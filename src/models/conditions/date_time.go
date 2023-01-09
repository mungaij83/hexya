package conditions

import (
	"github.com/hexya-erp/hexya/src/models/operator"
	"github.com/hexya-erp/hexya/src/models/types/dates"
)

// A DatesDateTimeConditionField is a partial ModelCondition when
// we have selected a field of type dates.DateTime and expecting an CondOperator.
type DatesDateTimeConditionField struct {
	*ConditionField
}

// Equals adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Equals(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) EqualsFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) EqualsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(expression),
	}
}

// NotEquals adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) NotEquals(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) NotEqualsFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) NotEqualsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(expression),
	}
}

// Greater adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Greater(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) GreaterFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) GreaterEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(expression),
	}
}

// GreaterOrEqual adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) GreaterOrEqual(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) GreaterOrEqualFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) GreaterOrEqualEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(expression),
	}
}

// Lower adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Lower(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) LowerFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) LowerEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(expression),
	}
}

// LowerOrEqual adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) LowerOrEqual(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) LowerOrEqualFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) LowerOrEqualEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(expression),
	}
}

// Like adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Like(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) LikeFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) LikeEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(expression),
	}
}

// Contains adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Contains(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) ContainsFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) ContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(expression),
	}
}

// NotContains adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) NotContains(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) NotContainsFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) NotContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(expression),
	}
}

// IContains adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) IContains(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) IContainsFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) IContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(expression),
	}
}

// NotIContains adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) NotIContains(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) NotIContainsFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) NotIContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(expression),
	}
}

// ILike adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) ILike(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) ILikeFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) ILikeEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(expression),
	}
}

// In adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) In(arg []dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(arg),
	}
}

// InFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) InFunc(arg func(RecordSet) []dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(arg),
	}
}

// InEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) InEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(expression),
	}
}

// NotIn adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) NotIn(arg []dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) NotInFunc(arg func(RecordSet) []dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) NotInEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(expression),
	}
}

// ChildOf adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) ChildOf(arg dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) ChildOfFunc(arg func(RecordSet) dates.DateTime) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c DatesDateTimeConditionField) ChildOfEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(expression),
	}
}

// IsNull checks if the current condition field is null
func (c DatesDateTimeConditionField) IsNull() ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IsNull(),
	}
}

// IsNotNull checks if the current condition field is not null
func (c DatesDateTimeConditionField) IsNotNull() ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IsNotNull(),
	}
}

// AddOperator adds a condition value to the condition with the given CondOperator and data
// If multi is true, a recordset will be converted into a slice of int64
// otherwise, it will return an int64 and panic if the recordset is not a singleton.
//
// This method is low level and should be avoided. Use CondOperator methods such as Equals() instead.
func (c DatesDateTimeConditionField) AddOperator(op operator.Operator, data interface{}) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.AddOperator(op, data),
	}
}
