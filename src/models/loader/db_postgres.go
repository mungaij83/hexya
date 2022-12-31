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
	"errors"
	"fmt"
	"github.com/hexya-erp/hexya/src/models/fieldtype"
	"github.com/hexya-erp/hexya/src/models/operator"
	"github.com/hexya-erp/hexya/src/tools/nbutils"
	"github.com/lib/pq"
)

type postgresAdapter struct {
	connector *DatabaseConnector
}

var pgOperators = map[operator.Operator]string{
	operator.Equals:         "= ?",
	operator.NotEquals:      "!= ?",
	operator.Contains:       "LIKE ?",
	operator.NotContains:    "NOT LIKE ?",
	operator.Like:           "LIKE ?",
	operator.IContains:      "ILIKE ?",
	operator.NotIContains:   "NOT ILIKE ?",
	operator.ILike:          "ILIKE ?",
	operator.In:             "IN (?)",
	operator.NotIn:          "NOT IN (?)",
	operator.Lower:          "< ?",
	operator.LowerOrEqual:   "<= ?",
	operator.Greater:        "> ?",
	operator.GreaterOrEqual: ">= ?",
}

var pgTypes = map[fieldtype.Type]string{
	fieldtype.Boolean:   "boolean",
	fieldtype.Char:      "character varying",
	fieldtype.Text:      "text",
	fieldtype.Date:      "date",
	fieldtype.DateTime:  "timestamp without time zone",
	fieldtype.Integer:   "integer",
	fieldtype.Float:     "numeric",
	fieldtype.HTML:      "text",
	fieldtype.Binary:    "bytea",
	fieldtype.Selection: "character varying",
	fieldtype.Many2One:  "integer",
	fieldtype.One2One:   "integer",
}

// connectionString returns the connection string for the given parameters
func (d *postgresAdapter) connectionString(params ConnectionParams, withDb bool) string {
	connectString := ""
	if withDb {
		connectString = fmt.Sprintf("dbname=%s", params.DBName)
	}
	if params.SSLMode != "" {
		connectString += fmt.Sprintf(" sslmode=%s", params.SSLMode)
	}
	if params.SSLCert != "" {
		connectString += fmt.Sprintf(" sslcert=%s", params.SSLCert)
	}
	if params.SSLKey != "" {
		connectString += fmt.Sprintf(" sslkey=%s", params.SSLKey)
	}
	if params.SSLCA != "" {
		connectString += fmt.Sprintf(" sslrootcert=%s", params.SSLCA)
	}
	if params.User != "" {
		connectString += fmt.Sprintf(" user=%s", params.User)
	}
	if params.Password != "" {
		connectString += fmt.Sprintf(" password=%s", params.Password)
	}
	if params.Host != "" {
		connectString += fmt.Sprintf(" host=%s", params.Host)
	}
	if params.Port != "" && params.Port != "5432" {
		connectString += fmt.Sprintf(" port=%s", params.Port)
	}
	return connectString
}

// operatorSQL returns the sql string and placeholders for the given DomainOperator
// Also modifies the given args to match the syntax of the operator.
func (d *postgresAdapter) operatorSQL(do operator.Operator, arg interface{}) (string, interface{}) {
	op := pgOperators[do]
	switch do {
	case operator.Contains, operator.IContains, operator.NotContains, operator.NotIContains:
		arg = fmt.Sprintf("%%%s%%", arg)
	}
	return op, arg
}
func (d *postgresAdapter) Connector() *DatabaseConnector {
	return d.connector
}

func (d *postgresAdapter) Connect() (rerr error) {
	defer func() {
		if err := recover(); err != nil {
			var ok bool
			rerr, ok = err.(error)
			if !ok {
				rerr = errors.New(err.(string))
			}
		}
	}()
	autoCreate := d.connector.connParams.AutoCreate
	if autoCreate {
		rerr = d.connector.DBConnect(d.connectionString(d.connector.DBParams(), false))

		if !d.connector.createDatabaseIfNotExist() {
			log.Debug("Failed to create database: ", "value", rerr)
		} else {
			// Close database and reconnect to the created database
			d.connector.DBClose()
			rerr = d.connector.DBConnect(d.connectionString(d.connector.connParams, true))
		}
	} else {
		rerr = d.connector.DBConnect(d.connectionString(d.connector.DBParams(), false))
	}
	return
}

// typeSQL returns the sql type string for the given Field
func (d *postgresAdapter) typeSQL(fi *Field) string {
	typ, _ := pgTypes[fi.FieldType]
	return typ
}

// columnSQLDefinition returns the SQL type string, including columns constraints if any
//
// If null is true, then the column will be nullable, whatever the field defines
func (d *postgresAdapter) columnSQLDefinition(fi *Field, null bool) string {
	var res string
	typ, ok := pgTypes[fi.FieldType]
	res = typ
	if !ok {
		log.Panic("Unknown column type", "type", fi.FieldType, "model", fi.model, "field", fi.name)
	}
	switch fi.FieldType {
	case fieldtype.Char:
		if fi.size > 0 {
			res = fmt.Sprintf("%s(%d)", res, fi.size)
		}
	case fieldtype.Float:
		emptyD := nbutils.Digits{}
		if fi.digits != emptyD {
			res = fmt.Sprintf("numeric(%d, %d)", fi.digits.Precision, fi.digits.Scale)
		}
	}
	if d.fieldIsNotNull(fi) && !null {
		res += " NOT NULL"
	}

	if fi.unique || fi.FieldType == fieldtype.One2One {
		res += " UNIQUE"
	}
	return res
}

// fieldIsNull returns true if the given Field results in a
// NOT NULL column in database.
func (d *postgresAdapter) fieldIsNotNull(fi *Field) bool {
	if fi.required {
		return true
	}
	return false
}

// tables returns a map of table names of the database
func (d *postgresAdapter) tables() map[string]bool {
	var resList []string
	query := "SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema NOT IN ('pg_catalog', 'information_schema')"
	if err := d.connector.DB().Select(&resList, query); err != nil {
		log.Panic("Unable to get list of tables from database", "error", err)
	}
	res := make(map[string]bool, len(resList))
	for _, tableName := range resList {
		res[tableName] = true
	}
	return res
}

// quoteTableName returns the given table name with sql quotes
func (d *postgresAdapter) quoteTableName(tableName string) string {
	return fmt.Sprintf(`"%s"`, tableName)
}

// columns returns a list of ColumnData for the given tableName
func (d *postgresAdapter) columns(tableName string) map[string]ColumnData {
	query := fmt.Sprintf(`
		SELECT column_name, data_type, is_nullable, column_default
		FROM information_schema.columns
		WHERE table_schema NOT IN ('pg_catalog', 'information_schema') AND table_name = '%s'
	`, tableName)
	var colData []ColumnData
	if err := d.connector.DB().Select(&colData, query); err != nil {
		log.Panic("Unable to get list of columns for table", "table", tableName, "error", err)
	}
	res := make(map[string]ColumnData, len(colData))
	for _, col := range colData {
		res[col.ColumnName] = col
	}
	return res
}

// indexExists returns true if an index with the given name exists in the given table
func (d *postgresAdapter) indexExists(table string, name string) bool {
	query := fmt.Sprintf("SELECT COUNT(*) FROM pg_indexes WHERE tablename = '%s' AND indexname = '%s'", table, name)
	var cnt int
	d.connector.dbGetNoTx(&cnt, query)
	return cnt > 0
}

// constraintExists returns true if a constraint with the given name exists in the given table
func (d *postgresAdapter) constraintExists(name string) bool {
	query := fmt.Sprintf("SELECT COUNT(*) FROM pg_constraint WHERE conname = '%s'", name)
	var cnt int
	d.connector.dbGetNoTx(&cnt, query)
	return cnt > 0
}

// constraints returns a list of all constraints matching the given SQL pattern
func (d *postgresAdapter) constraints(pattern string) []string {
	query := "SELECT conname FROM pg_constraint WHERE conname ILIKE ?"
	var res []string
	d.connector.dbSelectNoTx(&res, query, pattern)
	return res
}

// createSequence creates a DB sequence with the given name
func (d *postgresAdapter) CreateSequence(name string, increment, start int64) {
	query := fmt.Sprintf("CREATE SEQUENCE %s INCREMENT BY %d START WITH %d", name, increment, start)
	d.connector.dbExecuteNoTx(query)
}

// dropSequence drops the DB sequence with the given name
func (d *postgresAdapter) DropSequence(name string) {
	query := fmt.Sprintf("DROP SEQUENCE IF EXISTS %s", name)
	d.connector.dbExecuteNoTx(query)
}

// alterSequence modifies the DB sequence given by name
func (d *postgresAdapter) AlterSequence(name string, increment, restart int64) {
	query := fmt.Sprintf(`ALTER SEQUENCE %s`, name)
	if increment != 0 {
		query += fmt.Sprintf(` INCREMENT BY %d`, increment)
	}
	if restart != 0 {
		query += fmt.Sprintf(` RESTART WITH %d`, restart)
	}
	d.connector.dbExecuteNoTx(query)
}

// nextSequenceValue returns the next value of the given given sequence
func (d *postgresAdapter) NextSequenceValue(name string) int64 {
	query := fmt.Sprintf("SELECT nextval('%s')", name)
	var val int64
	d.connector.dbGetNoTx(&val, query)
	return val
}

// sequences returns a list of all sequences matching the given SQL pattern
func (d *postgresAdapter) Sequences(pattern string) []seqData {
	query := "SELECT sequence_name, start_value, increment FROM information_schema.sequences WHERE sequence_name ILIKE ?"
	log.Info("Sequence query:", "query", query)
	var res []seqData
	d.connector.dbSelectNoTx(&res, query, pattern)
	return res
}

// setTransactionIsolation returns the SQL string to set the
// transaction isolation level to serializable
func (d *postgresAdapter) setTransactionIsolation() string {
	return "SET TRANSACTION ISOLATION LEVEL SERIALIZABLE"
}

// childrenIdsQuery returns a query that finds all descendant of the given
// a record from table including itself. The query has a placeholder for the
// record's ID
func (d *postgresAdapter) childrenIdsQuery(table string) string {
	res := fmt.Sprintf(`
WITH RECURSIVE "recursive_query_children_ids" AS
(
	SELECT  id
	FROM    %s "m1"
	WHERE   id = ?
UNION ALL
	SELECT  "m2".id
	FROM    %s "m2"
	JOIN    "recursive_query_children_ids"
	ON      "m2".parent_id = "recursive_query_children_ids".id
)
SELECT  id
FROM    recursive_query_children_ids`, d.quoteTableName(table), d.quoteTableName(table))
	return res
}

// substituteErrorMessage substitutes the given error's message by newMsg
func (d *postgresAdapter) substituteErrorMessage(err error, newMsg string) error {
	pgError, ok := err.(*pq.Error)
	if !ok {
		return err
	}
	pgError.Message = newMsg
	return pgError
}

// isSerializationError returns true if the given error is a serialization error
// and that the failed transaction should be retried.
func (d *postgresAdapter) isSerializationError(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Class() == "40" {
		return true
	}
	return false
}

func (d *postgresAdapter) Close() (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			ok = false
		}
	}()
	d.Connector().DBClose()
	ok = true
	return
}
