// Copyright 2016 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package loader

import (
	"database/sql"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"

	"github.com/hexya-erp/hexya/src/models/operator"
	"github.com/hexya-erp/hexya/src/tools/strutils"
)

// ConnectionParams are the database agnostic parameters to connect to the database
type ConnectionParams struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	SSLCert  string
	SSLKey   string
	SSLCA    string
}

// A ColumnData holds information from the connector schema about one column
type ColumnData struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnDefault sql.NullString
}

// A seqData holds the data of a sequence in the database
type seqData struct {
	Name       string `connector:"sequence_name"`
	StartValue int64  `connector:"start_value"`
	Increment  int64  `connector:"increment"`
}

type DbAdapter interface {
	// connectionString returns the connection string for the given parameters
	connectionString(ConnectionParams) string
	// operatorSQL returns the sql string and placeholders for the given DomainOperator
	operatorSQL(operator.Operator, interface{}) (string, interface{})
	// typeSQL returns the SQL type string, including columns constraints if any
	typeSQL(fi *Field) string
	Connect() error
	Connector() *DatabaseConnector
	// columnSQLDefinition returns the SQL type string, including columns constraints if any
	//
	// If null is true, then the column will be nullable, whatever the field defines
	columnSQLDefinition(fi *Field, null bool) string
	// tables returns a map of table names of the database
	tables() map[string]bool
	// columns returns a list of ColumnData for the given tableName
	columns(tableName string) map[string]ColumnData
	// fieldIsNull returns true if the given Field results in a
	// NOT NULL column in database.
	fieldIsNotNull(fi *Field) bool
	// quoteTableName returns the given table name with sql quotes
	quoteTableName(string) string
	// indexExists returns true if an index with the given name exists in the given table
	indexExists(table string, name string) bool
	// constraintExists returns true if a constraint with the given name exists
	constraintExists(name string) bool
	// constraints returns a list of all constraints matching the given SQL pattern
	constraints(pattern string) []string
	// setTransactionIsolation returns the SQL string to set the transaction isolation
	// level to serializable
	setTransactionIsolation() string
	// createSequence creates a DB sequence with the given name
	CreateSequence(name string, increment, start int64)
	// dropSequence drop the DB sequence with the given name
	DropSequence(name string)
	// alterSequence modifies the DB sequence given by name
	AlterSequence(name string, increment, restart int64)
	// nextSequenceValue returns the next value of the given given sequence
	NextSequenceValue(name string) int64
	// sequences returns a list of all sequences matching the given SQL pattern
	sequences(pattern string) []seqData
	// childrenIdsQuery returns a query that finds all descendant of the given
	// a record from table including itself. The query has a placeholder for the
	// record's ID
	childrenIdsQuery(table string) string
	// substituteErrorMessage substitutes the given error's message by newMsg
	substituteErrorMessage(err error, newMsg string) error
	// isSerializationError returns true if the given error is a serialization error
	// and that the failed transaction should be retried.
	isSerializationError(err error) bool
}

// Cursor is a wrapper around a database transaction
type Cursor struct {
	adapter DbAdapter
	tx      *gorm.DB
}

// Execute a query without returning any rows. It panics in case of error.
// The args are for any placeholder parameters in the query.
func (c *Cursor) Execute(query string, args ...interface{}) (*sql.Rows, int64) {
	return c.adapter.Connector().dbExecute(c.tx, query, args...)
}

// Get queries a row into the database and maps the result into dest.
// The query must return only one row. Get panics on errors
func (c *Cursor) Get(dest interface{}, query string, args ...interface{}) {
	c.adapter.Connector().dbGet(c.tx, dest, query, args...)
}

// Select queries multiple rows and map the result into dest which must be a slice.
// Select panics on errors.
func (c *Cursor) Select(dest interface{}, query string, args ...interface{}) {
	c.adapter.Connector().dbSelect(c.tx, dest, query, args...)
}

// newCursor returns a new connector cursor on the given database
func newCursor(db *gorm.DB, adapter DbAdapter) *Cursor {
	adapter.Connector().dbExecute(db, adapter.setTransactionIsolation())
	return &Cursor{
		tx:      db,
		adapter: adapter,
	}
}

type DatabaseConnector struct {
	connParams ConnectionParams
	db         *gorm.DB
}

func DBConnect(params ConnectionParams) DbAdapter {
	connector := NewDatabaseConnector(&params)
	adapter = &postgresAdapter{
		connector: &connector,
	}
	adapter.Connect()
	return adapter
}

func NewDatabaseConnector(params *ConnectionParams) DatabaseConnector {
	if params == nil {
		params = &ConnectionParams{
			Driver:   viper.GetString("DB.Driver"),
			Host:     viper.GetString("DB.Host"),
			Port:     viper.GetString("DB.Port"),
			User:     viper.GetString("DB.User"),
			Password: viper.GetString("DB.Password"),
			DBName:   viper.GetString("DB.Name"),
			SSLMode:  viper.GetString("DB.SSLMode"),
			SSLCert:  viper.GetString("DB.SSLCert"),
			SSLKey:   viper.GetString("DB.SSLKey"),
			SSLCA:    viper.GetString("DB.SSLCA"),
		}
	}
	dd := DatabaseConnector{
		connParams: *params,
	}
	return dd
}

// DBParams returns the DB connection parameters currently in use
func (db *DatabaseConnector) DBParams() ConnectionParams {
	return db.connParams
}

func (db *DatabaseConnector) DB() *gorm.DB {
	return db.db
}

func (db *DatabaseConnector) MustExec(query string, args ...interface{}) int64 {
	var count int64
	err := db.db.Exec(query, args).Count(&count).Error
	if err != nil {
		count = -1
	}
	return count
}

// DBConnect connects to a database using the given driver and arguments.
func (db *DatabaseConnector) DBConnect(connStr string) {
	params := db.DBParams()
	var err error
	dbi, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Panic("Could not connect to database,", "driver", params.Driver, "connStr", connStr, "err", err)
	} else {
		db.db = dbi
		log.Info("Connected to database", "driver", params.Driver, "connStr", connStr)
	}
}

// DBClose is a wrapper around sqlx.Close
// It closes the connection to the database
func (db *DatabaseConnector) DBClose() {
	dbi, err := db.db.DB()
	if err != nil {
		log.Info("Closed database", "error", err)
	} else {
		err = dbi.Close()
		if err != nil {
			log.Info("Failed to closed database", "error", err)
		}
	}
}

// dbExecute is a wrapper around sqlx.MustExec
// It executes a query that returns no row
func (db *DatabaseConnector) dbExecute(cr *gorm.DB, query string, args ...interface{}) (*sql.Rows, int64) {
	query, args = db.sanitizeQuery(query, args...)
	t := time.Now()
	var cnt int64
	res, err := cr.Select(query, args...).Count(&cnt).Rows()
	logSQLResult(err, t, query, args...)
	return res, cnt
}

// dbExecuteNoTx simply executes the given query in the database without any transaction
func (db *DatabaseConnector) dbExecuteNoTx(query string, args ...interface{}) gorm.Rows {
	query, args = db.sanitizeQuery(query, args...)
	t := time.Now()
	res, err := db.db.Exec(query, args...).Rows()
	logSQLResult(err, t, query, args...)
	return res
}

// dbGet is a wrapper around sqlx.Get
// It gets the value of a single row found by the given query and arguments
// It panics in case of error
func (db DatabaseConnector) dbGet(cr *gorm.DB, dest interface{}, query string, args ...interface{}) {
	query, args = db.sanitizeQuery(query, args...)
	t := time.Now()
	err := cr.Select(query, args...).Scan(dest).Error
	logSQLResult(err, t, query, args)
}

// dbGetNoTx is a wrapper around sqlx.Get outside a transaction
// It gets the value of a single row found by the
// given query and arguments
func (db *DatabaseConnector) dbGetNoTx(dest interface{}, query string, args ...interface{}) {
	query, args = db.sanitizeQuery(query, args...)
	t := time.Now()
	err := db.db.Exec(query, args...).Scan(dest).Error
	logSQLResult(err, t, query, args)
}

// dbSelect is a wrapper around sqlx.Select
// It gets the value of a multiple rows found by the given query and arguments
// dest must be a slice. It panics in case of error
func (db *DatabaseConnector) dbSelect(cr *gorm.DB, dest interface{}, query string, args ...interface{}) {
	query, args = db.sanitizeQuery(query, args...)
	t := time.Now()
	err := cr.Select(query, args...).Scan(dest).Error
	logSQLResult(err, t, query, args)
}

// dbSelect is a wrapper around sqlx.Select outside a transaction
// It gets the value of a multiple rows found by the given query and arguments
// dest must be a slice. It panics in case of error
func (db DatabaseConnector) dbSelectNoTx(dest interface{}, query string, args ...interface{}) {
	query, args = db.sanitizeQuery(query, args...)
	t := time.Now()
	err := db.db.Select(query, args...).Scan(dest).Error
	logSQLResult(err, t, query, args)
}

// dbQuery is a wrapper around sqlx.Queryx
// It returns a sqlx.Rowsx found by the given query and arguments
// It panics in case of error
func (db *DatabaseConnector) dbQuery(cr *gorm.DB, query string, args ...interface{}) (*sql.Rows, int64) {
	query, args = db.sanitizeQuery(query, args...)
	var cnt int64
	t := time.Now()
	rows, err := cr.Select(query, args...).Count(&cnt).Rows()
	logSQLResult(err, t, query, args)
	return rows, cnt
}

// sanitizeQuery calls 'In' expansion and 'Rebind' on the given query and
// returns the new values to use. It panics in case of error
func (db *DatabaseConnector) sanitizeQuery(query string, args ...interface{}) (string, []interface{}) {
	log.Info("Unable to expand 'IN' statement", "query", query, "args", args)
	return query, args
}

// Log the result of the given sql query started at start time with the
// given args, and error. This function panics after logging if error is not nil.
func logSQLResult(err error, start time.Time, query string, args ...interface{}) {
	logCtx := log.New("query", query, "args", strutils.TrimArgs(args), "duration", time.Now().Sub(start))
	if err != nil {
		// We don't log.Panic to keep connector error information in recovery
		logCtx.Error("Error while executing query", "error", err)
		panic(err)
	}
	logCtx.Debug("Query executed")
}
