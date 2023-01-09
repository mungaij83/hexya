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
	"fmt"
	"reflect"

	"github.com/hexya-erp/hexya/src/models/operator"
)

// A ConditionPredicate of a condition in the form 'Field = arg'
type ConditionPredicate struct {
	Exprs        []FieldName
	CondOperator operator.Operator
	Arg          interface{}
	Cond         *Condition
	IsOr         bool
	IsNot        bool
	IsCond       bool
}

// Field returns the field name of this ConditionPredicate
func (p ConditionPredicate) Field() FieldName {
	return JoinFieldNames(p.Exprs, ExprSep)
}

// Operator returns the CondOperator of this ConditionPredicate
func (p ConditionPredicate) Operator() operator.Operator {
	return p.CondOperator
}

// Argument returns the argument of this ConditionPredicate
func (p ConditionPredicate) Argument() interface{} {
	return p.Arg
}

// AlterField changes the field of this ConditionPredicate
func (p *ConditionPredicate) AlterField(f FieldName) *ConditionPredicate {
	if f == nil || f.Name() == "" {
		log.Panic("AlterField must be called with a field name", "field", f)
	}
	p.Exprs = SplitFieldNames(f, ExprSep)
	return p
}

// AlterOperator changes the CondOperator of this ConditionPredicate
func (p *ConditionPredicate) AlterOperator(op operator.Operator) *ConditionPredicate {
	p.CondOperator = op
	return p
}

// AlterArgument changes the argument of this ConditionPredicate
func (p *ConditionPredicate) AlterArgument(arg interface{}) *ConditionPredicate {
	p.Arg = arg
	return p
}

// A Condition represents a WHERE clause of an SQL query.
type Condition struct {
	Predicates []ConditionPredicate
}

// newCondition returns a new condition struct
func NewCondition() *Condition {
	c := &Condition{}
	return c
}

// And completes the current condition with a simple AND clause : c.And().nextCond => c AND nextCond.
//
// No brackets are added so AND precedence over OR applies.
func (c Condition) And() *ConditionStart {
	res := ConditionStart{cond: c}
	return &res
}

// AndCond completes the current condition with the given cond as an AND clause
// between brackets : c.And(cond) => (c) AND (cond)
func (c Condition) AndCond(cond *Condition) *Condition {
	if !cond.IsEmpty() {
		c.Predicates = append(c.Predicates, ConditionPredicate{Cond: cond, IsCond: true})
	}
	return &c
}

// AndNot completes the current condition with a simple AND NOT clause :
// c.AndNot().nextCond => c AND NOT nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c Condition) AndNot() *ConditionStart {
	res := ConditionStart{cond: c}
	res.nextIsNot = true
	return &res
}

// AndNotCond completes the current condition with an AND NOT clause between
// brackets : c.AndNot(cond) => (c) AND NOT (cond)
func (c Condition) AndNotCond(cond *Condition) *Condition {
	if !cond.IsEmpty() {
		c.Predicates = append(c.Predicates, ConditionPredicate{Cond: cond, IsCond: true, IsNot: true})
	}
	return &c
}

// Or completes the current condition both with a simple OR clause : c.Or().nextCond => c OR nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c Condition) Or() *ConditionStart {
	res := ConditionStart{cond: c}
	res.nextIsOr = true
	return &res
}

// OrCond completes the current condition both with an OR clause between
// brackets : c.Or(cond) => (c) OR (cond)
func (c Condition) OrCond(cond *Condition) *Condition {
	if !cond.IsEmpty() {
		c.Predicates = append(c.Predicates, ConditionPredicate{Cond: cond, IsCond: true, IsOr: true})
	}
	return &c
}

// OrNot completes the current condition both with a simple OR NOT clause : c.OrNot().nextCond => c OR NOT nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c Condition) OrNot() *ConditionStart {
	res := ConditionStart{cond: c}
	res.nextIsNot = true
	res.nextIsOr = true
	return &res
}

// OrNotCond completes the current condition both with an OR NOT clause between
// brackets : c.OrNot(cond) => (c) OR NOT (cond)
func (c Condition) OrNotCond(cond *Condition) *Condition {
	if !cond.IsEmpty() {
		c.Predicates = append(c.Predicates, ConditionPredicate{Cond: cond, IsCond: true, IsOr: true, IsNot: true})
	}
	return &c
}

// Serialize returns the condition as a list which mimics Odoo domains.
func (c Condition) Serialize() []interface{} {
	return serializePredicates(c.Predicates)
}

// HasField returns true if the given field is in at least one of the
// the predicates of this condition or of one of its nested conditions.
func (c Condition) HasField(jsonName string) bool {
	preds := c.PredicatesWithField(jsonName)
	return len(preds) > 0
}

// PredicatesWithField returns all predicates of this condition (including
// nested conditions) that concern the given field.
func (c Condition) PredicatesWithField(jsonName string) []*ConditionPredicate {
	var res []*ConditionPredicate
	for i, pred := range c.Predicates {
		if len(pred.Exprs) > 0 {
			if JoinFieldNames(pred.Exprs, ExprSep).JSON() == jsonName {
				res = append(res, &c.Predicates[i])
			}
		}
		if pred.Cond != nil {
			res = append(res, c.Predicates[i].Cond.PredicatesWithField(jsonName)...)
		}
	}
	return res
}

// String method for the Condition. Recursively print all predicates.
func (c Condition) String() string {
	var res string
	for _, p := range c.Predicates {
		if p.IsOr {
			res += "OR "
		} else {
			res += "AND "
		}
		if p.IsNot {
			res += "NOT "
		}
		if p.IsCond {
			res += fmt.Sprintf("(\n%s\n)\n", p.Cond.String())
			continue
		}
		res += fmt.Sprintf("%s %s %v\n", JoinFieldNames(p.Exprs, ExprSep).Name(), p.CondOperator, p.Arg)
	}
	return res
}

// Underlying returns the underlying Condition (i.e. itself)
func (c Condition) Underlying() *Condition {
	return &c
}

var _ Conditioner = Condition{}

// A ConditionStart is an object representing a Condition when
// we just added a logical CondOperator (AND, OR, ...) and we are
// about to add a ConditionPredicate.
type ConditionStart struct {
	cond      Condition
	nextIsOr  bool
	nextIsNot bool
}

func (cs ConditionStart) FieldName(name string) *ConditionField {
	return cs.Field(NewFieldName(name, name))
}

// Field adds a field path (dot separated) to this condition
func (cs ConditionStart) Field(name FieldName) *ConditionField {
	newExprs := SplitFieldNames(name, ExprSep)
	cp := ConditionField{cs: cs}
	cp.Exprs = append(cp.Exprs, newExprs...)
	return &cp
}

// FilteredOn adds a condition with a table join on the given field and
// filters the result with the given condition
func (cs ConditionStart) FilteredOn(field FieldName, condition *Condition) *Condition {
	res := cs.cond
	for i, p := range condition.Predicates {
		condition.Predicates[i].Exprs = append([]FieldName{field}, p.Exprs...)
	}
	condition.Predicates[0].IsOr = cs.nextIsOr
	condition.Predicates[0].IsNot = cs.nextIsNot
	res.Predicates = append(res.Predicates, condition.Predicates...)
	return &res
}

var _ FieldName = ConditionField{}

// AddOperator adds a condition value to the condition with the given CondOperator and data
// If multi is true, a recordset will be converted into a slice of int64
// otherwise, it will return an int64 and panic if the recordset is not
// a singleton.
//
// This method is low level and should be avoided. Use CondOperator methods such as Equals()
// instead.
func (c ConditionField) AddOperator(op operator.Operator, data interface{}) *Condition {
	cond := c.cs.cond
	data = SanitizeArgs(data, op.IsMulti())
	if data != nil && op.IsMulti() && reflect.ValueOf(data).Kind() == reflect.Slice && reflect.ValueOf(data).Len() == 0 {
		// field in [] => ID = -1
		cond.Predicates = []ConditionPredicate{{
			Exprs:        []FieldName{ID},
			CondOperator: operator.Equals,
			Arg:          -1,
		}}
		return &cond
	}
	cond.Predicates = append(cond.Predicates, ConditionPredicate{
		Exprs:        c.Exprs,
		CondOperator: op,
		Arg:          data,
		IsNot:        c.cs.nextIsNot,
		IsOr:         c.cs.nextIsOr,
	})
	return &cond
}

// SanitizeArgs returns the given args suitable for SQL query
// In particular, retrieves the ids of a recordset if args is one.
// If multi is true, a recordset will be converted into a slice of int64
// otherwise, it will return an int64 and panic if the recordset is not
// a singleton
func SanitizeArgs(args interface{}, multi bool) interface{} {
	if rs, ok := args.(RecordSet); ok {
		if multi {
			return rs.Ids()
		}
		if len(rs.Ids()) > 1 {
			log.Panic("Trying to extract a single ID from a non singleton", "args", args)
		}
		if len(rs.Ids()) == 0 {
			return nil
		}
		return rs.Ids()[0]
	}
	return args
}

// Equals appends the '=' CondOperator to the current Condition
func (c ConditionField) Equals(data interface{}) *Condition {
	return c.AddOperator(operator.Equals, data)
}

// NotEquals appends the '!=' CondOperator to the current Condition
func (c ConditionField) NotEquals(data interface{}) *Condition {
	return c.AddOperator(operator.NotEquals, data)
}

// Greater appends the '>' CondOperator to the current Condition
func (c ConditionField) Greater(data interface{}) *Condition {
	return c.AddOperator(operator.Greater, data)
}

// GreaterOrEqual appends the '>=' CondOperator to the current Condition
func (c ConditionField) GreaterOrEqual(data interface{}) *Condition {
	return c.AddOperator(operator.GreaterOrEqual, data)
}

// Lower appends the '<' CondOperator to the current Condition
func (c ConditionField) Lower(data interface{}) *Condition {
	return c.AddOperator(operator.Lower, data)
}

// LowerOrEqual appends the '<=' CondOperator to the current Condition
func (c ConditionField) LowerOrEqual(data interface{}) *Condition {
	return c.AddOperator(operator.LowerOrEqual, data)
}

// Like appends the 'LIKE' CondOperator to the current Condition
func (c ConditionField) Like(data interface{}) *Condition {
	return c.AddOperator(operator.Like, data)
}

// ILike appends the 'ILIKE' CondOperator to the current Condition
func (c ConditionField) ILike(data interface{}) *Condition {
	return c.AddOperator(operator.ILike, data)
}

// Contains appends the 'LIKE %%' CondOperator to the current Condition
func (c ConditionField) Contains(data interface{}) *Condition {
	return c.AddOperator(operator.Contains, data)
}

// NotContains appends the 'NOT LIKE %%' CondOperator to the current Condition
func (c ConditionField) NotContains(data interface{}) *Condition {
	return c.AddOperator(operator.NotContains, data)
}

// IContains appends the 'ILIKE %%' CondOperator to the current Condition
func (c ConditionField) IContains(data interface{}) *Condition {
	return c.AddOperator(operator.IContains, data)
}

// NotIContains appends the 'NOT ILIKE %%' CondOperator to the current Condition
func (c ConditionField) NotIContains(data interface{}) *Condition {
	return c.AddOperator(operator.NotIContains, data)
}

// In appends the 'IN' CondOperator to the current Condition
func (c ConditionField) In(data interface{}) *Condition {
	return c.AddOperator(operator.In, data)
}

// NotIn appends the 'NOT IN' CondOperator to the current Condition
func (c ConditionField) NotIn(data interface{}) *Condition {
	return c.AddOperator(operator.NotIn, data)
}

// ChildOf appends the 'child of' CondOperator to the current Condition
func (c ConditionField) ChildOf(data interface{}) *Condition {
	return c.AddOperator(operator.ChildOf, data)
}

// IsNull checks if the current condition field is null
func (c ConditionField) IsNull() *Condition {
	return c.AddOperator(operator.Equals, nil)
}

// IsNotNull checks if the current condition field is not null
func (c ConditionField) IsNotNull() *Condition {
	return c.AddOperator(operator.NotEquals, nil)
}

// IsEmpty check the condition arguments are empty or not.
func (c *Condition) IsEmpty() bool {
	switch {
	case c == nil:
		return true
	case len(c.Predicates) == 0:
		return true
	case len(c.Predicates) == 1 && c.Predicates[0].Cond != nil && c.Predicates[0].Cond.IsEmpty():
		return true
	}
	return false
}

// // GetAllExpressions returns a list of all exprs used in this condition,
// // and recursively in all subconditions.
// // Expressions are given in field json format
func (c Condition) GetAllExpressions() [][]FieldName {
	var res [][]FieldName
	for _, p := range c.Predicates {
		res = append(res, p.Exprs)
		if p.Cond != nil {
			res = append(res, p.Cond.GetAllExpressions()...)
		}
	}
	return res
}

// // SubstituteExprs recursively replaces condition exprs that match substs keys
// // with the corresponding substs values.
func (c *Condition) SubstituteExprs(substs map[FieldName][]FieldName) {
	for i, p := range c.Predicates {
		for k, v := range substs {
			if len(p.Exprs) > 0 && JoinFieldNames(p.Exprs, ExprSep) == k {
				c.Predicates[i].Exprs = v
			}
		}
		if p.Cond != nil {
			p.Cond.SubstituteExprs(substs)
		}
	}
}

// // SubstituteChildOfOperator recursively replaces in the condition the
// // predicates with ChildOf CondOperator by the predicates to actually execute.
func (c *Condition) SubstituteChildOfOperator(hasParentFunc func(name FieldName, args interface{}) (bool, []int64)) {
	for i, p := range c.Predicates {
		if p.Cond != nil {
			p.Cond.SubstituteChildOfOperator(hasParentFunc)
		}
		if p.CondOperator != operator.ChildOf {
			continue
		}
		hasParent, parentIds := hasParentFunc(JoinFieldNames(p.Exprs, ExprSep), p.Arg)
		if !hasParent {
			// If we have no parent field, then we fetch only the "parent" record
			c.Predicates[i].CondOperator = operator.Equals
			continue
		}
		c.Predicates[i].CondOperator = operator.In
		c.Predicates[i].Arg = parentIds
	}
}
