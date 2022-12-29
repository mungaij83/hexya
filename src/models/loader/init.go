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
	"github.com/hexya-erp/hexya/src/tools/logging"
	"reflect"
)

var (
	log         logging.Logger
	adapter     DbAdapter
	lookupModel func(name string) *Model
	adapters    map[string]DbAdapter
	// recordSetWrappers is a map that stores the available types of RecordSet
	RecordSetWrappers map[string]reflect.Type
)

// registerDBAdapter adds a adapter to the adapters registry
// name of the adapter should match the database/sql driver name
func RegisterDBAdapter(name string, adapter DbAdapter) {
	adapters[name] = adapter
}

func RegisterModelLoader(loaderFunc func(string) *Model) {
	lookupModel = loaderFunc
}

func init() {
	log = logging.GetLogger("models")
	// DB drivers
	adapters = make(map[string]DbAdapter)
	RecordSetWrappers = make(map[string]reflect.Type)
	modelDataWrappers = make(map[string]reflect.Type)
	RegisterDBAdapter("postgres", new(postgresAdapter))
}

func GetAdapter() DbAdapter {
	return adapter
}
