package conditions

import (
	"strings"
)

// JoinFieldNames returns a field name that is the join of fn with sep
func JoinFieldNames(fn []FieldName, sep string) FieldName {
	var ntoks, jtoks []string
	for _, f := range fn {
		ntoks = append(ntoks, f.Name())
		jtoks = append(jtoks, f.JSON())
	}
	return fieldName{
		name: strings.Join(ntoks, sep),
		json: strings.Join(jtoks, sep),
	}
}

// SplitFieldNames splits the field name at sep, returning the result as a slice
func SplitFieldNames(f FieldName, sep string) []FieldName {
	ntoks := strings.Split(f.Name(), sep)
	jtoks := strings.Split(f.JSON(), sep)
	if len(ntoks) != len(jtoks) {
		log.Panic("name and json paths lengths are inconsistent", "fieldName", f)
	}
	res := make([]FieldName, len(ntoks))
	for i := 0; i < len(ntoks); i++ {
		res[i] = fieldName{
			name: ntoks[i],
			json: jtoks[i],
		}
	}
	return res
}

// serializePredicates returns a list that mimics Odoo domains from the given
// condition predicates.
func serializePredicates(predicates []ConditionPredicate) []interface{} {
	var res []interface{}
	i := 0
	for i < len(predicates) {
		if predicates[i].IsOr {
			subRes := []interface{}{"|"}
			subRes = appendPredicateToSerial(subRes, predicates[i])
			subRes, i = consumeAndPredicates(i+1, predicates, subRes)
			res = append(subRes, res...)
		} else {
			res, i = consumeAndPredicates(i, predicates, res)
		}
	}
	return res
}

// consumeAndPredicates appends res with all successive AND predicates
// starting from position i and returns the next position as second argument.
func consumeAndPredicates(i int, predicates []ConditionPredicate, res []interface{}) ([]interface{}, int) {
	if i >= len(predicates) || predicates[i].IsOr {
		return res, i
	}
	j := i
	for j < len(predicates)-1 {
		if predicates[j+1].IsOr {
			break
		}
		j++
	}
	for k := i; k < j; k++ {
		res = append(res, "&")
		res = appendPredicateToSerial(res, predicates[k])
	}
	res = appendPredicateToSerial(res, predicates[j])
	return res, j + 1
}

// appendPredicateToSerial appends the given ConditionPredicate to the given serialized
// ConditionPredicate list and returns the result.
func appendPredicateToSerial(res []interface{}, predicate ConditionPredicate) []interface{} {
	if predicate.IsCond {
		res = append(res, serializePredicates(predicate.Cond.Predicates)...)
	} else {
		res = append(res, []interface{}{JoinFieldNames(predicate.Exprs, ExprSep).JSON(), predicate.CondOperator, predicate.Arg})
	}
	return res
}

// addNameSearchesToCondition recursively modifies the given condition to search
// on the name of the related records if they point to a relation field.
func AddNameSearchesToCondition(isRelationField func(f FieldName) (bool, FieldName), addNameSearchToExprs func(FieldName, []FieldName) []FieldName, cond *Condition) {
	for i, p := range cond.Predicates {
		if p.Cond != nil {
			AddNameSearchesToCondition(isRelationField, addNameSearchToExprs, p.Cond)
		}
		if len(p.Exprs) == 0 {
			continue
		}
		isRelated, name := isRelationField(JoinFieldNames(p.Exprs, ExprSep))
		if isRelated {
			continue
		}
		switch p.Arg.(type) {
		case bool:
			cond.Predicates[i].Arg = int64(0)
		case string:
			cond.Predicates[i].Exprs = addNameSearchToExprs(name, p.Exprs)
		}
	}
}
