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

package conditions

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

var (
	// ID is a FieldName that represents the PK of a model
	ID = NewFieldName("ID", "id")
	// Name is a fieldName that represents the Name field of a model
	Name = NewFieldName("Name", "name")
)

// A RecordRef uniquely identifies a Record by giving its model and ID.
type RecordRef struct {
	ModelName string
	ID        int64
}

// RecordSet identifies a type that holds a set of records of
// a given model.
type RecordSet interface {
	sql.Scanner
	fmt.Stringer
	// ModelName returns the name of the model of this RecordSet
	ModelName() string
	// Ids returns the ids in this set of Records
	Ids() []int64
	// Env returns the current Environment of this RecordSet
	Env() interface{}
	// Len returns the number of records in this RecordSet
	Len() int
	// IsValid returns true if this RecordSet has been initialized.
	IsValid() bool
	// IsEmpty returns true if this RecordSet has no records
	IsEmpty() bool
	// IsNotEmpty returns true if this RecordSet has at least one record
	IsNotEmpty() bool
	// Call executes the given method (as string) with the given arguments
	Call(string, ...interface{}) interface{}
	// Collection returns the underlying RecordCollection instance
	Collection() interface{}
	// Get returns the value of the given fieldName for the first record of this RecordCollection.
	// It returns the type's zero value if the RecordCollection is empty.
	Get(FieldName) interface{}
	// Set sets field given by fieldName to the given value. If the RecordSet has several
	// Records, all of them will be updated. Each call to Set makes an update query in the
	// database. It panics if it is called on an empty RecordSet.
	Set(FieldName, interface{})
	// T translates the given string to the language specified by
	// the 'lang' key of rc.Env().Context(). If for any reason the
	// string cannot be translated, then src is returned.
	//
	// You MUST pass a string literal as src to have it extracted automatically (and not a variable)
	//
	// The translated string will be passed to fmt.Sprintf with the optional args
	// before being returned.
	T(string, ...interface{}) string
	// EnsureOne panics if this Recordset is not a singleton
	EnsureOne()
}

// A FieldName is a type that can represents a field in a model.
// It can yield the field name or the field's JSON name as a string
type FieldName interface {
	Name() string
	JSON() string
}

// fieldName is a simple implementation of FieldName
type fieldName struct {
	name string
	json string
}

// Name returns the field's name
func (f fieldName) Name() string {
	return f.name
}

// JSON returns the field's json name
func (f fieldName) JSON() string {
	return f.json
}

// NewFieldName returns a fieldName instance with the given name and json
func NewFieldName(name, json string) FieldName {
	return fieldName{name: name, json: json}
}

// FieldNames is a slice of FieldName that can be sorted
type FieldNames []FieldName

// Len returns the length of the FieldName slice
func (f FieldNames) Len() int {
	return len(f)
}

// Less returns true if f[i] < f[j]. FieldNames are ordered by JSON names
func (f FieldNames) Less(i, j int) bool {
	return f[i].JSON() < f[j].JSON()
}

// Swap i and j indexes
func (f FieldNames) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// UnmarshalJSON for the FieldNames type
func (f *FieldNames) UnmarshalJSON(data []byte) error {
	var aux []string
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	for _, v := range aux {
		*f = append(*f, NewFieldName(v, v))
	}
	return nil
}

// Names returns a slice with the names of each field
func (f FieldNames) Names() []string {
	var res []string
	for _, fn := range f {
		res = append(res, fn.Name())
	}
	return res
}

// JSON returns a slice with the JSON names of each field
func (f FieldNames) JSON() []string {
	var res []string
	for _, fn := range f {
		res = append(res, fn.JSON())
	}
	return res
}

// A Conditioner can return a Condition object through its Underlying() method
type Conditioner interface {
	Underlying() *Condition
}
