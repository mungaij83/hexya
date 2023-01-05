package models

import (
	"database/sql"
	"fmt"
	"github.com/hexya-erp/hexya/src/i18n"
	"github.com/hexya-erp/hexya/src/models/fieldtype"
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/models/operator"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/tools/nbutils"
	"reflect"
	"strings"
)

type ModelExtension[T any] interface {
	ExtensionName() string
}

// DefaultMixinExtension defines default mixin extension where T defines the model being extended
type DefaultMixinExtension[T any] struct {
	adapter loader.DbAdapter
}

func (DefaultMixinExtension[T]) ExtensionName() string {
	return "default_mixin"
}

func (DefaultMixinExtension[T]) New(rc *loader.RecordCollection, data loader.RecordData) *loader.RecordCollection {
	return rc.New(data)
}

// Create inserts a record in the database from the given data.
// Returns the created RecordCollection.
func (DefaultMixinExtension[T]) Create(rc *loader.RecordCollection, data loader.RecordData) *loader.RecordCollection {
	return rc.Create(data)
}

// Read reads the database and returns a slice of FieldMap of the given model.
func (DefaultMixinExtension[T]) Read(rc *loader.RecordCollection, fields loader.FieldNames) []loader.RecordData {
	var res []loader.RecordData
	// Check if we have id in fields, and add it otherwise
	fields = loader.AddIDIfNotPresent(fields)
	// Do the actual reading
	for _, rec := range rc.Records() {
		fData := loader.NewModelData(rc.Model())
		for _, fName := range fields {
			fData.Underlying().Set(fName, rec.Get(fName))
		}
		res = append(res, fData)
	}
	return res
}

// Load looks up cache for fields of the RecordCollection and
// query database for missing values.
// fields are the fields to retrieve in the expression format,
// i.e. "User.Profile.Age" or "user_id.profile_id.age".
// If no fields are given, all DB columns of the RecordCollection's
// model are retrieved.
func (DefaultMixinExtension[T]) Load(rc *loader.RecordCollection, fields ...loader.FieldName) *loader.RecordCollection {
	return rc.Load(fields...)
}

// Write is the base implementation of the 'Write' method which updates
// records in the database with the given data.
// Data can be either a struct pointer or a FieldMap.`,
func (DefaultMixinExtension[T]) Write(rc *loader.RecordCollection, data loader.RecordData) bool {
	return rc.Update(data)
}

// Unlink deletes the given records in the database.
func (DefaultMixinExtension[T]) Unlink(rc *loader.RecordCollection) int64 {
	return rc.Unlink()
}

// BrowseOne returns a new RecordSet with only the record with the given id.
// Note that this function is just a shorcut for Search on a given id.
func (DefaultMixinExtension[T]) BrowseOne(rc *loader.RecordCollection, id int64) *loader.RecordCollection {
	return rc.Call("Search", rc.Model().Field(loader.ID).Equals(id)).(loader.RecordSet).Collection()
}

// SearchCount fetch from the database the number of records that match the RecordSet conditions.
func (DefaultMixinExtension[T]) SearchCount(rc *loader.RecordCollection) int {
	return rc.SearchCount()
}

// CopyData copies given record's data with all its fields values.
//
// overrides contains field values to override in the original values of the copied record.
func (DefaultMixinExtension[T]) CopyData(rc *loader.RecordCollection, overrides loader.RecordData) *loader.ModelData {
	rc.EnsureOne()
	// Handle case when overrides is nil
	oVal := reflect.ValueOf(overrides)
	if !oVal.IsValid() || (oVal.Kind() != reflect.Struct && oVal.IsNil()) {
		overrides = loader.NewModelDataFromRS(rc)
	}

	// Create the RecordData
	res := loader.NewModelDataFromRS(rc)
	for _, fi := range rc.Model().Fields().NameRegistry() {
		fName := rc.Model().FieldName(fi.Name())
		if overrides.Underlying().Has(fName) {
			// Overrides are applied below
			continue
		}
		if fi.NoCopy || fi.IsComputedField() {
			continue
		}
		switch fi.FieldType {
		case fieldtype.One2One:
			// One2one related records must be copied to avoid duplicate keys on FK
			res = res.Create(fName, rc.Get(fName).(loader.RecordSet).Collection().Call("CopyData", nil).(loader.RecordData).Underlying())
		case fieldtype.One2Many, fieldtype.Rev2One:
			for _, rec := range rc.Get(fName).(loader.RecordSet).Collection().Records() {
				res = res.Create(fName, rec.Call("CopyData", nil).(loader.RecordData).Underlying().Unset(fi.RelatedModel.FieldName(fi.ReverseFK)))
			}
		default:
			res.Set(fName, rc.Get(fName))
		}
	}
	// Apply overrides
	res.RemovePK()
	res.MergeWith(overrides.Underlying())
	return res
}

// Copy duplicates the given records.
//
// overrides contains field values to override in the original values of the copied record.`,
func (DefaultMixinExtension[T]) Copy(rc *loader.RecordCollection, overrides loader.RecordData) *loader.RecordCollection {
	rc.EnsureOne()
	data := rc.Call("CopyData", overrides).(loader.RecordData).Underlying()
	newRs := rc.Call("Create", data).(loader.RecordSet).Collection()
	return newRs
}

// NameGet retrieves the human readable name of this record.`,
func (DefaultMixinExtension[T]) NameGet(rc *loader.RecordCollection) string {
	if _, nameExists := rc.Model().Fields().Get("Name"); nameExists {
		switch name := rc.Get(rc.Model().FieldName("Name")).(type) {
		case string:
			return name
		case fmt.Stringer:
			return name.String()
		default:
			log.Panic("Name field is neither a string nor a fmt.Stringer", "model", rc.Model())
		}
	}
	return rc.String()
}

// SearchByName searches for records that have a display name matching the given
// "name" pattern when compared with the given "op" operator, while also
// matching the optional search condition ("additionalCond").
//
// This is used for example to provide suggestions based on a partial
// value for a relational field. Sometimes be seen as the inverse
// function of NameGet but it is not guaranteed to be.
func (DefaultMixinExtension[T]) SearchByName(rc *loader.RecordCollection, name string, op operator.Operator, additionalCond loader.Conditioner, limit int) *loader.RecordCollection {
	if op == "" {
		op = operator.IContains
	}
	cond := rc.Model().Field(rc.Model().FieldName("Name")).AddOperator(op, name)
	if !additionalCond.Underlying().IsEmpty() {
		cond = cond.AndCond(additionalCond.Underlying())
	}
	return rc.Model().Search(rc.Env(), cond).Limit(limit)
}

// FieldsGet returns the definition of each field.
// The embedded fields are included.
// The string, help, and selection (if present) attributes are translated.
//
// The result map is indexed by the fields JSON names.
func (DefaultMixinExtension[T]) FieldsGet(rc *loader.RecordCollection, args loader.FieldsGetArgs) map[string]*loader.FieldInfo {
	// Get the field informations
	res := rc.Model().FieldsGet(args.Fields...)

	// Translate attributes when required
	lang := rc.Env().Context().GetString("lang")
	for fName, fInfo := range res {
		res[fName].Help = i18n.Registry.TranslateFieldHelp(lang, rc.Model().Name(), fInfo.Name, fInfo.Help)
		res[fName].String = i18n.Registry.TranslateFieldDescription(lang, rc.Model().Name(), fInfo.Name, fInfo.String)
		res[fName].Selection = i18n.Registry.TranslateFieldSelection(lang, rc.Model().Name(), fInfo.Name, fInfo.Selection)
	}
	return res
}

// FieldGet returns the definition of the given field.
// The string, help, and selection (if present) attributes are translated.
func (DefaultMixinExtension[T]) FieldGet(rc *loader.RecordCollection, field loader.FieldName) *loader.FieldInfo {
	args := loader.FieldsGetArgs{
		Fields: []loader.FieldName{field},
	}
	return rc.Call("FieldsGet", args).(map[string]*loader.FieldInfo)[field.JSON()]
}

// DefaultGet returns a Params map with the default values for the model.
func (DefaultMixinExtension[T]) DefaultGet(rc *loader.RecordCollection) *loader.ModelData {
	res := rc.GetDefaults(rc.Env().Context().GetBool("hexya_ignore_computed_defaults"))
	return res
}

// CheckRecursion verifies that there is no loop in a hierarchical structure of records,
// by following the parent relationship using the 'Parent' field until a loop is detected or
// until a top-level record is found.
//
// It returns true if no loop was found, false otherwise`,
func (mx DefaultMixinExtension[T]) CheckRecursion(rc *loader.RecordCollection) bool {
	if _, exists := rc.Model().Fields().Get("Parent"); !exists {
		// No Parent field in model, so no loop
		return true
	}
	if rc.HasNegIds {
		// We have a negative id, so we can't have a loop
		return true
	}
	// We use direct SQL query to bypass access control
	query := fmt.Sprintf(`SELECT parent_id FROM %s WHERE id = ?`, mx.adapter.QuoteTableName(rc.Model().TableName()))
	rc.Load(rc.Model().FieldName("Parent"))
	for _, record := range rc.Records() {
		currentID := record.Ids()[0]
		for {
			var parentID sql.NullInt64
			rc.Env().Cursor().Get(&parentID, query, currentID)
			if !parentID.Valid {
				break
			}
			currentID = parentID.Int64
			if currentID == record.Ids()[0] {
				return false
			}
		}
	}
	return true
}

// Onchange returns the values that must be modified according to each field's Onchange
// method in the pseudo-record given as params.Values`,
func (DefaultMixinExtension[T]) OnChange(rc *loader.RecordCollection, params loader.OnchangeParams) loader.OnchangeResult {
	var retValues *loader.ModelData
	var warnings []string
	filters := make(map[loader.FieldName]loader.Conditioner)

	err := loader.SimulateInNewEnvironment(rc.Env().Uid(), func(env loader.Environment) {
		values := params.Values.Underlying().FieldMap
		data := loader.NewModelDataFromRS(rc.WithEnv(env), values)
		if rc.IsNotEmpty() {
			data.Set(loader.ID, rc.Ids()[0])
		}
		retValues = loader.NewModelDataFromRS(rc.WithEnv(env))
		var rs *loader.RecordCollection
		if id, _ := nbutils.CastToInteger(data.Get(loader.ID)); id != 0 {
			rs = rc.WithEnv(env).WithIds([]int64{id})
			rs = rs.WithContext("hexya_onchange_origin", rs.First().Wrap())
			rs.WithContext("hexya_force_compute_write", true).Update(data)
		} else {
			rs = rc.WithEnv(env).WithContext("hexya_force_compute_write", true).Create(data)
		}
		// Set inverse fields
		for field := range values {
			fName := rs.Model().FieldName(field)
			fi := rs.Model().GetRelatedFieldInfo(fName)
			if fi.Inverse != "" {
				fVal := data.Get(fName)
				rs.Call(fi.Inverse, fVal)
			}
		}
		todo := params.Fields
		done := make(map[string]bool)
		// Apply onchanges or compute
		for len(todo) > 0 {
			field := todo[0]
			todo = todo[1:]
			if done[field.JSON()] {
				continue
			}
			done[field.JSON()] = true
			if params.Onchange[field.Name()] == "" && params.Onchange[field.JSON()] == "" {
				continue
			}
			fi := rs.Model().GetRelatedFieldInfo(field)
			fnct := fi.OnChange
			if fnct == "" {
				fnct = fi.Compute
			}
			rrs := rs
			toks := loader.SplitFieldNames(field, loader.ExprSep)
			if len(toks) > 1 {
				rrs = rs.Get(loader.JoinFieldNames(toks[:len(toks)-1], loader.ExprSep)).(loader.RecordSet).Collection()
			}
			// Values
			if fnct != "" {
				vals := rrs.Call(fnct).(loader.RecordData)
				for _, f := range vals.Underlying().FieldNames() {
					if !done[f.JSON()] {
						todo = append(todo, f)
					}
				}
				rrs.WithContext("hexya_force_compute_write", true).Call("Write", vals)
			}
			// Warning
			if fi.OnChangeWarning != "" {
				w := rrs.Call(fi.OnChangeWarning).(string)
				if w != "" {
					warnings = append(warnings, w)
				}
			}
			// Filters
			if fi.OnChangeFilters != "" {
				ff := rrs.Call(fi.OnChangeFilters).(map[loader.FieldName]loader.Conditioner)
				for k, v := range ff {
					filters[k] = v
				}
			}
		}
		// Collect modified values
		for field, val := range values {
			fName := rs.Model().FieldName(field)
			if fName.JSON() == "__last_update" {
				continue
			}
			fi := rs.Collection().Model().GetRelatedFieldInfo(fName)
			newVal := rs.Get(fName)
			switch {
			case fi.FieldType.IsRelationType():
				v := rs.ConvertToRecordSet(val, fi.RelatedModelName)
				nv := rs.ConvertToRecordSet(newVal, fi.RelatedModelName)
				if !v.Equals(nv) {
					retValues.Set(fName, newVal)
				}
			default:
				if val != newVal {
					retValues.Set(fName, newVal)
				}
			}
		}
	})
	if err != nil {
		panic(err)
	}
	retValues.Unset(loader.ID)
	return loader.OnchangeResult{
		Value:   retValues,
		Warning: strings.Join(warnings, "\n\n"),
		Filters: filters,
	}
}

// Search returns a new RecordSet filtering on the current one with the
// additional given Condition.
func (DefaultMixinExtension[T]) Search(rc *loader.RecordCollection, cond loader.Conditioner) *loader.RecordCollection {
	return rc.Search(cond.Underlying())
}

// Browse returns a new RecordSet with only the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (DefaultMixinExtension[T]) Browse(rc *loader.RecordCollection, ids []int64) *loader.RecordCollection {
	return rc.Call("Search", rc.Model().Field(loader.ID).In(ids)).(loader.RecordSet).Collection()
}

// Fetch query the database with the current filter and returns a RecordSet
// with the queries ids.
//
// Fetch is lazy and only return ids. Use Load() instead if you want to fetch all fields.
func (DefaultMixinExtension[T]) Fetch(rc *loader.RecordCollection) *loader.RecordCollection {
	return rc.Fetch()
}

// SearchAll returns a RecordSet with all items of the table, regardless of the
// current RecordSet query. It is mainly meant to be used on an empty RecordSet.
func (DefaultMixinExtension[T]) SearchAll(rc *loader.RecordCollection) *loader.RecordCollection {
	return rc.SearchAll()
}

// GroupBy returns a new RecordSet grouped with the given GROUP BY expressions.
func (DefaultMixinExtension[T]) GroupBy(rc *loader.RecordCollection, exprs ...loader.FieldName) *loader.RecordCollection {
	return rc.GroupBy(exprs...)
}

// Limit returns a new RecordSet with only the first 'limit' records.
func (DefaultMixinExtension[T]) Limit(rc *loader.RecordCollection, limit int) *loader.RecordCollection {
	return rc.Limit(limit)
}

// Offset returns a new RecordSet with only the records starting at offset
func (DefaultMixinExtension[T]) Offset(rc *loader.RecordCollection, offset int) *loader.RecordCollection {
	return rc.Offset(offset)
}

// OrderBy returns a new RecordSet ordered by the given ORDER BY expressions.
// Each expression contains a field name and optionally one of "asc" or "desc", such as:
//
// rs.OrderBy("Company", "Name desc")
func (DefaultMixinExtension[T]) OrderBy(rc *loader.RecordCollection, exprs ...string) *loader.RecordCollection {
	return rc.OrderBy(exprs...)
}

// Union returns a new RecordSet that is the union of this RecordSet and the given
// "other" RecordSet. The result is guaranteed to be a set of unique records.
func (DefaultMixinExtension[T]) Union(rc *loader.RecordCollection, other loader.RecordSet) *loader.RecordCollection {
	return rc.Union(other)
}

// Subtract returns a RecordSet with the Records that are in this
// RecordCollection but not in the given 'other' one.
// The result is guaranteed to be a set of unique records.
func (DefaultMixinExtension[T]) Subtract(rc *loader.RecordCollection, other loader.RecordSet) *loader.RecordCollection {
	return rc.Subtract(other)
}

// Intersect returns a new RecordCollection with only the records that are both
// in this RecordCollection and in the other RecordSet.
func (DefaultMixinExtension[T]) Intersect(rc *loader.RecordCollection, other loader.RecordSet) *loader.RecordCollection {
	return rc.Intersect(other)
}

// CartesianProduct returns the cartesian product of this RecordCollection with others.
func (DefaultMixinExtension[T]) CartesianProduct(rc *loader.RecordCollection, other ...loader.RecordSet) []*loader.RecordCollection {
	return rc.CartesianProduct(other...)
}

// Equals returns true if this RecordSet is the same as other
// i.e. they are of the same model and have the same ids
func (DefaultMixinExtension[T]) Equals(rc *loader.RecordCollection, other loader.RecordSet) bool {
	return rc.Equals(other)
}

// Sorted returns a new RecordCollection sorted according to the given less function.
//
// The less function should return true if rs1 < rs2`,
func (DefaultMixinExtension[T]) Sorted(rc *loader.RecordCollection, less func(rs1 loader.RecordSet, rs2 loader.RecordSet) bool) *loader.RecordCollection {
	return rc.Sorted(less)
}

// SortedDefault returns a new record set with the same records as rc but sorted according
// to the default order of this model
func (DefaultMixinExtension[T]) SortedDefault(rc *loader.RecordCollection) *loader.RecordCollection {
	return rc.SortedDefault()
}

// SortedByField returns a new record set with the same records as rc but sorted by the given field.
// If reverse is true, the sort is done in reversed order
func (DefaultMixinExtension[T]) SortedByField(rc *loader.RecordCollection, namer loader.FieldName, reverse bool) *loader.RecordCollection {
	return rc.SortedByField(namer, reverse)
}

// Filtered returns a new record set with only the elements of this record set
// for which test is true.
//
// Note that if this record set is not fully loaded, this function will call the database
// to load the fields before doing the filtering. In this case, it might be more efficient
// to search the database directly with the filter condition.
func (DefaultMixinExtension[T]) Filtered(rc *loader.RecordCollection, test func(rs loader.RecordSet) bool) *loader.RecordCollection {
	return rc.Filtered(test)
}

// GetRecord returns the Recordset with the given externalID. It panics if the externalID does not exist.
func (DefaultMixinExtension[T]) GetRecord(rc *loader.RecordCollection, externalID string) *loader.RecordCollection {
	return rc.GetRecord(externalID)
}

// CheckExecutionPermission panics if the current user is not allowed to execute the given method.
//
// If dontPanic is false, this function will panic, otherwise it returns true
// if the user has the execution permission and false otherwise.
func (DefaultMixinExtension[T]) CheckExecutionPermission(rc *loader.RecordCollection, method *loader.Method, dontPanic ...bool) bool {
	return rc.CheckExecutionPermission(method, dontPanic...)
}

// SQLFromCondition returns the WHERE clause sql and arguments corresponding to
// the given condition.`,
func (DefaultMixinExtension[T]) SQLFromCondition(rc *loader.RecordCollection, c *loader.Condition) (string, loader.SQLParams) {
	return rc.SQLFromCondition(c)
}

// WithEnv returns a copy of the current RecordSet with the given Environment.
func (DefaultMixinExtension[T]) WithEnv(rc *loader.RecordCollection, env loader.Environment) *loader.RecordCollection {
	return rc.WithEnv(env)
}

// WithContext returns a copy of the current RecordSet with
// its context extended by the given key and value.
func (DefaultMixinExtension[T]) WithContext(rc *loader.RecordCollection, key string, value interface{}) *loader.RecordCollection {
	return rc.WithContext(key, value)
}

// WithNewContext returns a copy of the current RecordSet with its context
// replaced by the given one.
func (DefaultMixinExtension[T]) WithNewContext(rc *loader.RecordCollection, context *types.Context) *loader.RecordCollection {
	return rc.WithNewContext(context)
}

// Sudo returns a new RecordSet with the given userID
// or the superuser ID if not specified
func (DefaultMixinExtension[T]) Sudo(rc *loader.RecordCollection, userID ...int64) *loader.RecordCollection {
	return rc.Sudo(userID...)
}
