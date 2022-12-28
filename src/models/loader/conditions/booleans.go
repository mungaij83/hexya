package conditions

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/operator"
)

// A BoolConditionField is a partial Condition when
// we have selected a field of type bool and expecting an operator.
type BoolConditionField struct {
	*models.ConditionField
}

// Equals adds a condition value to the ConditionPath
func (c BoolConditionField) Equals(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) EqualsFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) EqualsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Equals(models.ClientEvaluatedString(expression)),
	}
}

// NotEquals adds a condition value to the ConditionPath
func (c BoolConditionField) NotEquals(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) NotEqualsFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) NotEqualsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotEquals(models.ClientEvaluatedString(expression)),
	}
}

// Greater adds a condition value to the ConditionPath
func (c BoolConditionField) Greater(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) GreaterFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) GreaterEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Greater(models.ClientEvaluatedString(expression)),
	}
}

// GreaterOrEqual adds a condition value to the ConditionPath
func (c BoolConditionField) GreaterOrEqual(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) GreaterOrEqualFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) GreaterOrEqualEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.GreaterOrEqual(models.ClientEvaluatedString(expression)),
	}
}

// Lower adds a condition value to the ConditionPath
func (c BoolConditionField) Lower(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) LowerFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) LowerEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Lower(models.ClientEvaluatedString(expression)),
	}
}

// LowerOrEqual adds a condition value to the ConditionPath
func (c BoolConditionField) LowerOrEqual(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) LowerOrEqualFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) LowerOrEqualEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.LowerOrEqual(models.ClientEvaluatedString(expression)),
	}
}

// Like adds a condition value to the ConditionPath
func (c BoolConditionField) Like(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) LikeFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) LikeEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Like(models.ClientEvaluatedString(expression)),
	}
}

// Contains adds a condition value to the ConditionPath
func (c BoolConditionField) Contains(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) ContainsFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) ContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Contains(models.ClientEvaluatedString(expression)),
	}
}

// NotContains adds a condition value to the ConditionPath
func (c BoolConditionField) NotContains(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) NotContainsFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) NotContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotContains(models.ClientEvaluatedString(expression)),
	}
}

// IContains adds a condition value to the ConditionPath
func (c BoolConditionField) IContains(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) IContainsFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) IContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.IContains(models.ClientEvaluatedString(expression)),
	}
}

// NotIContains adds a condition value to the ConditionPath
func (c BoolConditionField) NotIContains(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) NotIContainsFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) NotIContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIContains(models.ClientEvaluatedString(expression)),
	}
}

// ILike adds a condition value to the ConditionPath
func (c BoolConditionField) ILike(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) ILikeFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) ILikeEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.ILike(models.ClientEvaluatedString(expression)),
	}
}

// In adds a condition value to the ConditionPath
func (c BoolConditionField) In(arg []bool) Condition {
	return Condition{
		Condition: c.ConditionField.In(arg),
	}
}

// InFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) InFunc(arg func(models.RecordSet) []bool) Condition {
	return Condition{
		Condition: c.ConditionField.In(arg),
	}
}

// InEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) InEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.In(models.ClientEvaluatedString(expression)),
	}
}

// NotIn adds a condition value to the ConditionPath
func (c BoolConditionField) NotIn(arg []bool) Condition {
	return Condition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) NotInFunc(arg func(models.RecordSet) []bool) Condition {
	return Condition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) NotInEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIn(models.ClientEvaluatedString(expression)),
	}
}

// ChildOf adds a condition value to the ConditionPath
func (c BoolConditionField) ChildOf(arg bool) Condition {
	return Condition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c BoolConditionField) ChildOfFunc(arg func(models.RecordSet) bool) Condition {
	return Condition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c BoolConditionField) ChildOfEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.ChildOf(models.ClientEvaluatedString(expression)),
	}
}

// IsNull checks if the current condition field is null
func (c BoolConditionField) IsNull() Condition {
	return Condition{
		Condition: c.ConditionField.IsNull(),
	}
}

// IsNotNull checks if the current condition field is not null
func (c BoolConditionField) IsNotNull() Condition {
	return Condition{
		Condition: c.ConditionField.IsNotNull(),
	}
}

// AddOperator adds a condition value to the condition with the given operator and data
// If multi is true, a recordset will be converted into a slice of int64
// otherwise, it will return an int64 and panic if the recordset is not a singleton.
//
// This method is low level and should be avoided. Use operator methods such as Equals() instead.
func (c BoolConditionField) AddOperator(op operator.Operator, data interface{}) Condition {
	return Condition{
		Condition: c.ConditionField.AddOperator(op, data),
	}
}
