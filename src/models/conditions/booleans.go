package conditions

import (
	"github.com/hexya-erp/hexya/src/models/operator"
)

// A BoolConditionField is a partial ModelCondition when
// we have selected a field of type bool and expecting an CondOperator.
type BoolConditionField struct {
	*ConditionField
}

// Equals adds a condition value to the ConditionPath
func (c BoolConditionField) Equals(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) EqualsFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) EqualsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(expression),
	}
}

// NotEquals adds a condition value to the ConditionPath
func (c BoolConditionField) NotEquals(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) NotEqualsFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) NotEqualsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(expression),
	}
}

// Greater adds a condition value to the ConditionPath
func (c BoolConditionField) Greater(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) GreaterFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) GreaterEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(expression),
	}
}

// GreaterOrEqual adds a condition value to the ConditionPath
func (c BoolConditionField) GreaterOrEqual(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) GreaterOrEqualFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) GreaterOrEqualEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(expression),
	}
}

// Lower adds a condition value to the ConditionPath
func (c BoolConditionField) Lower(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) LowerFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) LowerEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(expression),
	}
}

// LowerOrEqual adds a condition value to the ConditionPath
func (c BoolConditionField) LowerOrEqual(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) LowerOrEqualFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) LowerOrEqualEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(expression),
	}
}

// Like adds a condition value to the ConditionPath
func (c BoolConditionField) Like(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) LikeFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) LikeEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(expression),
	}
}

// Contains adds a condition value to the ConditionPath
func (c BoolConditionField) Contains(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) ContainsFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) ContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(expression),
	}
}

// NotContains adds a condition value to the ConditionPath
func (c BoolConditionField) NotContains(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) NotContainsFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) NotContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(expression),
	}
}

// IContains adds a condition value to the ConditionPath
func (c BoolConditionField) IContains(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) IContainsFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) IContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(expression),
	}
}

// NotIContains adds a condition value to the ConditionPath
func (c BoolConditionField) NotIContains(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) NotIContainsFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) NotIContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(expression),
	}
}

// ILike adds a condition value to the ConditionPath
func (c BoolConditionField) ILike(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) ILikeFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) ILikeEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(expression),
	}
}

// In adds a condition value to the ConditionPath
func (c BoolConditionField) In(arg []bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(arg),
	}
}

// InFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) InFunc(arg func(RecordSet) []bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(arg),
	}
}

// InEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) InEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(expression),
	}
}

// NotIn adds a condition value to the ConditionPath
func (c BoolConditionField) NotIn(arg []bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) NotInFunc(arg func(RecordSet) []bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) NotInEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(expression),
	}
}

// ChildOf adds a condition value to the ConditionPath
func (c BoolConditionField) ChildOf(arg bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) ChildOfFunc(arg func(RecordSet) bool) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c BoolConditionField) ChildOfEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(expression),
	}
}

// IsNull checks if the current condition field is null
func (c BoolConditionField) IsNull() ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IsNull(),
	}
}

// IsNotNull checks if the current condition field is not null
func (c BoolConditionField) IsNotNull() ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IsNotNull(),
	}
}

// AddOperator adds a condition value to the condition with the given CondOperator and data
// If multi is true, a recordset will be converted into a slice of int64
// otherwise, it will return an int64 and panic if the recordset is not a singleton.
//
// This method is low level and should be avoided. Use CondOperator methods such as Equals() instead.
func (c BoolConditionField) AddOperator(op operator.Operator, data interface{}) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.AddOperator(op, data),
	}
}
