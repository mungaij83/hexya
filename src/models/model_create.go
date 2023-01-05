package models

import (
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/hexya/src/tools"
	"reflect"
)

// getOrCreateModel checks if the given model has been created
// and creates it if it is not the case/
func getOrCreateModel(name string, options tools.Option) *loader.Model {
	model, ok := Registry.Get(name)
	if !ok {
		model = loader.CreateModel(name, options)
	}
	if model.Created {
		log.Panic("Trying to add already existing model", "model", name)
	}
	model.Created = true
	return model
}

func GetRepository[T any]() Repository[any, int64] {
	return Registry.getRepo(new(T))
}

// NewModel creates a new model with the given name.
func NewModel(name string) *loader.Model {
	model := getOrCreateModel(name, 0)
	if model == nil {
		log.Warn("failed to register model", "Error", "Failed to create")
		return nil
	}
	//model.InheritModel(Registry.MustGet("ModelMixin"))
	return model
}

// NewMixinModel creates a new mixin model with the given name.
func NewMixinModel(name string) *loader.Model {
	model := getOrCreateModel(name, tools.MixinModel)
	return model
}

// NewTransientModel creates a new mixin model with the given name.
func NewTransientModel(name string) *loader.Model {
	model := getOrCreateModel(name, tools.TransientModel)
	//model.InheritModel(Registry.MustGet("BaseMixin"))
	return model
}

// NewManualModel creates a model whose table is not automatically generated
// in the database. This is particularly useful for SQL view models.
func NewManualModel(name string) *loader.Model {
	model := getOrCreateModel(name, tools.ManualModel)
	//model.InheritModel(Registry.MustGet("CommonMixin"))
	return model
}

// RegisterRecordSetWrapper registers the object passed as obj as the RecordSet type
// for the given model.
//
// - typ must be a struct that embeds *RecordCollection
// - modelName must be the name of a model that exists in the registry
func RegisterRecordSetWrapper(modelName string, obj interface{}) {
	Registry.MustGet(modelName)
	typ := reflect.TypeOf(obj)
	if typ.Kind() != reflect.Struct {
		log.Panic("trying to register a non struct type as Wrapper", "modelName", modelName, "type", typ)
	}
	if typ.Field(0).Type != reflect.TypeOf(new(loader.RecordCollection)) {
		log.Panic("trying to register a struct that don't embed *RecordCollection", "modelName", modelName, "type", typ)
	}
	loader.RecordSetWrappers[modelName] = typ
}

// FreeTransientModels remove transient models records from database which are
// older than the given timeout.
func FreeTransientModels() {
	for _, model := range Registry.registryByTableName {
		if model.IsTransient() {
			err := loader.ExecuteInNewEnvironment(security.SuperUserID, func(env loader.Environment) {
				//createDate := model.FieldName("CreateDate")
				//model.Search(env, model.GetField(createDate).Lower(dates.Now().Add(-transientModelTimeout))).Call("Unlink")
			})
			if err != nil {
				log.Warn("Failed to free transient models")
			}
		}
	}
}
