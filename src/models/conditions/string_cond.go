package conditions

import (
	"github.com/hexya-erp/hexya/src/models/operator"
)

// A StringConditionField is a partial ModelCondition when
// we have selected a field of type string and expecting an CondOperator.
type StringConditionField struct {
	*ConditionField
}

// Equals adds a condition value to the ConditionPath
func (c StringConditionField) Equals(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) EqualsFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) EqualsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(expression),
	}
}

// NotEquals adds a condition value to the ConditionPath
func (c StringConditionField) NotEquals(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) NotEqualsFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) NotEqualsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(expression),
	}
}

// Greater adds a condition value to the ConditionPath
func (c StringConditionField) Greater(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) GreaterFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) GreaterEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(expression),
	}
}

// GreaterOrEqual adds a condition value to the ConditionPath
func (c StringConditionField) GreaterOrEqual(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) GreaterOrEqualFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) GreaterOrEqualEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(expression),
	}
}

// Lower adds a condition value to the ConditionPath
func (c StringConditionField) Lower(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) LowerFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) LowerEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(expression),
	}
}

// LowerOrEqual adds a condition value to the ConditionPath
func (c StringConditionField) LowerOrEqual(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) LowerOrEqualFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) LowerOrEqualEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(expression),
	}
}

// Like adds a condition value to the ConditionPath
func (c StringConditionField) Like(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) LikeFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) LikeEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(expression),
	}
}

// Contains adds a condition value to the ConditionPath
func (c StringConditionField) Contains(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) ContainsFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) ContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(expression),
	}
}

// NotContains adds a condition value to the ConditionPath
func (c StringConditionField) NotContains(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) NotContainsFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) NotContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(expression),
	}
}

// IContains adds a condition value to the ConditionPath
func (c StringConditionField) IContains(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) IContainsFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) IContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(expression),
	}
}

// NotIContains adds a condition value to the ConditionPath
func (c StringConditionField) NotIContains(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) NotIContainsFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) NotIContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(expression),
	}
}

// ILike adds a condition value to the ConditionPath
func (c StringConditionField) ILike(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) ILikeFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) ILikeEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(expression),
	}
}

// In adds a condition value to the ConditionPath
func (c StringConditionField) In(arg []string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(arg),
	}
}

// InFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) InFunc(arg func(RecordSet) []string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(arg),
	}
}

// InEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) InEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(expression),
	}
}

// NotIn adds a condition value to the ConditionPath
func (c StringConditionField) NotIn(arg []string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) NotInFunc(arg func(RecordSet) []string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) NotInEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(expression),
	}
}

// ChildOf adds a condition value to the ConditionPath
func (c StringConditionField) ChildOf(arg string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c StringConditionField) ChildOfFunc(arg func(RecordSet) string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c StringConditionField) ChildOfEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(expression),
	}
}

// IsNull checks if the current condition field is null
func (c StringConditionField) IsNull() ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IsNull(),
	}
}

// IsNotNull checks if the current condition field is not null
func (c StringConditionField) IsNotNull() ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IsNotNull(),
	}
}

// AddOperator adds a condition value to the condition with the given CondOperator and data
// If multi is true, a recordset will be converted into a slice of int64
// otherwise, it will return an int64 and panic if the recordset is not a singleton.
//
// This method is low level and should be avoided. Use CondOperator methods such as Equals() instead.
func (c StringConditionField) AddOperator(op operator.Operator, data interface{}) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.AddOperator(op, data),
	}
}
