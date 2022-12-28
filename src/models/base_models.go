package models

import "time"

const (
	// TransientModel means that the records of this model will be automatically
	// removed periodically. Transient models are mainly used for wizards.
	TransientModel Option = 1 << iota
	// MixinModel means that this model will not be accessible like a regular model
	// but is meant to be mixed in other models.
	MixinModel
	// Many2ManyLinkModel is a model that abstracts the link
	// table of a many2many relationship
	Many2ManyLinkModel
	// ContextsModel is a model for holding fields values that depend on contexts
	ContextsModel
	// ManualModel is a model whose table is not automatically generated in the
	// database. Such models include SQL views and materialized SQL views.
	ManualModel
	// SystemModel is a model that is used internally by the Hexya Framework
	SystemModel
)

type DataModel interface {
	IsModel() bool
	IsTransient() bool
	IsAbstract() bool
}

// HexyaBaseModel For Database persisted models and data
type HexyaBaseModel struct {
	CreateDate  time.Time `json:"create_date" hexya:"type=datetime;display_name=Created On;noCopy"`
	CreateUID   int64     `json:"create_uid" hexya:"display_name=Created By;noCopy"`
	WriteDate   time.Time `json:"write_date" hexya:"type=datetime;display_name=Updated On;noCopy"`
	WriteUID    int64     `json:"write_uid" hexya:"display_name=Updated By;noCopy"`
	LastUpdate  time.Time `json:"__last_update" hexya:"type=datetime;display_name=Updated On;noCopy"`
	DisplayName string    `json:"display_name" hexya:"type=compute;display_name=Display Name;noCopy"`
}

func (HexyaBaseModel) IsModel() bool {
	return true
}
func (HexyaBaseModel) IsTransient() bool {
	return false
}

func (HexyaBaseModel) IsAbstract() bool {
	return false
}

// HexyaTransientModel extends base model
// For models to be cleared after sometime
type HexyaTransientModel struct {
	HexyaBaseModel
}

func (HexyaTransientModel) IsTransient() bool {
	return true
}

// HexyaAbstractModel To be inherited by other models
type HexyaAbstractModel struct {
}

func (HexyaAbstractModel) IsModel() bool {
	return false
}
func (HexyaAbstractModel) IsTransient() bool {
	return false
}

func (HexyaAbstractModel) IsAbstract() bool {
	return true
}
