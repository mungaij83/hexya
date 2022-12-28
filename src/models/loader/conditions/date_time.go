package conditions

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/operator"
	"github.com/hexya-erp/hexya/src/models/types/dates"
)

// A DatesDateTimeConditionField is a partial Condition when
// we have selected a field of type dates.DateTime and expecting an operator.
type DatesDateTimeConditionField struct {
	*models.ConditionField
}

// Equals adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Equals(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) EqualsFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) EqualsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Equals(models.ClientEvaluatedString(expression)),
	}
}

// NotEquals adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) NotEquals(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) NotEqualsFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) NotEqualsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotEquals(models.ClientEvaluatedString(expression)),
	}
}

// Greater adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Greater(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) GreaterFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) GreaterEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Greater(models.ClientEvaluatedString(expression)),
	}
}

// GreaterOrEqual adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) GreaterOrEqual(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) GreaterOrEqualFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) GreaterOrEqualEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.GreaterOrEqual(models.ClientEvaluatedString(expression)),
	}
}

// Lower adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Lower(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) LowerFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) LowerEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Lower(models.ClientEvaluatedString(expression)),
	}
}

// LowerOrEqual adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) LowerOrEqual(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) LowerOrEqualFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) LowerOrEqualEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.LowerOrEqual(models.ClientEvaluatedString(expression)),
	}
}

// Like adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Like(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) LikeFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) LikeEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Like(models.ClientEvaluatedString(expression)),
	}
}

// Contains adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) Contains(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) ContainsFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) ContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.Contains(models.ClientEvaluatedString(expression)),
	}
}

// NotContains adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) NotContains(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) NotContainsFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) NotContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotContains(models.ClientEvaluatedString(expression)),
	}
}

// IContains adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) IContains(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) IContainsFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) IContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.IContains(models.ClientEvaluatedString(expression)),
	}
}

// NotIContains adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) NotIContains(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) NotIContainsFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) NotIContainsEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIContains(models.ClientEvaluatedString(expression)),
	}
}

// ILike adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) ILike(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) ILikeFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) ILikeEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.ILike(models.ClientEvaluatedString(expression)),
	}
}

// In adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) In(arg []dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.In(arg),
	}
}

// InFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) InFunc(arg func(models.RecordSet) []dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.In(arg),
	}
}

// InEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) InEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.In(models.ClientEvaluatedString(expression)),
	}
}

// NotIn adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) NotIn(arg []dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) NotInFunc(arg func(models.RecordSet) []dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) NotInEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.NotIn(models.ClientEvaluatedString(expression)),
	}
}

// ChildOf adds a condition value to the ConditionPath
func (c DatesDateTimeConditionField) ChildOf(arg dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c DatesDateTimeConditionField) ChildOfFunc(arg func(models.RecordSet) dates.DateTime) Condition {
	return Condition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting Condition cannot
// be used server-side.
func (c DatesDateTimeConditionField) ChildOfEval(expression string) Condition {
	return Condition{
		Condition: c.ConditionField.ChildOf(models.ClientEvaluatedString(expression)),
	}
}

// IsNull checks if the current condition field is null
func (c DatesDateTimeConditionField) IsNull() Condition {
	return Condition{
		Condition: c.ConditionField.IsNull(),
	}
}

// IsNotNull checks if the current condition field is not null
func (c DatesDateTimeConditionField) IsNotNull() Condition {
	return Condition{
		Condition: c.ConditionField.IsNotNull(),
	}
}

// AddOperator adds a condition value to the condition with the given operator and data
// If multi is true, a recordset will be converted into a slice of int64
// otherwise, it will return an int64 and panic if the recordset is not a singleton.
//
// This method is low level and should be avoided. Use operator methods such as Equals() instead.
func (c DatesDateTimeConditionField) AddOperator(op operator.Operator, data interface{}) Condition {
	return Condition{
		Condition: c.ConditionField.AddOperator(op, data),
	}
}
