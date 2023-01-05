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

package models

import (
	"fmt"
	"github.com/hexya-erp/hexya/src/models/loader"
	"os"
	"testing"

	"github.com/hexya-erp/hexya/src/tools/logging"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var dbArgs = struct {
	Driver   string
	User     string
	Password string
	DB       string
	Debug    string
}{}

var TestAdapter loader.DbAdapter

func TestMain(m *testing.M) {
	initializeTests()
	res := m.Run()
	tearDownTests()
	os.Exit(res)
}

func initializeTests() {
	fmt.Printf("Initializing database for models\n")
	dbArgs.Driver = os.Getenv("HEXYA_DB_DRIVER")
	if dbArgs.Driver == "" {
		dbArgs.Driver = "postgres"
	}
	dbArgs.User = os.Getenv("HEXYA_DB_USER")
	if dbArgs.User == "" {
		dbArgs.User = "hexya"
	}
	dbArgs.Password = os.Getenv("HEXYA_DB_PASSWORD")
	if dbArgs.Password == "" {
		dbArgs.Password = "hexya"
	}
	prefix := os.Getenv("HEXYA_DB_PREFIX")
	if prefix == "" {
		prefix = "hexya"
	}

	dbArgs.DB = fmt.Sprintf("%s_models_tests", prefix)
	dbArgs.Debug = os.Getenv("HEXYA_DEBUG")

	viper.Set("LogLevel", "panic")
	if dbArgs.Debug == "" {
		viper.Set("Debug", true)
		viper.Set("LogLevel", "debug")
		viper.Set("LogStdout", true)
	}
	logging.Initialize()

	TestAdapter = loader.DBConnect(loader.ConnectionParams{
		Driver:     dbArgs.Driver,
		DBName:     dbArgs.DB,
		Debug:      true,
		AutoCreate: true,
		User:       dbArgs.User,
		Password:   dbArgs.Password,
		SSLMode:    "disable",
	})
}

func tearDownTests() {
	keepDB := os.Getenv("HEXYA_KEEP_TEST_DB")
	if !TestAdapter.Connector().DBParams().AutoCreate {
		log.Debug("Keep database", "value", keepDB)
		TestAdapter.Close()
		return
	}
	fmt.Printf("Tearing down database for models\n")
	TestAdapter.DropDatabase(dbArgs.DB)
	fmt.Printf("Close database for models\n")
	TestAdapter.Close()
}
