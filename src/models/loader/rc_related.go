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
	"github.com/hexya-erp/hexya/src/models/conditions"
	"github.com/hexya-erp/hexya/src/models/fieldtype"
	"github.com/hexya-erp/hexya/src/tools/strutils"
)

// substituteRelatedFields returns a copy of the given fields slice with related fields substituted by their related
// field path. It also adds the fk and pk fields of all records in the related paths.
//
// The second returned value is a map the keys of which are the related field paths, and the values are the
// corresponding original fields if they exist.
//
// This method removes duplicates and change all field names to their json names.
func (rc *RecordCollection) substituteRelatedFields(fields []conditions.FieldName) ([]conditions.FieldName, map[string]string) {
	resFields := make([]conditions.FieldName, len(fields))
	resSubsts := make(map[string]string)
	for i, field := range fields {
		relPath := rc.substituteRelatedInPath(field)
		resSubsts[relPath.JSON()] = field.JSON()
		resFields[i] = relPath
	}
	resFields = rc.addIntermediatePaths(resFields)
	return resFields, resSubsts
}

// addIntermediatePaths adds the paths that compose fields and returns a New slice.
//
// e.g. given [User.Address.Country Note Partner.Age] will return
// [User User.Address User.address.country Note Partner Partner.Age]
//
// This method removes duplicates
func (rc *RecordCollection) addIntermediatePaths(fields []conditions.FieldName) []conditions.FieldName {
	// Create a keys map with our fields to avoid duplicates
	keys := make(map[conditions.FieldName]bool)
	// Add intermediate records to our map
	for _, field := range fields {
		keys[field] = true
		exprs := conditions.SplitFieldNames(field, conditions.ExprSep)
		if len(exprs) == 1 {
			continue
		}
		var curPath conditions.FieldName
		for _, expr := range exprs {
			if curPath != nil {
				curPath = conditions.JoinFieldNames(append([]conditions.FieldName{curPath}, expr), conditions.ExprSep)
			} else {
				curPath = expr
			}
			keys[curPath] = true
		}
	}
	// Extract keys from our map to res
	res := make([]conditions.FieldName, len(keys))
	var i int
	for key := range keys {
		res[i] = key
		i++
	}
	return res
}

// substituteRelatedFieldsInMap returns a copy of the given FieldMap with related fields
// substituted by their related field path.
//
// This method substitute the first level only (to work with data structs)
func (rc *RecordCollection) substituteRelatedFieldsInMap(fMap FieldMap) FieldMap {
	res := make(FieldMap)
	for field, value := range fMap {
		// Inflate our related fields
		path := rc.substituteRelatedInPath(rc.model.FieldName(field))
		res[path.JSON()] = value
	}
	return res
}

// substituteRelatedInQuery returns a New RecordCollection with related fields
// substituted in the query.
func (rc *RecordCollection) substituteRelatedInQuery() *RecordCollection {
	// Substitute in RecordCollection query
	substs := make(map[conditions.FieldName][]conditions.FieldName)
	queryExprs := rc.query.getAllExpressions()
	for _, exprs := range queryExprs {
		if len(exprs) == 0 {
			continue
		}
		var curPath conditions.FieldName
		var resExprs []conditions.FieldName
		for _, expr := range exprs {
			resExprs = append(resExprs, expr)
			curPath = conditions.JoinFieldNames(resExprs, conditions.ExprSep)
			fi := rc.model.GetRelatedFieldInfo(curPath)
			curFI := fi
			for curFI.isRelatedField() {
				// We loop because target field may be related itself
				reLen := len(resExprs)
				jsonPath := curFI.relatedPath
				resExprs = append(resExprs[:reLen-1], conditions.SplitFieldNames(jsonPath, conditions.ExprSep)...)
				curFI = rc.model.GetRelatedFieldInfo(conditions.JoinFieldNames(resExprs, conditions.ExprSep))
			}
		}
		substs[conditions.JoinFieldNames(exprs, conditions.ExprSep)] = resExprs
	}
	rc.query.substituteConditionExprs(substs)

	return rc
}

// substituteRelatedInPath recursively substitutes path for its related value.
// If path is not a related field, it is returned as is.
func (rc *RecordCollection) substituteRelatedInPath(path conditions.FieldName) conditions.FieldName {
	exprs := conditions.SplitFieldNames(path, conditions.ExprSep)
	prefix := exprs[0]
	fi := rc.model.GetRelatedFieldInfo(prefix)
	if fi.isRelatedField() {
		newPath := fi.relatedPath
		if len(exprs) > 1 {
			newPath = conditions.JoinFieldNames(append([]conditions.FieldName{newPath}, exprs[1:]...), conditions.ExprSep)
		}
		return rc.substituteRelatedInPath(newPath)
	}
	if len(exprs) == 1 {
		return prefix
	}
	suffix := conditions.JoinFieldNames(exprs[1:], conditions.ExprSep)
	model := rc.Model().getRelatedModelInfo(prefix)
	res := conditions.JoinFieldNames(append([]conditions.FieldName{prefix}, rc.Env().(Environment).Pool(model.name).substituteRelatedInPath(suffix)), conditions.ExprSep)
	return res
}

// createRelatedRecord creates Records at the given path, starting from this recordset.
// This method does not check whether such a records already exists or not.
func (rc *RecordCollection) createRelatedRecord(path conditions.FieldName, vals RecordData) *RecordCollection {
	log.Debug("Creating related record", "recordset", rc, "path", path, "vals", vals)
	rc.EnsureOne()
	fi := rc.model.GetRelatedFieldInfo(path)
	exprs := conditions.SplitFieldNames(path, conditions.ExprSep)
	switch fi.FieldType {
	case fieldtype.Many2One, fieldtype.One2One, fieldtype.Many2Many:
		resRS := rc.createRelatedFKRecord(fi, vals)
		rc.Set(path, resRS.Collection())
		return resRS.Collection().(*RecordCollection)
	case fieldtype.One2Many, fieldtype.Rev2One:
		target := rc
		if len(exprs) > 1 {
			target = rc.Get(conditions.JoinFieldNames(exprs[:len(exprs)-1], conditions.ExprSep)).(conditions.RecordSet).Collection().(*RecordCollection)
			if target.IsEmpty() {
				log.Panic("Target record does not exist", "recordset", rc, "path", conditions.JoinFieldNames(exprs[:len(exprs)-1], conditions.ExprSep))
			}
			target = target.Records()[0]
		}
		vals.Underlying().Set(fi.RelatedModel.FieldName(fi.ReverseFK), target)
		return rc.env.Pool(fi.RelatedModel.name).Call("Create", vals).(conditions.RecordSet).Collection().(*RecordCollection)
	}
	return rc.env.Pool(rc.ModelName())
}

// createRelatedFKRecord creates a single related record for the given FK field
func (rc *RecordCollection) createRelatedFKRecord(fi *Field, data RecordData) *RecordCollection {
	rSet := rc.env.Pool(fi.RelatedModel.name)
	if fi.embed {
		rSet = rSet.WithContext("default_hexya_external_id", fmt.Sprintf("%s_%s", rc.Get(rc.model.FieldName("HexyaExternalID")), strutils.SnakeCase(fi.RelatedModel.name)))
	}
	res := rSet.Call("Create", data)
	return res.(conditions.RecordSet).Collection().(*RecordCollection)
}
