package conditions

import (
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/models/operator"
)

// A StringConditionField is a partial Condition when
// we have selected a field of type string and expecting an operator.
type StringConditionField struct {
	*loader.ConditionField
}

// Equals adds a condition value to the ConditionPath
func (c StringConditionField) Equals(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) EqualsFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) EqualsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Equals(loader.ClientEvaluatedString(expression)),
	}
}

// NotEquals adds a condition value to the ConditionPath
func (c StringConditionField) NotEquals(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) NotEqualsFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) NotEqualsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotEquals(loader.ClientEvaluatedString(expression)),
	}
}

// Greater adds a condition value to the ConditionPath
func (c StringConditionField) Greater(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) GreaterFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) GreaterEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Greater(loader.ClientEvaluatedString(expression)),
	}
}

// GreaterOrEqual adds a condition value to the ConditionPath
func (c StringConditionField) GreaterOrEqual(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) GreaterOrEqualFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) GreaterOrEqualEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.GreaterOrEqual(loader.ClientEvaluatedString(expression)),
	}
}

// Lower adds a condition value to the ConditionPath
func (c StringConditionField) Lower(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) LowerFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) LowerEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Lower(loader.ClientEvaluatedString(expression)),
	}
}

// LowerOrEqual adds a condition value to the ConditionPath
func (c StringConditionField) LowerOrEqual(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) LowerOrEqualFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) LowerOrEqualEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.LowerOrEqual(loader.ClientEvaluatedString(expression)),
	}
}

// Like adds a condition value to the ConditionPath
func (c StringConditionField) Like(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) LikeFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) LikeEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Like(loader.ClientEvaluatedString(expression)),
	}
}

// Contains adds a condition value to the ConditionPath
func (c StringConditionField) Contains(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) ContainsFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) ContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Contains(loader.ClientEvaluatedString(expression)),
	}
}

// NotContains adds a condition value to the ConditionPath
func (c StringConditionField) NotContains(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) NotContainsFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) NotContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotContains(loader.ClientEvaluatedString(expression)),
	}
}

// IContains adds a condition value to the ConditionPath
func (c StringConditionField) IContains(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) IContainsFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) IContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.IContains(loader.ClientEvaluatedString(expression)),
	}
}

// NotIContains adds a condition value to the ConditionPath
func (c StringConditionField) NotIContains(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) NotIContainsFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) NotIContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIContains(loader.ClientEvaluatedString(expression)),
	}
}

// ILike adds a condition value to the ConditionPath
func (c StringConditionField) ILike(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) ILikeFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) ILikeEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.ILike(loader.ClientEvaluatedString(expression)),
	}
}

// In adds a condition value to the ConditionPath
func (c StringConditionField) In(arg []string) Condition {
	return Condition{
		Condition: c.ConditionField.In(arg),
	}
}

// InFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) InFunc(arg func(loader.RecordSet) []string) Condition {
	return Condition{
		Condition: c.ConditionField.In(arg),
	}
}

// InEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) InEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.In(loader.ClientEvaluatedString(expression)),
	}
}

// NotIn adds a condition value to the ConditionPath
func (c StringConditionField) NotIn(arg []string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) NotInFunc(arg func(loader.RecordSet) []string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) NotInEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIn(loader.ClientEvaluatedString(expression)),
	}
}

// ChildOf adds a condition value to the ConditionPath
func (c StringConditionField) ChildOf(arg string) Condition {
	return Condition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) ChildOfFunc(arg func(loader.RecordSet) string) Condition {
	return Condition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c StringConditionField) ChildOfEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.ChildOf(loader.ClientEvaluatedString(expression)),
	}
}

// IsNull checks if the current condition field is null
func (c StringConditionField) IsNull() Condition {
	return Condition{
		Condition: c.ConditionField.IsNull(),
	}
}

// IsNotNull checks if the current condition field is not null
func (c StringConditionField) IsNotNull() Condition {
	return Condition{
		Condition: c.ConditionField.IsNotNull(),
	}
}

// AddOperator adds a condition value to the condition with the given operator and data
// If multi is true, a recordset will be converted into a slice of int64
// otherwise, it will return an int64 and panic if the recordset is not a singleton.
//
// This method is low level and should be avoided. Use operator methods such as Equals() instead.
func (c StringConditionField) AddOperator(op operator.Operator, data interface{}) Condition {
	return Condition{
		Condition: c.ConditionField.AddOperator(op, data),
	}
}
