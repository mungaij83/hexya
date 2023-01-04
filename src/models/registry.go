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
	"errors"
	"fmt"
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/tools/strutils"
	"sync"
	"time"
)

// transientModelTimeout is the timeout after which transient model
// records can be removed from the database
var transientModelTimeout = 30 * time.Minute

// Registry is the registry of all Model instances.
var Registry *modelCollection

type modelCollection struct {
	sync.RWMutex
	loader              *ModelLoader
	bootstrapped        bool
	registryByTableName map[string]Repository[any, int64]
	sequences           map[string]*Sequence
}

// Get the given Model by name or by table name
func (mc *modelCollection) Get(nameOrJSON string) (mi *loader.Model, ok bool) {
	repo, ok := mc.registryByTableName[nameOrJSON]
	if !ok {
		return nil, false
	}
	mi, ok = repo.GetModel()
	return
}

// MustGet the given Model by name or by table name.
// It panics if the Model does not exist
func (mc *modelCollection) MustGet(nameOrJSON string) *loader.Model {
	mi, ok := mc.Get(nameOrJSON)
	if !ok {
		log.Panic("Unknown model", "model", nameOrJSON)
	}
	return mi
}

// GetSequence the given Sequence by name or by db name
func (mc *modelCollection) GetSequence(nameOrJSON string) (s *Sequence, ok bool) {
	s, ok = mc.sequences[nameOrJSON]
	if !ok {
		jsonBoot := strutils.SnakeCase(nameOrJSON) + "_bootseq"
		s, ok = mc.sequences[jsonBoot]
		if !ok {
			jsonMan := strutils.SnakeCase(nameOrJSON) + "_manseq"
			s, ok = mc.sequences[jsonMan]
		}
	}
	return
}

// MustGetSequence gets the given sequence by name or by db name.
// It panics if the Sequence does not exist
func (mc *modelCollection) MustGetSequence(nameOrJSON string) *Sequence {
	s, ok := mc.GetSequence(nameOrJSON)
	if !ok {
		log.Panic("Unknown sequence", "sequence", nameOrJSON)
	}
	return s
}

// add the given Model to the modelCollection
func (mc *modelCollection) add(mi Repository[any, int64]) error {
	// Initialize repository
	err := mi.validateAndInitialize(mc.loader)
	if err != nil {
		log.Warn("Failed to initialize model repository", "Error", err)
		return err
	}
	// Initialize table
	if _, exists := mc.Get(mi.TableName()); exists {
		log.Warn("Trying to add already existing model", "model", mi.TableName())
		return errors.New(fmt.Sprintf("trying to initialize an existing model: %v", mi.TableName()))
	}
	// Register default model extension
	err = mi.RegisterExtension(DefaultMixinExtension[any]{})
	if err != nil {
		log.Warn("register extension for this model")
	}
	// Register model
	mc.registryByTableName[mi.TableName()] = mi
	return nil
}

// add the given Model to the modelCollection
func (mc *modelCollection) addSequence(s *Sequence) {
	if _, exists := mc.GetSequence(s.JSON); exists {
		log.Panic("Trying to add already existing sequence", "sequence", s.JSON)
	}
	mc.Lock()
	defer mc.Unlock()
	mc.sequences[s.JSON] = s
}

// newModelCollection returns a pointer to a new modelCollection
func newModelCollection() *modelCollection {
	return &modelCollection{
		loader:              &ModelLoader{},
		registryByTableName: make(map[string]Repository[any, int64]),
		sequences:           make(map[string]*Sequence),
	}
}
