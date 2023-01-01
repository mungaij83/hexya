package models

import (
	"errors"
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/tools"
)

// RecordOptions extra options to be used with GORM models
type RecordOptions struct {
	EagerLoad []string
}

// PrimaryKeys represent the list of field types that can act as primary keys to a table
type PrimaryKeys interface {
	~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | string
}

// Repository Type of repository to represent any query link between
type Repository[T any, K PrimaryKeys] interface {
	Save(v T) (*T, error)
	Search(cond loader.Condition) ([]T, error)
	TableName() string
	Delete(v T) (*T, error)
	validateAndInitialize(modelLoader *ModelLoader) error
	setEnv(v *loader.Environment) error
	FindById(id K) (*T, error)
	FindByIdWithOptions(id K, options RecordOptions) (*T, error)
	FindByIds(id []K) ([]T, error)
	FindByIdsWithOptions(id []K, option RecordOptions) ([]T, error)
	Methods() *loader.MethodsCollection
	Fields() *loader.FieldsCollection
	GetFieldName(s string) loader.FieldName
	GetModel() (*loader.Model, bool)
	IsMixin() bool
	IsManual() bool
	isSystem() bool
	isContext() bool
	IsM2MLink() bool
	IsTransient() bool
	hasParentField() bool
}

// ModelRepository default implementation of model repository
type ModelRepository[T DataModel, K PrimaryKeys] struct {
	env       *loader.Environment
	tableName string
	model     *loader.Model
}

func (mr ModelRepository[T, K]) validateAndInitialize(loader *ModelLoader) error {
	if mr.env == nil {
		return errors.New("environment is not set for this repository, call set env first")
	}
	mdl, err := loader.LoadBaseModel(new(T))
	if err != nil {
		return err
	}
	// Migrate this model
	err = mr.env.Cr().AutoMigrate(new(T))
	if err != nil {
		return err
	}
	mr.model = mdl
	// Resolve model table name
	mr.tableName = mr.env.Cr().Unscoped().Model(new(T)).Name()
	return nil
}
func (mr ModelRepository[T, K]) GetModel() (*loader.Model, bool) {
	return mr.model, mr.model != nil
}
func (mr ModelRepository[T, K]) setEnv(env *loader.Environment) error {
	if mr.env != nil {
		return errors.New("tried to reinitialize environment")
	}
	mr.env = env
	return nil
}

func (mr ModelRepository[T, K]) Save(v T) (*T, error) {
	err := mr.env.Cr().Save(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}
func (mr ModelRepository[T, K]) Search(cond loader.Condition) ([]T, error) {
	var vv []T
	err := mr.env.Cr().Find(&vv).Error
	if err != nil {
		return nil, err
	}
	return vv, nil
}
func (mr ModelRepository[T, K]) TableName() string {
	return mr.tableName
}

func (mr ModelRepository[T, K]) Methods() *loader.MethodsCollection {
	return mr.model.Methods()
}

func (mr ModelRepository[T, K]) Fields() *loader.FieldsCollection {
	return mr.Fields()
}

func (mr ModelRepository[T, K]) GetFieldName(s string) loader.FieldName {
	dd, ok := mr.model.Fields().Get(s)
	if !ok {
		return nil
	}
	return loader.NewFieldName(dd.Name(), dd.JSON())
}

func (mr ModelRepository[T, K]) Delete(v T) (*T, error) {
	err := mr.env.Cr().Delete(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (mr ModelRepository[T, K]) FindById(id K) (*T, error) {
	var v T
	err := mr.env.Cr().First(&v, id).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (mr ModelRepository[T, K]) FindByIdWithOptions(id K, options RecordOptions) (*T, error) {
	var v T
	db := mr.env.Cr()
	if len(options.EagerLoad) > 0 {
		for _, opt := range options.EagerLoad {
			db = db.Preload(opt)
		}
	}
	err := db.First(&v, id).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (mr ModelRepository[T, K]) FindByIds(ids []K) ([]T, error) {
	var v []T
	err := mr.env.Cr().Find(&v, ids).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (mr ModelRepository[T, K]) FindByIdsWithOptions(id []K, options RecordOptions) ([]T, error) {
	var v []T
	db := mr.env.Cr()
	if len(options.EagerLoad) > 0 {
		for _, opt := range options.EagerLoad {
			db.Preload(opt)
		}
	}
	err := db.Find(&v, id).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (mr ModelRepository[T, K]) IsMixin() bool {
	if mr.model.Options()&tools.MixinModel > 0 {
		return true
	}
	return false
}

// IsManual returns true if this is a manual model.
func (mr ModelRepository[T, K]) IsManual() bool {
	if mr.model.Options()&tools.ManualModel > 0 {
		return true
	}
	return false
}

// isSystem returns true if this is a system model.
func (mr ModelRepository[T, K]) isSystem() bool {
	if mr.model.Options()&tools.SystemModel > 0 {
		return true
	}
	return false
}

// isContext returns true if this is a context model.
func (mr ModelRepository[T, K]) isContext() bool {
	if mr.model.Options()&tools.ContextsModel > 0 {
		return true
	}
	return false
}

// IsM2MLink returns true if this is an M2M Link model.
func (mr ModelRepository[T, K]) IsM2MLink() bool {
	if mr.model.Options()&tools.Many2ManyLinkModel > 0 {
		return true
	}
	return false
}

// IsTransient returns true if this Model is transient
func (mr ModelRepository[T, K]) IsTransient() bool {
	return mr.model.Options() == tools.TransientModel
}

// hasParentField returns true if this model is recursive and has a Parent field.
func (mr ModelRepository[T, K]) hasParentField() bool {
	_, parentExists := mr.Fields().Get("Parent")
	return parentExists
}
