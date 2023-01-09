package models

import (
	"errors"
	"fmt"
	"github.com/hexya-erp/hexya/src/models/conditions"
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/tools"
	"gorm.io/gorm"
	"reflect"
)

// RecordOptions extra options to be used with GORM models
type RecordOptions struct {
	EagerLoad []string
}

//	type RecordTypes interface {
//		~HexyaBaseModel| ~HexyaAbstractModel| ~HexyaTransientModel
//	}
//
// PrimaryKeys represent the list of field types that can act as primary keys to a table
type PrimaryKeys interface {
	~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | string
}

// Repository Type of repository to represent any query link between
type Repository[T any, K PrimaryKeys] interface {
	Save(v interface{}) error
	Search(cond *conditions.Condition) (interface{}, error)
	TableName() string
	ModelName() string
	Delete(v interface{}) (interface{}, error)
	validateAndInitialize(modelLoader *ModelLoader) error
	setEnv(v *loader.Environment) error
	FindById(id K) (interface{}, error)
	FindByIdWithOptions(id K, options RecordOptions) (interface{}, error)
	FindByIds(id []K) (interface{}, error)
	FindByIdsWithOptions(id []K, option RecordOptions) (interface{}, error)
	Methods() *loader.MethodsCollection
	Fields() *loader.FieldsCollection
	GetFieldName(s string) conditions.FieldName
	GetModel() (*loader.Model, bool)
	Condition() *conditions.ModelConditionStart
	Query(environment *loader.Environment) *loader.RecordCollection
	IsMixin() bool
	IsManual() bool
	isSystem() bool
	isContext() bool
	IsM2MLink() bool
	IsTransient() bool
	hasParentField() bool
	registerExtension(d any) error
}

// BaseRepository abstract type for repository functions
type BaseRepository[T any] struct {
	extensions       map[string]interface{}
	extensionMethods map[string]map[string]*loader.Method
}

func (mr BaseRepository[T]) initialize() {
	mr.extensions = make(map[string]interface{})
	mr.extensionMethods = make(map[string]map[string]*loader.Method)
}

func (mr BaseRepository[T]) RegisterExtension(ext interface{}) error {
	if mr.extensions == nil {
		mr.extensions = make(map[string]interface{})
	}
	extensionName := ""
	e, ok := ext.(ModelExtension[T])
	if !ok {
		return errors.New("invalid extension signature, require method: ExtensionName")
	}
	extensionName = e.ExtensionName()
	_, ok = mr.extensions[extensionName]
	if ok {
		return errors.New(fmt.Sprintf("extension with name already registered: %v", extensionName))
	}
	if mr.extensionMethods == nil {
		mr.extensionMethods = make(map[string]map[string]*loader.Method)
	}
	mr.extensions[extensionName] = ext
	// Load methods from extension
	var err error
	extMap := make(map[string]*loader.Method)
	structType := reflect.TypeOf(reflect.ValueOf(ext).Interface())
	for i := 0; i < structType.NumMethod(); i++ {
		mthd := structType.Method(i)
		log.Debug("Defined methods: ", "methodName", mthd.Name)
		extMap[mthd.Name], err = loader.NewReflectedMethod(mthd, "extension")
		if err != nil {
			log.Warn("Error creating method type", "error", err)
		}
	}
	mr.extensionMethods[extensionName] = extMap
	return nil
}

// ModelRepository default implementation of model repository
type ModelRepository[T any, K PrimaryKeys] struct {
	BaseRepository[T]
	env       *loader.Environment
	db        *gorm.DB
	tableName string
	model     *loader.Model
}

func (mr ModelRepository[T, K]) connection() *gorm.DB {
	if mr.env == nil {
		return loader.GetAdapter().Connector().DB()
	}
	return mr.env.Cr()
}
func (mr ModelRepository[T, K]) validateAndInitialize(modelLoader *ModelLoader) error {
	mdl, err := modelLoader.LoadBaseModel(new(T))
	if err != nil {
		return err
	}
	// Migrate this model
	err = mr.connection().AutoMigrate(*new(T))
	if err != nil {
		return err
	}
	mr.model = mdl
	// Resolve model table name
	mr.tableName = mr.connection().Unscoped().Model(new(T)).Name()
	mr.initialize()
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
func (mr ModelRepository[T, K]) validate() error {
	if mr.env != nil && mr.env.Cr() == nil {
		return errors.New("invalid state: database not initialized")
	}
	return nil
}
func (mr ModelRepository[T, K]) Save(v interface{}) error {
	err := mr.validate()
	if err != nil {
		return err
	}
	err = mr.connection().Save(v).Error
	if err != nil {
		return err
	}
	return nil
}
func (mr ModelRepository[T, K]) Search(cond *conditions.Condition) (interface{}, error) {
	var vv []T
	err := mr.connection().Find(&vv).Error
	if err != nil {
		return nil, err
	}
	return vv, nil
}
func (mr ModelRepository[T, K]) TableName() string {
	return mr.tableName
}
func (mr ModelRepository[T, K]) ModelName() string {
	mdl, ok := mr.GetModel()
	if ok {
		return mdl.Name()
	}
	return ""
}
func (mr ModelRepository[T, K]) Methods() *loader.MethodsCollection {
	return mr.model.Methods()
}

func (mr ModelRepository[T, K]) Query(env *loader.Environment) *loader.RecordCollection {
	var rc *loader.RecordCollection
	if env == nil {
		rc = loader.NewRecordCollection(mr.env, mr.model)
	} else {
		rc = loader.NewRecordCollection(env, mr.model)
	}
	return rc
}

func (mr ModelRepository[T, K]) Condition() *conditions.ModelConditionStart {
	return &conditions.ModelConditionStart{&conditions.ConditionStart{}}
}

func (mr ModelRepository[T, K]) Fields() *loader.FieldsCollection {
	return mr.Fields()
}

func (mr ModelRepository[T, K]) GetFieldName(s string) conditions.FieldName {
	dd, ok := mr.model.Fields().Get(s)
	if !ok {
		return nil
	}
	return conditions.NewFieldName(dd.Name(), dd.JSON())
}

func (mr ModelRepository[T, K]) Delete(v interface{}) (interface{}, error) {
	err := mr.connection().Delete(&v).Error
	if err != nil {
		return *new(T), err
	}
	return v, nil
}

func (mr ModelRepository[T, K]) FindById(id K) (interface{}, error) {
	var v T
	err := mr.connection().First(&v, id).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (mr ModelRepository[T, K]) FindByIdWithOptions(id K, options RecordOptions) (interface{}, error) {
	var v T
	db := mr.connection()
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

func (mr ModelRepository[T, K]) FindByIds(ids []K) (interface{}, error) {
	var v []T
	err := mr.connection().Find(&v, ids).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (mr ModelRepository[T, K]) FindByIdsWithOptions(id []K, options RecordOptions) (any, error) {
	var v []T
	db := mr.connection()
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

func (mr ModelRepository[T, K]) registerExtension(d interface{}) error {
	return mr.RegisterExtension(d)
}
