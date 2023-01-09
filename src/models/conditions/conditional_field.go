package conditions

// A ConditionField is a partial Condition when we have set
// a field name in a ConditionPredicate and are about to add an CondOperator.
type ConditionField struct {
	cs    ConditionStart
	Exprs []FieldName
}

// JSON returns the json field name of this ConditionField
func (c ConditionField) JSON() string {
	return JoinFieldNames(c.Exprs, ExprSep).JSON()
}

// Name method for ConditionField
func (c ConditionField) Name() string {
	return JoinFieldNames(c.Exprs, ExprSep).Name()
}
