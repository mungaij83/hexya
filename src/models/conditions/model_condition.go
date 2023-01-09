package conditions

// Expression separation symbols
const (
	ExprSep    = "."
	SqlSep     = "__"
	ContextSep = "|"
)

// ------- CONDITION ---------

// A ModelCondition is a type safe WHERE clause in an SQL query
type ModelCondition struct {
	*Condition
}

// newCondition returns a new condition struct
func NewModelCondition() *ModelCondition {
	c := &Condition{}
	return &ModelCondition{c}
}

// And completes the current condition with a simple And clause : c.And().nextCond => c And nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c ModelCondition) And() ModelConditionStart {
	return ModelConditionStart{
		ConditionStart: c.Condition.And(),
	}
}

// AndCond completes the current condition with the given cond as an And clause
// between brackets : c.And(cond) => c And (cond)
func (c ModelCondition) AndCond(cond ModelCondition) ModelCondition {
	return ModelCondition{
		Condition: c.Condition.AndCond(cond.Condition),
	}
}

// AndNot completes the current condition with a simple AndNot clause : c.AndNot().nextCond => c AndNot nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c ModelCondition) AndNot() ModelConditionStart {
	return ModelConditionStart{
		ConditionStart: c.Condition.AndNot(),
	}
}

// AndNotCond completes the current condition with the given cond as an AndNot clause
// between brackets : c.AndNot(cond) => c AndNot (cond)
func (c ModelCondition) AndNotCond(cond ModelCondition) ModelCondition {
	return ModelCondition{
		Condition: c.Condition.AndNotCond(cond.Condition),
	}
}

// Or completes the current condition with a simple Or clause : c.Or().nextCond => c Or nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c ModelCondition) Or() ModelConditionStart {
	return ModelConditionStart{
		ConditionStart: c.Condition.Or(),
	}
}

// OrCond completes the current condition with the given cond as an Or clause
// between brackets : c.Or(cond) => c Or (cond)
func (c ModelCondition) OrCond(cond ModelCondition) ModelCondition {
	return ModelCondition{
		Condition: c.Condition.OrCond(cond.Condition),
	}
}

// OrNot completes the current condition with a simple OrNot clause : c.OrNot().nextCond => c OrNot nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c ModelCondition) OrNot() ModelConditionStart {
	return ModelConditionStart{
		ConditionStart: c.Condition.OrNot(),
	}
}

// OrNotCond completes the current condition with the given cond as an OrNot clause
// between brackets : c.OrNot(cond) => c OrNot (cond)
func (c ModelCondition) OrNotCond(cond ModelCondition) ModelCondition {
	return ModelCondition{
		Condition: c.Condition.OrNotCond(cond.Condition),
	}
}

// Underlying returns the underlying models.Condition instance
func (c ModelCondition) Underlying() *Condition {
	return c.Condition
}

// AttachmentConditionHexyaFunc is a dummy function to uniquely match interfaces.
func (c ModelCondition) AttachmentConditionHexyaFunc() {}

// ------- CONDITION START ---------

// A ConditionStart is an object representing a ModelCondition when
// we just added a logical CondOperator (AND, OR, ...) and we are
// about to add a ConditionPredicate.
type ModelConditionStart struct {
	*ConditionStart
}

// NewCondition returns a valid empty ModelCondition
func (cs ConditionStart) NewCondition() ModelCondition {
	return ModelCondition{
		Condition: &Condition{},
	}
}

// CompanyFilteredOn adds a condition with a table join on the given field and
// filters the result with the given condition
func (cs ConditionStart) ModelField(fieldName string, cond ModelCondition) ModelCondition {
	return ModelCondition{
		Condition: cs.FilteredOn(NewFieldName(fieldName, ""), cond.Underlying()),
	}
}

// CreateDate adds the "CreateDate" field to the ModelCondition
func (cs ConditionStart) CreateDate() DatesDateTimeConditionField {
	return DatesDateTimeConditionField{
		ConditionField: cs.Field(NewFieldName("CreateDate", "create_date")),
	}
}

// CreateUID adds the "CreateUID" field to the ModelCondition
func (cs ConditionStart) CreateUID() NumberConditionField[int64] {
	return NumberConditionField[int64]{
		ConditionField: cs.Field(NewFieldName("CreateUID", "create_uid")),
	}
}

// DBDatas adds the "DBDatas" field to the ModelCondition
func (cs ConditionStart) DBDatas() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName("DBDatas", "db_datas")),
	}
}

// DisplayName adds the "DisplayName" field to the ModelCondition
func (cs ConditionStart) DisplayName() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName("DisplayName", "display_name")),
	}
}

// FileSize adds the "FileSize" field to the ModelCondition
func (cs ConditionStart) Int(fieleName string) NumberConditionField[int] {
	return NumberConditionField[int]{
		ConditionField: cs.Field(NewFieldName(fieleName, "")),
	}
}

// HexyaExternalID adds the "HexyaExternalID" field to the ModelCondition
func (cs ConditionStart) HexyaExternalID() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName("HexyaExternalID", "hexya_external_id")),
	}
}

// HexyaVersion adds the "HexyaVersion" field to the ModelCondition
func (cs ConditionStart) HexyaVersion() NumberConditionField[int] {
	return NumberConditionField[int]{
		ConditionField: cs.Field(NewFieldName("HexyaVersion", "hexya_version")),
	}
}

// ID adds the "ID" field to the ModelCondition
func (cs ConditionStart) ID() NumberConditionField[int64] {
	return NumberConditionField[int64]{
		ConditionField: cs.Field(NewFieldName("ID", "id")),
	}
}

// LastUpdate adds the "LastUpdate" field to the ModelCondition
func (cs ConditionStart) LastUpdate() DatesDateTimeConditionField {
	return DatesDateTimeConditionField{
		ConditionField: cs.Field(NewFieldName("LastUpdate", "__last_update")),
	}
}

// Name adds the "Name" field to the ModelCondition
func (cs ConditionStart) String(fieldName string) StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName(fieldName, "")),
	}
}

// Public adds the "Public" field to the ModelCondition
func (cs ConditionStart) Bool(fieldName string) BoolConditionField {
	return BoolConditionField{
		ConditionField: cs.Field(NewFieldName(fieldName, "")),
	}
}

// ResField adds the "ResField" field to the ModelCondition
func (cs ConditionStart) ResField() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName("ResField", "res_field")),
	}
}

// ResID adds the "ResID" field to the ModelCondition
func (cs ConditionStart) ResID() NumberConditionField[int64] {
	return NumberConditionField[int64]{
		ConditionField: cs.Field(NewFieldName("ResID", "res_id")),
	}
}

// ResModel adds the "ResModel" field to the ModelCondition
func (cs ConditionStart) ResModel() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName("ResModel", "res_model")),
	}
}

// ResName adds the "ResName" field to the ModelCondition
func (cs ConditionStart) ResName() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName("ResName", "res_name")),
	}
}

// StoreFname adds the "StoreFname" field to the ModelCondition
func (cs ConditionStart) StoreFname() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName("StoreFname", "store_fname")),
	}
}

// Type adds the "Type" field to the ModelCondition
func (cs ConditionStart) Type() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName("Type", "type")),
	}
}

// URL adds the "URL" field to the ModelCondition
func (cs ConditionStart) URL() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(NewFieldName("URL", "url")),
	}
}

// WriteDate adds the "WriteDate" field to the ModelCondition
func (cs ConditionStart) WriteDate() DatesDateTimeConditionField {
	return DatesDateTimeConditionField{
		ConditionField: cs.Field(NewFieldName("WriteDate", "write_date")),
	}
}

// WriteUID adds the "WriteUID" field to the ModelCondition
func (cs ConditionStart) WriteUID() NumberConditionField[int64] {
	return NumberConditionField[int64]{
		ConditionField: cs.Field(NewFieldName("WriteUID", "write_uid")),
	}
}
