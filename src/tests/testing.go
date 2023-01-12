// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package tests

import (
	"fmt"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/loader"
	"os"
	"path/filepath"
	"testing"

	"github.com/hexya-erp/hexya/src/actions"
	"github.com/hexya-erp/hexya/src/controllers"
	"github.com/hexya-erp/hexya/src/menus"
	"github.com/hexya-erp/hexya/src/reports"
	"github.com/hexya-erp/hexya/src/server"
	"github.com/hexya-erp/hexya/src/templates"
	"github.com/hexya-erp/hexya/src/tools/logging"
	"github.com/hexya-erp/hexya/src/views"
	"github.com/spf13/viper"
)

var driver, user, password, prefix, debug string
var adapter loader.DbAdapter

// RunTests initializes the database, run the tests given by m and
// tears the database down.
//
// It is meant to be used for modules testing. Initialize your module's
// tests with:
//
//	    import (
//	        "testing"
//	        "github.com/hexya-erp/hexya/src/tests"
//	    )
//
//	    func TestMain(m *testing.M) {
//		       tests.RunTests(m, "my_module")
//	    }
func RunTests(m *testing.M, moduleName string, preHookFnct func()) {
	var res int
	defer func() {
		TearDownTests(moduleName)
		if r := recover(); r != nil {
			panic(r)
		}
		os.Exit(res)
	}()
	server.RegisterModule(&server.Module{
		Name:    moduleName,
		PreInit: preHookFnct,
	})
	InitializeTests(moduleName)
	res = m.Run()

}

// InitializeTests initializes a database for the tests of the given module.
// You probably want to use RunTests instead.
func InitializeTests(moduleName string) {
	fmt.Printf("Initializing tests for module %s\n", moduleName)
	driver = os.Getenv("HEXYA_DB_DRIVER")
	if driver == "" {
		driver = "postgres"
	}
	user = os.Getenv("HEXYA_DB_USER")
	if user == "" {
		user = "hexya"
	}
	password = os.Getenv("HEXYA_DB_PASSWORD")
	if password == "" {
		password = "hexya"
	}
	prefix = os.Getenv("HEXYA_DB_PREFIX")
	if prefix == "" {
		prefix = "hexya"
	}
	dbName := fmt.Sprintf("%s_%s_tests", prefix, moduleName)
	debug = os.Getenv("HEXYA_DEBUG")
	if debug == "" {
		debug = "Yes"
	}
	logTests := os.Getenv("HEXYA_LOG")

	viper.Set("LogLevel", "panic")
	if logTests != "" {
		viper.Set("LogLevel", "info")
		viper.Set("LogStdout", true)
	}
	if debug != "" {
		viper.Set("Debug", true)
		viper.Set("LogLevel", "debug")
		viper.Set("LogStdout", true)
	}
	logging.Initialize()

	server.PreInit()
	adapter = loader.DBConnect(loader.ConnectionParams{
		Driver:     driver,
		DBName:     dbName,
		User:       user,
		Debug:      true,
		AutoCreate: true,
		Password:   password,
		SSLMode:    "disable",
	})
	if adapter == nil {
		panic("Failed to initialize database")
	}
	keepDB := os.Getenv("HEXYA_KEEP_TEST_DB") != ""
	var count int64
	count = adapter.Connector().MustExec(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbName))
	if count <= 0 {
		fmt.Printf("Error: %v=%v", "value", count)
	}

	models.BootStrap()
	resourceDir, _ := filepath.Abs(filepath.Join(".", "res"))
	server.ResourceDir = resourceDir
	server.LoadInternalResources(resourceDir)
	if count > 0 || !keepDB {
		fmt.Println("Upgrading schemas in database", dbName)
		//models.SyncDatabase()
		fmt.Println("Loading resources into database", dbName)
		server.LoadDataRecords(resourceDir)
		server.LoadDemoRecords(resourceDir)
	}
	views.BootStrap()
	templates.BootStrap()
	actions.BootStrap()
	reports.BootStrap()
	controllers.BootStrap()
	menus.BootStrap()
	server.PostInit()
}

// TearDownTests tears down the tests for the given module
func TearDownTests(moduleName string) {
	keepDB := os.Getenv("HEXYA_KEEP_TEST_DB")
	if keepDB != "" {
		return
	}
	fmt.Printf("Tearing down database for module %s...", moduleName)
	dbName := fmt.Sprintf("%s_%s_tests", prefix, moduleName)
	//adapter.DropDatabase(dbName)
	fmt.Println("Ok: " + dbName)
	// Close connection
	adapter.Connector().DBClose()
}
