package models

import (
	"github.com/hexya-erp/hexya/src/models/fieldtype"
	"github.com/hexya-erp/hexya/src/models/types"
	"reflect"
)

// FieldInfo is the exportable field information struct
type FieldInfo struct {
	ChangeDefault    bool                                  `json:"change_default"`
	Help             string                                `json:"help"`
	Searchable       bool                                  `json:"searchable"`
	Views            map[string]interface{}                `json:"views"`
	Required         bool                                  `json:"required"`
	Manual           bool                                  `json:"manual"`
	ReadOnly         bool                                  `json:"readonly"`
	Depends          []string                              `json:"depends"`
	CompanyDependent bool                                  `json:"company_dependent"`
	Sortable         bool                                  `json:"sortable"`
	Translate        bool                                  `json:"translate"`
	Type             fieldtype.Type                        `json:"type"`
	Store            bool                                  `json:"store"`
	String           string                                `json:"string"`
	Relation         string                                `json:"relation"`
	Selection        types.Selection                       `json:"selection"`
	Domain           interface{}                           `json:"domain"`
	OnChange         bool                                  `json:"-"`
	ReverseFK        string                                `json:"-"`
	Name             string                                `json:"-"`
	JSON             string                                `json:"-"`
	ReadOnlyFunc     func(Environment) (bool, Conditioner) `json:"-"`
	RequiredFunc     func(Environment) (bool, Conditioner) `json:"-"`
	InvisibleFunc    func(Environment) (bool, Conditioner) `json:"-"`
	DefaultFunc      func(Environment) interface{}         `json:"-"`
	GoType           reflect.Type                          `json:"-"`
	Index            bool                                  `json:"-"`
}

// FieldsGetArgs is the args struct for the FieldsGet method
type FieldsGetArgs struct {
	// Fields is a list of fields to document, all if empty or not provided
	Fields FieldNames `json:"allfields"`
}

// OnchangeParams is the args struct of the Onchange function
type OnchangeParams struct {
	Values   RecordData        `json:"values"`
	Fields   FieldNames        `json:"field_name"`
	Onchange map[string]string `json:"field_onchange"`
}

// OnchangeResult is the result struct type of the Onchange function
type OnchangeResult struct {
	Value   RecordData                `json:"value"`
	Warning string                    `json:"warning"`
	Filters map[FieldName]Conditioner `json:"domain"`
}
