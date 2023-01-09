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
	"github.com/hexya-erp/hexya/src/models/conditions"
	"strings"
)

// jsonizeExpr returns an expression slice with field names changed to the fields json names
// Computation is made relatively to the given Model
// e.g. [User Profile Name] -> [user_id profile_id name]
func jsonizeExpr(mi *Model, exprs []string) []string {
	if len(exprs) == 0 {
		return []string{}
	}
	var res []string
	fi := mi.fields.MustGet(exprs[0])
	res = append(res, fi.json)
	if len(exprs) > 1 {
		if fi.RelatedModel != nil {
			res = append(res, jsonizeExpr(fi.RelatedModel, exprs[1:])...)
		} else {
			log.Panic("Field is not a relation in model", "field", exprs[0], "model", mi.name)
		}
	}
	return res
}

// addNameSearchesToCondition recursively modifies the given condition to search
// on the name of the related records if they point to a relation field.
func addNameSearchesToCondition(mi *Model, cond *conditions.Condition) {
	// Implemented with callback function to remove cyclic dependencies between condition and loader modules
	conditions.AddNameSearchesToCondition(
		func(f conditions.FieldName) (bool, conditions.FieldName) {
			fi := mi.GetRelatedFieldInfo(f)
			if !fi.IsRelationField() {
				return false, nil
			}
			return true, conditions.NewFieldName(fi.name, fi.json)
		},
		func(s conditions.FieldName, names []conditions.FieldName) []conditions.FieldName {
			fi := mi.GetRelatedFieldInfo(s)
			return addNameSearchToExprs(fi, names)
		},
		cond,
	)
}

// addNameSearchToExprs modifies the given exprs to search on the name of the related record
// if it points to a relation field.
func addNameSearchToExprs(fi *Field, exprs []conditions.FieldName) []conditions.FieldName {
	relFI, exists := fi.RelatedModel.fields.Get("name")
	if !exists {
		return exprs
	}
	exprsToAppend := []conditions.FieldName{conditions.Name}
	if relFI.isRelatedField() {
		exprsToAppend = conditions.SplitFieldNames(relFI.relatedPath, conditions.ExprSep)
	}
	exprs = append(exprs, exprsToAppend...)
	return exprs
}

// jsonizePath returns a path with field names changed to the field json names
// Computation is made relatively to the given Model
// e.g. User.Profile.Name -> user_id.profile_id.name
func jsonizePath(mi *Model, path string) string {
	exprs := strings.Split(path, conditions.ExprSep)
	exprs = jsonizeExpr(mi, exprs)
	return strings.Join(exprs, conditions.ExprSep)
}

// filterOnDBFields returns the given fields slice with only stored fields.
// This function also adds the "id" field to the list if not present unless dontAddID is true
func filterOnDBFields(mi *Model, fields []conditions.FieldName, dontAddID ...bool) []conditions.FieldName {
	var res []conditions.FieldName
	// Check if fields are stored
	for _, field := range fields {
		fieldExprs := conditions.SplitFieldNames(field, conditions.ExprSep)
		fi := mi.fields.MustGet(fieldExprs[0].JSON())
		fn := mi.FieldName(fi.json)
		// Single field
		if len(fieldExprs) == 1 {
			if fi.isStored() {
				res = append(res, fn)
			}
			continue
		}

		// Depends field (e.g. User.Profile.Age)
		if fi.RelatedModel == nil {
			log.Panic("Field is not a relation in model", "field", fieldExprs[0], "model", mi.name)
		}
		subFieldName := conditions.JoinFieldNames(fieldExprs[1:], conditions.ExprSep)
		subFieldRes := filterOnDBFields(fi.RelatedModel, []conditions.FieldName{subFieldName}, dontAddID...)
		if len(subFieldRes) == 0 {
			// Our last expr is not stored after all, we don't add anything
			continue
		}

		if !fi.isStored() {
			// We re-add our first expr as it has been removed above (not stored)
			res = append(res, fn)
		}
		for _, sfr := range subFieldRes {
			resExprs := []conditions.FieldName{fn}
			resExprs = append(resExprs, sfr)
			res = append(res, conditions.JoinFieldNames(resExprs, conditions.ExprSep))
		}
	}
	if len(dontAddID) == 0 || !dontAddID[0] {
		res = AddIDIfNotPresent(res)
	}
	return res
}

// AddIDIfNotPresent returns a new fields slice including ID if it
// is not already present. Otherwise returns the original slice.
func AddIDIfNotPresent(fields []conditions.FieldName) []conditions.FieldName {
	var hadID bool
	for _, fName := range fields {
		if fName.JSON() == "id" {
			hadID = true
		}
	}
	if !hadID {
		fields = append(fields, conditions.ID)
	}
	return fields
}

// getGroupCondition returns the condition to retrieve the individual aggregated rows in vals
// knowing that they were grouped by Groups and that we had the given initial condition
func getGroupCondition(groups []conditions.FieldName, vals map[string]interface{}, initialCondition *conditions.Condition) *conditions.Condition {
	res := initialCondition
	for _, group := range groups {
		res = res.And().Field(group).Equals(vals[group.JSON()])
	}
	return res
}

// substituteKeys returns a new map with its keys substituted following substMap after changing sqlSep into ExprSep.
// vals keys that are not found in substMap are not returned
func substituteKeys(vals map[string]interface{}, substMap map[string]string) map[string]interface{} {
	res := make(FieldMap)
	for f, v := range vals {
		k := strings.Replace(f, conditions.SqlSep, conditions.ExprSep, -1)
		sk, ok := substMap[k]
		if !ok {
			continue
		}
		res[sk] = v
	}
	return res
}

// DefaultValue returns a function that is suitable for the Default parameter of
// model fields and that simply returns value.
func DefaultValue(value interface{}) func(env Environment) interface{} {
	return func(env Environment) interface{} {
		return value
	}
}

// cartesianProductSlices returns the cartesian product of the given RecordCollection slices.
//
// This function panics if all records are not pf the same model
func cartesianProductSlices(records ...[]*RecordCollection) []*RecordCollection {
	switch len(records) {
	case 0:
		return []*RecordCollection{}
	case 1:
		return records[0]
	case 2:
		res := make([]*RecordCollection, len(records[0])*len(records[1]))
		for i, v1 := range records[0] {
			for j, v2 := range records[1] {
				res[i*len(records[1])+j] = v1.Union(v2)
			}
		}
		return res
	default:
		return cartesianProductSlices(records[0], cartesianProductSlices(records[1:]...))
	}
}
