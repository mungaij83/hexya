package loader

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/hexya-erp/hexya/src/models/conditions"
	"github.com/hexya-erp/hexya/src/models/fieldtype"
	"reflect"
	"strconv"
	"time"
)

// FieldContexts define the different contexts for a field, that will define different
// values for this field.
//
// The key is a context name and the value is a function that returns the context
// value for the given recordset.
type FieldContexts map[string]func(conditions.RecordSet) string

// A FieldMapper is an object that can convert itself into a FieldMap
type FieldMapper interface {
	// Underlying returns the object converted to a FieldMap.
	Underlying() FieldMap
}

// A Methoder can return a Method data object through its Underlying() method
type Methoder interface {
	Underlying() *Method
}

// A Modeler can return a Model data object through its Underlying() method
type Modeler interface {
	Underlying() *Model
}

// A GroupAggregateRow holds a row of results of a query with a Group by clause
// - Values holds the values of the actual query
// - Count is the number of lines aggregated into this one
// - Condition can be used to query the aggregated rows separately if needed
type GroupAggregateRow struct {
	Values    *ModelData
	Count     int
	Condition *conditions.Condition
}

// A RecordData can return a ModelData object through its Underlying() method
type RecordData interface {
	sql.Scanner
	Underlying() *ModelData
}

// A ModelData is used to hold values of an object instance for creating or
// updating a RecordSet. It is mainly designed to be embedded in a type-safe
// struct.
type ModelData struct {
	FieldMap
	ToCreate map[string][]*ModelData
	Model    *Model
}

var _ RecordData = new(ModelData)

// Scan implements sql.Scanner
func (md *ModelData) Scan(src interface{}) error {
	switch val := src.(type) {
	case nil:
		return nil
	case FieldMapper:
		md.FieldMap = val.Underlying()
	case map[string]interface{}:
		md.FieldMap = val
	default:
		return fmt.Errorf("unexpected type %T to represent RecordData: %s", src, src)
	}
	return nil
}

// Get returns the value of the given field.
//
// The field can be either its name or is JSON name.
func (md *ModelData) Get(field conditions.FieldName) interface{} {
	res, _ := md.FieldMap.Get(field)
	return res
}

// The field can be either its name or is JSON name.
func (md *ModelData) GetField(field string) interface{} {
	res, _ := md.FieldMap.Get(conditions.NewFieldName(field, ""))
	return res
}

func (md *ModelData) GetJsonField(field string) interface{} {
	res, _ := md.FieldMap.Get(conditions.NewFieldName("", field))
	return res
}

func (md *ModelData) GetDateTime(field string) *time.Time {
	res, ok := md.FieldMap.Get(conditions.NewFieldName(field, ""))
	if ok {
		return nil
	}
	v, ok := res.(time.Time)
	if ok {
		return &v
	}
	return nil
}

//func (md *ModelData) GetData[M any](field string) M {
//	res := md.GetField(field)
//	if res == nil {
//		return nil
//	}
//	v, ok := res.(M)
//	if ok {
//		return v
//	}
//	return nil
//}

func (md *ModelData) GetInt64(field string) int64 {
	res := md.GetField(field)
	if res == nil {
		return 0
	}
	v, ok := res.(int64)
	if ok {
		return v
	}
	return 0
}

func (md *ModelData) GetString(field string) string {
	res := md.GetField(field)
	if res == nil {
		return ""
	}
	v, ok := res.(string)
	if ok {
		return v
	}
	return ""
}

// Has returns true if this ModelData has values for the given field.
//
// The field can be either its name or is JSON name.
func (md *ModelData) Has(field conditions.FieldName) bool {
	if _, ok := md.FieldMap.Get(field); ok {
		return true
	}
	if _, ok := md.ToCreate[field.JSON()]; ok {
		return true
	}
	return false
}

// Set sets the given field with the given value.
// If the field already exists, then it is updated with value.
// Otherwise, a new entry is inserted.
//
// It returns the given ModelData so that calls can be chained
func (md *ModelData) Set(field conditions.FieldName, value interface{}) *ModelData {
	md.FieldMap.Set(field, value)
	return md
}

// It returns the given ModelData so that calls can be chained
func (md *ModelData) SetValue(field string, value interface{}) *ModelData {
	md.FieldMap.Set(conditions.NewFieldName(field, ""), value)
	return md
}

// Unset removes the value of the given field if it exists.
//
// It returns the given ModelData so that calls can be chained
func (md *ModelData) Unset(field conditions.FieldName) *ModelData {
	md.FieldMap.Delete(field)
	delete(md.ToCreate, field.JSON())
	return md
}

// Create stores the related ModelData to be used to create
// a related record on the fly and link it to this field.
//
// This method can be called multiple times to create multiple records
func (md *ModelData) Create(field conditions.FieldName, related *ModelData) *ModelData {
	fi := md.Model.GetRelatedFieldInfo(field)
	if related.Model != fi.RelatedModel {
		log.Panic("create data must be of the model of the relation field", "fieldModel", fi.RelatedModel, "dataModel", related.Model)
	}
	md.ToCreate[field.JSON()] = append(md.ToCreate[field.JSON()], related)
	return md
}

// Copy returns a copy of this ModelData
func (md *ModelData) Copy() *ModelData {
	ntc := make(map[string][]*ModelData)
	for k, v := range md.ToCreate {
		ntc[k] = v
	}
	return &ModelData{
		Model:    md.Model,
		FieldMap: md.FieldMap.Copy(),
		ToCreate: ntc,
	}
}

// MergeWith updates this ModelData with the given other ModelData.
// If a key of the other ModelData already exists here, the value is overridden,
// otherwise, the key is inserted with its json name.
func (md *ModelData) MergeWith(other *ModelData) {
	// 1. We unset all entries existing in other to remove both FieldMap and ToCreate entries
	for field := range other.FieldMap {
		if md.Has(md.Model.FieldName(field)) {
			md.Unset(md.Model.FieldName(field))
		}
	}
	for field := range other.ToCreate {
		if md.Has(md.Model.FieldName(field)) {
			md.Unset(md.Model.FieldName(field))
		}
	}
	// 2. We set other values in md
	md.FieldMap.MergeWith(other.FieldMap, other.Model)
	for field, toCreate := range other.ToCreate {
		md.ToCreate[field] = append(md.ToCreate[field], toCreate...)
	}
}

// FieldNames returns the ModelData keys as a slice of FieldNames.
func (md *ModelData) FieldNames() conditions.FieldNames {
	return md.FieldMap.FieldNames(md.Model)
}

// MarshalJSON function for ModelData. Returns the FieldMap.
func (md *ModelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(md.FieldMap)
}

// Underlying returns the ModelData
func (md *ModelData) Underlying() *ModelData {
	return md
}

// fixFieldValue changes the given value for the given field by applying several fixes
func fixFieldValue(v interface{}, fi *Field) interface{} {
	if _, ok := v.(bool); ok && fi.FieldType != fieldtype.Boolean {
		// Client returns false when empty
		v = reflect.Zero(fi.structField.Type).Interface()
	}
	if _, ok := v.([]byte); ok && fi.FieldType == fieldtype.Float {
		// DB can return numeric types as []byte
		switch fi.structField.Type.Kind() {
		case reflect.Float64:
			if res, err := strconv.ParseFloat(string(v.([]byte)), 64); err == nil {
				v = res
			}
		case reflect.Float32:
			if res, err := strconv.ParseFloat(string(v.([]byte)), 32); err == nil {
				v = float32(res)
			}
		}
	}
	if _, ok := v.(float64); ok && fi.FieldType == fieldtype.Integer {
		// JSON unmarshals int to float64. Convert back to the Go type of fi.
		val := reflect.ValueOf(v)
		typ := fi.structField.Type
		val = val.Convert(typ)
		v = val.Interface()
	}
	return v
}

// NewModelData returns a pointer to a new instance of ModelData
// for the given model. If FieldMaps are given they are added to
// the ModelData.
func NewModelData(model *Model, fm ...FieldMap) *ModelData {
	fMap := make(FieldMap)
	for _, f := range fm {
		for k, v := range f {
			fi := model.Underlying().GetRelatedFieldInfo(model.FieldName(k))
			v = fixFieldValue(v, fi)
			fMap[fi.json] = v
		}
	}
	return &ModelData{
		FieldMap: fMap,
		ToCreate: make(map[string][]*ModelData),
		Model:    model,
	}
}

// NewModelDataFromRS creates a pointer to a new instance of ModelData.
// If FieldMaps are given they are added to the ModelData.
//
// Unlike NewModelData, this method translates relation fields in64 and
// []int64 values as RecordSets
func NewModelDataFromRS(rs conditions.RecordSet, fm ...FieldMap) *ModelData {
	fMap := make(FieldMap)
	for _, f := range fm {
		for k, v := range f {
			fi := rs.Collection().(*RecordCollection).Model().GetRelatedFieldInfo(rs.Collection().(*RecordCollection).Model().FieldName(k))
			if fi.isRelatedField() {
				v = rs.Collection().(*RecordCollection).ConvertToRecordSet(v, fi.RelatedModelName)
			}
			v = fixFieldValue(v, fi)
			fMap[fi.json] = v
		}
	}
	return &ModelData{
		FieldMap: fMap,
		ToCreate: make(map[string][]*ModelData),
		Model:    rs.Collection().(*RecordCollection).model,
	}
}
