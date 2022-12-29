// Copyright 2019 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package loader

import (
	"reflect"
)

// Wrap returns the given RecordCollection embedded into a RecordSet Wrapper type
//
// If modelName is defined, wrap in a modelName Wrapper type instead (use for mixins).
func (rc *RecordCollection) Wrap(modelName ...string) interface{} {
	modName := rc.ModelName()
	if len(modelName) > 0 {
		modName = modelName[0]
	}
	typ, ok := RecordSetWrappers[modName]
	if !ok {
		log.Panic("unable to wrap RecordCollection", "model", modName)
	}
	val := reflect.New(typ).Elem()
	val.Field(0).Set(reflect.ValueOf(rc))
	return val.Interface()
}

// recordSetWrappers is a map that stores the available types of ModelData
var modelDataWrappers map[string]reflect.Type

// Wrap returns the given ModelData embedded into a RecordSet Wrapper type.
// This method returns a pointer.
func (md ModelData) Wrap() interface{} {
	typ, ok := modelDataWrappers[md.Model.name]
	if !ok {
		return &md
	}
	val := reflect.New(typ)
	val.Elem().Field(0).Set(reflect.ValueOf(&md))
	return val.Interface()
}
