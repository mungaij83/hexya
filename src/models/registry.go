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
	"sync"
	"time"

	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/hexya/src/tools/strutils"
)

// transientModelTimeout is the timeout after which transient model
// records can be removed from the database
var transientModelTimeout = 30 * time.Minute

// Registry is the registry of all Model instances.
var Registry *modelCollection

// Option describes a optional feature of a model
type Option int

type modelCollection struct {
	sync.RWMutex
	bootstrapped        bool
	registryByTableName map[string]Repository[any, int64]
	sequences           map[string]*Sequence
}

// Get the given Model by name or by table name
func (mc *modelCollection) Get(nameOrJSON string) (mi Repository[any, int64], ok bool) {
	mi, ok = mc.registryByTableName[nameOrJSON]
	if !ok {
		mi, ok = mc.registryByTableName[nameOrJSON]
	}
	return
}

// MustGet the given Model by name or by table name.
// It panics if the Model does not exist
func (mc *modelCollection) MustGet(nameOrJSON string) Repository[any, int64] {
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
func (mc *modelCollection) add(mi Repository[any, int64]) {
	if _, exists := mc.Get(mi.TableName()); exists {
		log.Panic("Trying to add already existing model", "model", mi.TableName())
	}
	mc.registryByTableName[mi.TableName()] = mi
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
		registryByTableName: make(map[string]Repository[any, int64]),
		sequences:           make(map[string]*Sequence),
	}
}

// A Sequence holds the metadata of a DB sequence
//
// There are two types of sequences: those created before bootstrap
// and those created after. The former will be created and updated at
// bootstrap and cannot be modified afterwards. The latter will be
// created, updated or dropped immediately.
type Sequence struct {
	JSON      string
	Increment int64
	Start     int64
	boot      bool
}

// CreateSequence creates a new Sequence in the database and returns a pointer to it
func CreateSequence(name string, increment, start int64) *Sequence {
	var boot bool
	suffix := "manseq"
	if !Registry.bootstrapped {
		boot = true
		suffix = "bootseq"
	}
	json := fmt.Sprintf("%s_%s", strutils.SnakeCase(name), suffix)
	seq := &Sequence{
		JSON:      json,
		Increment: increment,
		Start:     start,
		boot:      boot,
	}
	if !boot {
		// Create the sequence on the fly if we already bootstrapped.
		// Otherwise, this will be done in Bootstrap
		adapters[connParams.Driver].createSequence(seq.JSON, seq.Increment, seq.Start)
	}
	Registry.addSequence(seq)
	return seq
}

// Drop this sequence and removes it from the database
func (s *Sequence) Drop() {
	Registry.Lock()
	defer Registry.Unlock()
	delete(Registry.sequences, s.JSON)
	if Registry.bootstrapped {
		// Drop the sequence on the fly if we already bootstrapped.
		// Otherwise, this will be done in Bootstrap
		if s.boot {
			log.Panic("Boot Sequences cannot be dropped after bootstrap")
		}
		adapters[connParams.Driver].dropSequence(s.JSON)
	}
}

// Alter alters this sequence by changing next number and/or increment.
// Set a parameter to 0 to leave it unchanged.
func (s *Sequence) Alter(increment, restart int64) {
	var boot bool
	if !Registry.bootstrapped {
		boot = true
	}
	if s.boot && !boot {
		log.Panic("Boot Sequences cannot be modified after bootstrap")
	}
	if restart > 0 {
		s.Start = restart
	}
	if increment > 0 {
		s.Increment = increment
	}
	if !boot {
		adapters[connParams.Driver].alterSequence(s.JSON, increment, restart)
	}
}

// NextValue returns the next value of this Sequence
func (s *Sequence) NextValue() int64 {
	adapter := adapters[connParams.Driver]
	return adapter.nextSequenceValue(s.JSON)
}

// FreeTransientModels remove transient models records from database which are
// older than the given timeout.
func FreeTransientModels() {
	for _, model := range Registry.registryByTableName {
		if model.IsTransient() {
			err := ExecuteInNewEnvironment(security.SuperUserID, func(env Environment) {
				//createDate := model.FieldName("CreateDate")
				//model.Search(env, model.GetField(createDate).Lower(dates.Now().Add(-transientModelTimeout))).Call("Unlink")
			})
			if err != nil {
				log.Warn("Failed to free transient models")
			}
		}
	}
}
