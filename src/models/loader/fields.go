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
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/tools/nbutils"
	"github.com/hexya-erp/hexya/src/tools/strutils"
	"reflect"
	"sync"
)

// An OnDeleteAction defines what to be done with this record when
// the target record is deleted.
type OnDeleteAction string

const (
	// SetNull sets the foreign key to null in referencing records. This is the default
	SetNull OnDeleteAction = "set null"
	// Restrict throws an error if there are record referencing the deleted one.
	Restrict OnDeleteAction = "restrict"
	// Cascade deletes all referencing records.
	Cascade OnDeleteAction = "cascade"
)

type ctxType int

const (
	ctxNone = iota
	ctxValue
	ctxContext
	ctxFK
)

// computeData holds data to recompute another field.
// - model is a pointer to the Model instance to recompute
// - fieldName is the name of the field to recompute in model.
// - compute is the name of the method to call on model
// - path is the search string that will be used to find records to update
// (e.g. path = "Profile.BestPost").
// - stored is true if the computed field is stored
type computeData struct {
	model     *Model
	stored    bool
	fieldName string
	compute   string
	path      string
}

// FieldsCollection is a collection of Field instances in a model.
type FieldsCollection struct {
	sync.RWMutex
	model                *Model
	registryByName       map[string]*Field
	registryByJSON       map[string]*Field
	computedFields       []*Field
	computedStoredFields []*Field
	relatedFields        []*Field
	bootstrapped         bool
}

// Get returns the Field of the field with the given name.
// name can be either the name of the field or its JSON name.
func (fc *FieldsCollection) Get(name string) (fi *Field, ok bool) {
	fi, ok = fc.registryByName[name]
	if !ok {
		fi, ok = fc.registryByJSON[name]
	}
	return
}

// MustGet returns the Field of the field with the given name or panics
// name can be either the name of the field or its JSON name.
func (fc *FieldsCollection) MustGet(name string) *Field {
	fi, ok := fc.Get(name)
	if !ok {
		log.Panic("Unknown field in model", "model", fc.model.TableName(), "field", name)
	}
	return fi
}

// storedFieldNames returns a slice with the names of all the stored fields
// If fields are given, return only names in the list
func (fc *FieldsCollection) storedFieldNames(fieldNames ...FieldName) []FieldName {
	var res []FieldName
	for fName, fi := range fc.registryByName {
		var keepField bool
		if len(fieldNames) == 0 {
			keepField = true
		} else {
			for _, f := range fieldNames {
				if fName == f.Name() {
					keepField = true
					break
				}
			}
		}
		if (fi.isStored() || fi.isRelatedField()) && keepField {
			res = append(res, fc.model.Fields().MustGet(fName))
		}
	}
	return res
}

// allFieldNames returns a slice with the name of all field's JSON names of this collection
func (fc *FieldsCollection) allFieldNames() FieldNames {
	res := make([]FieldName, len(fc.registryByJSON))
	var i int
	for f := range fc.registryByName {
		res[i] = fc.model.Fields().MustGet(f)
		i++
	}
	return res
}

// getComputedFields returns the slice of Field of the computed, but not
// stored fields of the given modelName.
// If fields are given, return only Field instances in the list
func (fc *FieldsCollection) getComputedFields(fields ...string) (fil []*Field) {
	fInfos := fc.computedFields
	if len(fields) > 0 {
		for _, f := range fields {
			for _, fInfo := range fInfos {
				if f == fInfo.name || f == fInfo.json {
					fil = append(fil, fInfo)
					continue
				}
			}
		}
	} else {
		fil = fInfos
	}
	return
}

// Model returns this FieldsCollection Model
func (fc *FieldsCollection) Model() *Model {
	return fc.model
}

// newFieldsCollection returns a pointer to a new empty FieldsCollection with
// all maps initialized.
func newFieldsCollection() *FieldsCollection {
	return &FieldsCollection{
		registryByName: make(map[string]*Field),
		registryByJSON: make(map[string]*Field),
	}
}

// add the given Field to the FieldsCollection.
func (fc *FieldsCollection) add(fInfo *Field) {
	if _, exists := fc.registryByName[fInfo.name]; exists {
		log.Panic("Trying to add already existing field", "model", fInfo.model, "field", fInfo.name)
	}
	fc.register(fInfo)
}

// register adds the given fInfo in the collection.
func (fc *FieldsCollection) register(fInfo *Field) {
	fc.Lock()
	defer fc.Unlock()

	checkFieldInfo(fInfo)
	name := fInfo.name
	jsonName := fInfo.json
	fc.registryByName[name] = fInfo
	fc.registryByJSON[jsonName] = fInfo
	if fInfo.isComputedField() {
		if fInfo.stored {
			fc.computedStoredFields = append(fc.computedStoredFields, fInfo)
		} else {
			fc.computedFields = append(fc.computedFields, fInfo)
		}
	}
	if fInfo.isRelatedField() {
		fc.relatedFields = append(fc.relatedFields, fInfo)
	}
}

// Field holds the meta information about a field
type Field struct {
	model            *Model
	name             string
	json             string
	description      string
	help             string
	stored           bool
	required         bool
	readOnly         bool
	requiredFunc     func(Environment) (bool, Conditioner)
	readOnlyFunc     func(Environment) (bool, Conditioner)
	invisibleFunc    func(Environment) (bool, Conditioner)
	unique           bool
	index            bool
	compute          string
	depends          []string
	RelatedModelName string
	RelatedModel     *Model
	reverseFK        string
	jsonReverseFK    string
	m2mRelModel      *Model
	m2mOurField      *Field
	m2mTheirField    *Field
	selection        types.Selection
	selectionFunc    func() types.Selection
	FieldType        fieldtype.Type
	groupOperator    string
	size             int
	digits           nbutils.Digits
	structField      reflect.StructField
	relatedPathStr   string
	relatedPath      FieldName
	dependencies     []computeData
	embed            bool
	noCopy           bool
	defaultFunc      func(Environment) interface{}
	onDelete         OnDeleteAction
	onChange         string
	onChangeWarning  string
	onChangeFilters  string
	constraint       string
	inverse          string
	filter           *Condition
	contexts         FieldContexts
	ctxType          ctxType
	updates          []map[string]interface{}
}

// isComputedField returns true if this field is computed
func (f *Field) isComputedField() bool {
	return f.compute != ""
}

// isComputedField returns true if this field is related
func (f *Field) isRelatedField() bool {
	return f.relatedPath != nil
}

// isRelationField returns true if this field points to another model
func (f *Field) isRelationField() bool {
	// We check on relatedModelName and not relatedModel to be able
	// to use this method even if the models have not been bootstrapped yet.
	return f.RelatedModelName != ""
}

// isStored returns true if this field is stored in database
func (f *Field) isStored() bool {
	if f.FieldType.IsNonStoredRelationType() {
		// reverse fields are not stored
		return false
	}
	if (f.isComputedField() || f.isRelatedField()) && !f.stored {
		// Computed and related non stored fields are not stored
		return false
	}
	return true
}

// isSettable returns true if the given field can be set directly
func (f *Field) isSettable() bool {
	if f.isComputedField() && f.inverse == "" {
		return false
	}
	return true
}

// isReadOnly returns true if this field must not be set directly
// by the user.
func (f *Field) isReadOnly() bool {
	if f.readOnly {
		return true
	}
	fInfo := f
	//if fInfo.isRelatedField() {
	//	fInfo = f.model.getRelatedFieldInfo(fInfo.relatedPath)
	//}
	if fInfo.compute != "" && fInfo.inverse == "" {
		return true
	}
	return false
}

// isContextedField returns true if the value of this field depends on contexts
func (f *Field) isContextedField() bool {
	if f.contexts != nil && len(f.contexts) > 0 {
		return true
	}
	return false
}

// JSON returns this field name as FieldName type
func (f *Field) JSON() string {
	return f.json
}

// Name returns the field's name.
func (f *Field) Name() string {
	return f.name
}

var _ FieldName = new(Field)

// checkFieldInfo makes sanity checks on the given Field.
// It panics in case of severe error and logs recoverable errors.
func checkFieldInfo(fi *Field) {
	if fi.FieldType.IsReverseRelationType() && fi.reverseFK == "" {
		log.Panic("'one2many' and 'rev2one' fields must define a 'ReverseFK' parameter", "model",
			fi.model.TableName(), "field", fi.name, "type", fi.FieldType)
	}

	if fi.embed && !fi.FieldType.IsFKRelationType() {
		log.Warn("'Embed' should be set only on many2one or one2one fields", "model", fi.model.TableName(), "field", fi.name,
			"type", fi.FieldType)
		fi.embed = false
	}

	if fi.structField.Type == reflect.TypeOf(RecordCollection{}) && fi.RelatedModel.TableName() == "" {
		log.Panic("Undefined relation model on related field", "model", fi.model.TableName(), "field", fi.name,
			"type", fi.FieldType)
	}

	if fi.stored && !fi.isComputedField() {
		log.Warn("'stored' should be set only on computed fields", "model", fi.model.TableName(), "field", fi.name,
			"type", fi.FieldType)
		fi.stored = false
	}
}

// SnakeCaseFieldName returns a snake cased field name, adding '_id' on x2one
// relation fields and '_ids' to x2many relation fields.
func SnakeCaseFieldName(fName string, typ fieldtype.Type) string {
	res := strutils.SnakeCase(fName)
	if typ.Is2OneRelationType() {
		res += "_id"
	} else if typ.Is2ManyRelationType() {
		res += "_ids"
	}
	return res
}

// checkMethType panics if the given method does not have
// the correct number and type of arguments and returns for a compute/onChange method
func checkMethType(method *Method, label string) error {
	methType := method.methodType
	var msg string
	switch {
	case methType.NumIn() != 1:
		msg = fmt.Sprintf("%s should have no arguments", label)
	case methType.NumOut() == 0:
		msg = fmt.Sprintf("%s should return a value", label)
	case methType.NumOut() > 1:
		msg = fmt.Sprintf("Too many return values for %s", label)
	case !methType.Out(0).Implements(reflect.TypeOf((*RecordData)(nil)).Elem()):
		msg = fmt.Sprintf("%s returned value must implement models.RecordData", label)
	}
	if msg != "" {
		return errors.New(msg)
	}
	return nil
}

// checkOnChangeWarningType panics if the given method does not have
// the correct number and type of arguments and returns for a onChangeWarning method
func checkOnChangeWarningType(method *Method) error {
	methType := method.methodType
	var msg string
	switch {
	case methType.NumIn() != 1:
		msg = "OnChangeWarning methods should have no arguments"
	case methType.NumOut() == 0:
		msg = "OnChangeWarning methods should return a value"
	case methType.NumOut() > 1:
		msg = "Too many return values for OnChangeWarning method"
	case methType.Out(0) != reflect.TypeOf("string"):
		msg = "OnChangeWarning methods returned value must be of type string"
	}
	if msg != "" {
		return errors.New(msg)
	}
	return nil
}

// checkOnChangeFiltersType panics if the given method does not have
// the correct number and type of arguments and returns for a onChangeFilters method
func checkOnChangeFiltersType(method *Method) error {
	methType := method.methodType
	var msg string
	switch {
	case methType.NumIn() != 1:
		msg = "OnChangeFilters methods should have no arguments"
	case methType.NumOut() == 0:
		msg = "OnChangeFilters methods should return a value"
	case methType.NumOut() > 1:
		msg = "Too many return values for OnChangeFilters method"
	case methType.Out(0) != reflect.TypeOf(map[FieldName]Conditioner{}):
		msg = "OnChangeFilters methods returned value must be of type map[models.FieldName]models.Conditioner"
	}
	if msg != "" {
		return errors.New(msg)
	}
	return nil
}
