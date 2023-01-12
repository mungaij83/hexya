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
	"fmt"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/tools/logging"
	"gorm.io/gorm"
)

// DBSerializationMaxRetries defines the number of time a
// transaction that failed due to serialization error should
// be retried.
const DBSerializationMaxRetries uint8 = 5

// maxRecursionDepth is the maximum allowed number of nested calls
// during a transaction.
const maxRecursionDepth uint8 = 100

// An Environment stores various contextual data used by the models:
// - the database cursor (current open transaction),
// - the current user ID (for access rights checking)
// - the current context (for storing arbitrary metadata).
// The Environment also stores caches.
type Environment struct {
	cr             *Cursor
	uid            int64
	driverName     string
	context        *types.Context
	cache          *cache
	currentLayer   *methodLayer
	previousMethod *Method
	super          bool
	testMode       bool
	recursions     uint8
	nextNegativeID int64
}

// Cr returns a pointer to the Cursor of the Environment
func (env Environment) Cr() *gorm.DB {
	if env.testMode {
		return env.cr.tx.Debug()
	}
	return env.cr.tx
}

func (env Environment) Cursor() *Cursor {
	return env.cr
}

// Uid returns the user id of the Environment
func (env Environment) Uid() int64 {
	return env.uid
}

// DriverName returns the user database driver name Environment
func (env Environment) DriverName() string {
	return env.cr.adapter.Connector().DBParams().Driver
}

// Context returns the Context of the Environment
func (env Environment) Context() *types.Context {
	return env.context
}

// commit the transaction of this environment.
//
// WARNING: Do NOT call Commit on Environment instances that you
// did not create yourself with NewEnvironment. The framework will
// automatically commit the Environment.
func (env Environment) commit() error {
	return env.Cr().Commit().Error
}

// rollback the transaction of this environment.
//
// WARNING: Do NOT call Rollback on Environment instances that you
// did not create yourself with NewEnvironment. Just panic instead
// for the framework to roll back automatically for you.
func (env Environment) rollback() error {
	return env.Cr().Rollback().Error
}

// checkRecursion panics if the recursion depth limit is reached
func (env Environment) checkRecursion() {
	if env.recursions > maxRecursionDepth {
		log.Panic("Max recursion depth exceeded")
	}
}

// DumpCache returns a human readable string of this Environment's
// cache for debugging purposes.
func (env Environment) DumpCache() string {
	res := "Data\n====\n"
	for modelName, model := range env.cache.data {
		for id, record := range model {
			res += fmt.Sprintf("> Model: %s, ID: %d\n", modelName, id)
			for field, value := range record {
				res += fmt.Sprintf("    %s: %#v\n", field, value)
			}
		}
	}
	res += "\nM2M Links\n=========\n"
	for relModel, pairs := range env.cache.m2mLinks {
		res += fmt.Sprintf("> RelModel: %s, pairs: %v\n", relModel, pairs)
	}
	res += "\nX2M Links\n=========\n"
	for modelName, records := range env.cache.x2mRelated {
		for id, record := range records {
			res += fmt.Sprintf("> Model: %s, ID: %d\n", modelName, id)
			for fieldName, fields := range record {
				res += fmt.Sprintf("    %s: %#v\n", fieldName, fields)
			}
		}
	}
	return res
}

// newEnvironment returns a new Environment for the given user ID
//
// WARNING: Callers to newEnvironment should ensure to either call commit()
// or rollback() on the returned Environment after operation to release
// the database connection.
func newEnvironment(uid int64) Environment {
	env := Environment{
		cr:       newCursor(adapter.Connector().DB(), adapter),
		uid:      uid,
		testMode: adapter.Connector().connParams.Debug,
		context:  types.NewContext(),
		cache:    newCache(),
	}
	return env
}

// ExecuteInNewEnvironment executes the given fnct in a new Environment
// within a new transaction.
//
// This function commits the transaction if everything went right or
// rolls it back otherwise, returning an arror. Database serialization
// errors are automatically retried several times before returning an
// error if they still occur.
func ExecuteInNewEnvironment(uid int64, fnct func(Environment)) error {
	return doExecuteInNewEnvironment(uid, 0, fnct)
}

func doExecuteInNewEnvironment(uid int64, retries uint8, fnct func(Environment)) (rError error) {
	env := newEnvironment(uid)
	defer func() {
		if r := recover(); r != nil {
			env.rollback()
			if err, ok := r.(error); ok && adapter.isSerializationError(err) {
				// Transaction error
				retries++
				if retries < DBSerializationMaxRetries {
					if doExecuteInNewEnvironment(uid, retries, fnct) == nil {
						rError = nil
						return
					}
				}
			}
			rError = logging.LogPanicData(r)
			return
		}
		env.commit()
	}()
	fnct(env)
	return nil
}

// SimulateInNewEnvironment executes the given fnct in a new Environment
// within a new transaction and rolls back the transaction at the end.
//
// This function always rolls back the transaction but returns an error
// only if fnct panicked during its execution.
func SimulateInNewEnvironment(uid int64, fnct func(Environment)) error {
	return doSimulateInNewEnvironment(uid, 0, fnct)
}

func doSimulateInNewEnvironment(uid int64, retries uint8, fnct func(Environment)) (rError error) {
	env := newEnvironment(uid)
	defer func() {
		env.rollback()
		if r := recover(); r != nil {
			if err, ok := r.(error); ok && adapter.isSerializationError(err) {
				// Transaction error. We try again even if we rollback anyway
				// to be as close as ExecuteInNewEnvironment as possible
				retries++
				if retries < DBSerializationMaxRetries {
					if doSimulateInNewEnvironment(uid, retries, fnct) == nil {
						rError = nil
						return
					}
				}
			}
			rError = logging.LogPanicData(r)
			return
		}
	}()
	fnct(env)
	return
}

// Pool returns an empty RecordCollection for the given modelName
func (env Environment) Pool(modelName string) *RecordCollection {
	return newRecordCollection(env, modelName)
}
