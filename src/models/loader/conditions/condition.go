package conditions

import "github.com/hexya-erp/hexya/src/models"

// ------- CONDITION ---------

// A Condition is a type safe WHERE clause in an SQL query
type Condition struct {
	*models.Condition
}

// And completes the current condition with a simple And clause : c.And().nextCond => c And nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c Condition) And() ConditionStart {
	return ConditionStart{
		ConditionStart: c.Condition.And(),
	}
}

// AndCond completes the current condition with the given cond as an And clause
// between brackets : c.And(cond) => c And (cond)
func (c Condition) AndCond(cond Condition) Condition {
	return Condition{
		Condition: c.Condition.AndCond(cond.Condition),
	}
}

// AndNot completes the current condition with a simple AndNot clause : c.AndNot().nextCond => c AndNot nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c Condition) AndNot() ConditionStart {
	return ConditionStart{
		ConditionStart: c.Condition.AndNot(),
	}
}

// AndNotCond completes the current condition with the given cond as an AndNot clause
// between brackets : c.AndNot(cond) => c AndNot (cond)
func (c Condition) AndNotCond(cond Condition) Condition {
	return Condition{
		Condition: c.Condition.AndNotCond(cond.Condition),
	}
}

// Or completes the current condition with a simple Or clause : c.Or().nextCond => c Or nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c Condition) Or() ConditionStart {
	return ConditionStart{
		ConditionStart: c.Condition.Or(),
	}
}

// OrCond completes the current condition with the given cond as an Or clause
// between brackets : c.Or(cond) => c Or (cond)
func (c Condition) OrCond(cond Condition) Condition {
	return Condition{
		Condition: c.Condition.OrCond(cond.Condition),
	}
}

// OrNot completes the current condition with a simple OrNot clause : c.OrNot().nextCond => c OrNot nextCond
//
// No brackets are added so AND precedence over OR applies.
func (c Condition) OrNot() ConditionStart {
	return ConditionStart{
		ConditionStart: c.Condition.OrNot(),
	}
}

// OrNotCond completes the current condition with the given cond as an OrNot clause
// between brackets : c.OrNot(cond) => c OrNot (cond)
func (c Condition) OrNotCond(cond Condition) Condition {
	return Condition{
		Condition: c.Condition.OrNotCond(cond.Condition),
	}
}

// Underlying returns the underlying models.Condition instance
func (c Condition) Underlying() *models.Condition {
	return c.Condition
}

// AttachmentConditionHexyaFunc is a dummy function to uniquely match interfaces.
func (c Condition) AttachmentConditionHexyaFunc() {}

// ------- CONDITION START ---------

// A ConditionStart is an object representing a Condition when
// we just added a logical operator (AND, OR, ...) and we are
// about to add a predicate.
type ConditionStart struct {
	*models.ConditionStart
}

// NewCondition returns a valid empty Condition
func (cs ConditionStart) NewCondition() Condition {
	return Condition{
		Condition: &models.Condition{},
	}
}

// CompanyFilteredOn adds a condition with a table join on the given field and
// filters the result with the given condition
func (cs ConditionStart) ModelField(fieldName string, cond Condition) Condition {
	return Condition{
		Condition: cs.FilteredOn(models.NewFieldName(fieldName, ""), cond.Underlying()),
	}
}

// CreateDate adds the "CreateDate" field to the Condition
func (cs ConditionStart) CreateDate() DatesDateTimeConditionField {
	return DatesDateTimeConditionField{
		ConditionField: cs.Field(models.NewFieldName("CreateDate", "create_date")),
	}
}

// CreateUID adds the "CreateUID" field to the Condition
func (cs ConditionStart) CreateUID() NumberConditionField[int64] {
	return NumberConditionField[int64]{
		ConditionField: cs.Field(models.NewFieldName("CreateUID", "create_uid")),
	}
}

// DBDatas adds the "DBDatas" field to the Condition
func (cs ConditionStart) DBDatas() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName("DBDatas", "db_datas")),
	}
}

// DisplayName adds the "DisplayName" field to the Condition
func (cs ConditionStart) DisplayName() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName("DisplayName", "display_name")),
	}
}

// FileSize adds the "FileSize" field to the Condition
func (cs ConditionStart) Int(fieleName string) NumberConditionField[int] {
	return NumberConditionField[int]{
		ConditionField: cs.Field(models.NewFieldName(fieleName, "")),
	}
}

// HexyaExternalID adds the "HexyaExternalID" field to the Condition
func (cs ConditionStart) HexyaExternalID() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName("HexyaExternalID", "hexya_external_id")),
	}
}

// HexyaVersion adds the "HexyaVersion" field to the Condition
func (cs ConditionStart) HexyaVersion() NumberConditionField[int] {
	return NumberConditionField[int]{
		ConditionField: cs.Field(models.NewFieldName("HexyaVersion", "hexya_version")),
	}
}

// ID adds the "ID" field to the Condition
func (cs ConditionStart) ID() NumberConditionField[int64] {
	return NumberConditionField[int64]{
		ConditionField: cs.Field(models.NewFieldName("ID", "id")),
	}
}

// LastUpdate adds the "LastUpdate" field to the Condition
func (cs ConditionStart) LastUpdate() DatesDateTimeConditionField {
	return DatesDateTimeConditionField{
		ConditionField: cs.Field(models.NewFieldName("LastUpdate", "__last_update")),
	}
}

// Name adds the "Name" field to the Condition
func (cs ConditionStart) String(fieldName string) StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName(fieldName, "")),
	}
}

// Public adds the "Public" field to the Condition
func (cs ConditionStart) Bool(fieldName string) BoolConditionField {
	return BoolConditionField{
		ConditionField: cs.Field(models.NewFieldName(fieldName, "")),
	}
}

// ResField adds the "ResField" field to the Condition
func (cs ConditionStart) ResField() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName("ResField", "res_field")),
	}
}

// ResID adds the "ResID" field to the Condition
func (cs ConditionStart) ResID() NumberConditionField[int64] {
	return NumberConditionField[int64]{
		ConditionField: cs.Field(models.NewFieldName("ResID", "res_id")),
	}
}

// ResModel adds the "ResModel" field to the Condition
func (cs ConditionStart) ResModel() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName("ResModel", "res_model")),
	}
}

// ResName adds the "ResName" field to the Condition
func (cs ConditionStart) ResName() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName("ResName", "res_name")),
	}
}

// StoreFname adds the "StoreFname" field to the Condition
func (cs ConditionStart) StoreFname() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName("StoreFname", "store_fname")),
	}
}

// Type adds the "Type" field to the Condition
func (cs ConditionStart) Type() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName("Type", "type")),
	}
}

// URL adds the "URL" field to the Condition
func (cs ConditionStart) URL() StringConditionField {
	return StringConditionField{
		ConditionField: cs.Field(models.NewFieldName("URL", "url")),
	}
}

// WriteDate adds the "WriteDate" field to the Condition
func (cs ConditionStart) WriteDate() DatesDateTimeConditionField {
	return DatesDateTimeConditionField{
		ConditionField: cs.Field(models.NewFieldName("WriteDate", "write_date")),
	}
}

// WriteUID adds the "WriteUID" field to the Condition
func (cs ConditionStart) WriteUID() NumberConditionField[int64] {
	return NumberConditionField[int64]{
		ConditionField: cs.Field(models.NewFieldName("WriteUID", "write_uid")),
	}
}
