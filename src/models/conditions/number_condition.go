package conditions

import (
	"github.com/hexya-erp/hexya/src/models/operator"
)

type Numbers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | uint32 | ~uint64
}

// A NumberConditionField is a partial ModelCondition when
// we have selected a field of type int and expecting an CondOperator.
type NumberConditionField[T Numbers] struct {
	*ConditionField
}

// Equals adds a condition value to the ConditionPath
func (c NumberConditionField[T]) Equals(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) EqualsFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(arg),
	}
}

// EqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) EqualsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Equals(expression),
	}
}

// NotEquals adds a condition value to the ConditionPath
func (c NumberConditionField[T]) NotEquals(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) NotEqualsFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(arg),
	}
}

// NotEqualsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) NotEqualsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotEquals(expression),
	}
}

// Greater adds a condition value to the ConditionPath
func (c NumberConditionField[T]) Greater(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) GreaterFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(arg),
	}
}

// GreaterEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) GreaterEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Greater(expression),
	}
}

// GreaterOrEqual adds a condition value to the ConditionPath
func (c NumberConditionField[T]) GreaterOrEqual(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) GreaterOrEqualFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(arg),
	}
}

// GreaterOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) GreaterOrEqualEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.GreaterOrEqual(expression),
	}
}

// Lower adds a condition value to the ConditionPath
func (c NumberConditionField[T]) Lower(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) LowerFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(arg),
	}
}

// LowerEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) LowerEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Lower(expression),
	}
}

// LowerOrEqual adds a condition value to the ConditionPath
func (c NumberConditionField[T]) LowerOrEqual(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) LowerOrEqualFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(arg),
	}
}

// LowerOrEqualEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) LowerOrEqualEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.LowerOrEqual(expression),
	}
}

// Like adds a condition value to the ConditionPath
func (c NumberConditionField[T]) Like(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) LikeFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(arg),
	}
}

// LikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) LikeEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Like(expression),
	}
}

// Contains adds a condition value to the ConditionPath
func (c NumberConditionField[T]) Contains(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) ContainsFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(arg),
	}
}

// ContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) ContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.Contains(expression),
	}
}

// NotContains adds a condition value to the ConditionPath
func (c NumberConditionField[T]) NotContains(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) NotContainsFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(arg),
	}
}

// NotContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) NotContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotContains(expression),
	}
}

// IContains adds a condition value to the ConditionPath
func (c NumberConditionField[T]) IContains(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) IContainsFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(arg),
	}
}

// IContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) IContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IContains(expression),
	}
}

// NotIContains adds a condition value to the ConditionPath
func (c NumberConditionField[T]) NotIContains(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) NotIContainsFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(arg),
	}
}

// NotIContainsEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) NotIContainsEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIContains(expression),
	}
}

// ILike adds a condition value to the ConditionPath
func (c NumberConditionField[T]) ILike(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) ILikeFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(arg),
	}
}

// ILikeEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) ILikeEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ILike(expression),
	}
}

// In adds a condition value to the ConditionPath
func (c NumberConditionField[T]) In(arg []T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(arg),
	}
}

// InFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) InFunc(arg func(RecordSet) []T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(arg),
	}
}

// InEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) InEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.In(expression),
	}
}

// NotIn adds a condition value to the ConditionPath
func (c NumberConditionField[T]) NotIn(arg []T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) NotInFunc(arg func(RecordSet) []T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(arg),
	}
}

// NotInEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) NotInEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.NotIn(expression),
	}
}

// ChildOf adds a condition value to the ConditionPath
func (c NumberConditionField[T]) ChildOf(arg T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfFunc adds a function value to the ConditionPath.
// The function will be evaluated when the query is performed and
// it will be given the RecordSet on which the query is made as parameter
func (c NumberConditionField[T]) ChildOfFunc(arg func(RecordSet) T) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(arg),
	}
}

// ChildOfEval adds an expression value to the ConditionPath.
// The expression value will be evaluated by the client with the
// corresponding execution context. The resulting ModelCondition cannot
// be used server-side.
func (c NumberConditionField[T]) ChildOfEval(expression string) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.ChildOf(expression),
	}
}

// IsNull checks if the current condition field is null
func (c NumberConditionField[T]) IsNull() ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IsNull(),
	}
}

// IsNotNull checks if the current condition field is not null
func (c NumberConditionField[T]) IsNotNull() ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.IsNotNull(),
	}
}

// AddOperator adds a condition value to the condition with the given CondOperator and data
// If multi is true, a recordset will be converted into a slice of int64
// otherwise, it will return an int64 and panic if the recordset is not a singleton.
//
// This method is low level and should be avoided. Use CondOperator methods such as Equals() instead.
func (c NumberConditionField[T]) AddOperator(op operator.Operator, data interface{}) ModelCondition {
	return ModelCondition{
		Condition: c.ConditionField.AddOperator(op, data),
	}
}
