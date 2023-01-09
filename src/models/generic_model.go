// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package models

import (
	"github.com/hexya-erp/hexya/src/models/conditions"
	"github.com/hexya-erp/hexya/src/models/loader"
)

// ------- MODEL ---------

// CompanyModel is a strongly typed model definition that is used
// to extend the Company model or to get a CompanySet through
// its NewSet() function.
//
// To get the unique instance of this type, call Company().
type ModelDefinition[M any] struct {
	*loader.Model
}

// NewSet returns a new CompanySet instance in the given Environment
func (md ModelDefinition[M]) NewSet(env loader.Environment) conditions.RecordSet {
	return env.Pool(md.Name())
}
func (md ModelDefinition[M]) Create(env loader.Environment, data *loader.ModelData) *loader.RecordCollection {
	return md.Model.Create(env, data)
}

func (md ModelDefinition[M]) Save(env loader.Environment, data M) *loader.RecordCollection {
	return md.Model.Create(env, data)
}
func (md ModelDefinition[M]) wrapBaseType(coll *loader.RecordCollection) *M {
	return nil
}

// Create creates a new Company record and returns the newly created
// CompanySet instance.
func (md ModelDefinition[M]) CreateModel(env loader.Environment, data *loader.ModelData) *M {
	dd := md.Create(env, data)
	if dd != nil {
		return nil
	}
	return md.wrapBaseType(dd)
}

// Search searches the database and returns a new CompanySet instance
// with the records found.
func (md ModelDefinition[M]) Search(env loader.Environment, cond conditions.Conditioner) *loader.RecordCollection {
	return md.Model.Search(env, cond)
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md ModelDefinition[M]) Browse(env loader.Environment, ids []int64) *loader.RecordCollection {
	return md.Model.Browse(env, ids)
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md ModelDefinition[M]) BrowseOne(env loader.Environment, id int64) *loader.RecordCollection {
	return md.Model.BrowseOne(env, id)
}

// NewData returns a pointer to a new empty CompanyData instance.
//
// Optional field maps if given will be used to populate the data.
func (md ModelDefinition[M]) NewData(fm ...loader.FieldMap) *loader.ModelData {
	return loader.NewModelData(md.Model, fm...)
}

// Fields returns the Field Collection of the Company Model
func (md ModelDefinition[M]) Fields() loader.FieldsCollections {
	return loader.FieldsCollections{
		FieldsCollection: md.Model.Fields(),
	}
}

//func (ModelDefinition[M]) Query() conditions.ConditionStart {
//	return conditions.ConditionStart{
//		ConditionStart: &ConditionStart{},
//	}
//}

// Methods returns the Method Collection of the Company Model
func (md ModelDefinition[M]) Methods() *loader.MethodsCollection {
	return md.Model.Methods()
}

// Underlying returns the underlying models.Model instance
func (md ModelDefinition[M]) Underlying() *loader.Model {
	return md.Model
}

// Coalesce takes a list of CompanySet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md ModelDefinition[M]) Coalesce(lst ...conditions.RecordSet) *conditions.RecordSet {
	var last conditions.RecordSet
	for _, elem := range lst {
		if elem.Collection().(*loader.RecordCollection).IsNotEmpty() {
			return &elem
		}
		last = elem
	}
	return &last
}

// Company returns the unique instance of the CompanyModel type
// which is used to extend the Company model or to get a CompanySet through
// its NewSet() function.
func NewModelDefinition[M any](mdl interface{}) ModelDefinition[M] {
	return NewTypedModel[M](mdl)
}
