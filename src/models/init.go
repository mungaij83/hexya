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
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/tools/logging"
)

var (
	log logging.Logger
	// Views is a map to store views created automatically.
	// It will be processed by the views package and added to the views registry.
	Views map[*loader.Model][]string
	// Method that cannot be overridden
	unauthorizedMethods = map[string]bool{
		"Load":   true,
		"Create": true,
		"Write":  true,
		"Unlink": true,
	}
)

func init() {
	log = logging.GetLogger("models")
	// model registry
	Registry = newModelCollection()
	loader.RegisterModelLoader(func(name string) *loader.Model {
		return Registry.MustGet(name)
	})
	Views = make(map[*loader.Model][]string)

}
