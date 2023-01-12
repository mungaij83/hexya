package loader

import (
	"database/sql"
	"fmt"
	"github.com/hexya-erp/hexya/src/models/conditions"
	"github.com/hexya-erp/hexya/src/tools"
	"github.com/hexya-erp/hexya/src/tools/strutils"
	"github.com/hexya-erp/hexya/src/tools/typesutils"
	"reflect"
	"strings"
)

// A Model is the definition of a business object (e.g. a partner, a sale order, etc.)
// including fields and methods.
type Model struct {
	name            string
	options         tools.Option
	rulesRegistry   *recordRuleRegistry
	tableName       string
	fields          *FieldsCollection
	methods         *MethodsCollection
	mixins          []*Model
	sqlConstraints  map[string]sqlConstraint
	sqlErrors       map[string]string
	defaultOrderStr []string
	defaultOrder    []orderPredicate
	Created         bool
}

// An sqlConstraint holds the data needed to create a table constraint in the database
type sqlConstraint struct {
	name        string
	sql         string
	errorString string
}

// Name returns the name of this model
func (m *Model) Name() string {
	log.Debug("Model Name() is:", "name", m)
	if len(m.name) > 0 {
		return m.name
	}
	return ""
}

func (m *Model) Options() tools.Option {
	return m.options
}

func (m *Model) Mixins() []*Model {
	return m.mixins
}

// getRelatedModelInfo returns the Model of the related model when
// following path.
// - If skipLast is true, getRelatedModelInfo does not follow the last part of the path
// - If the last part of path is a non relational field, it is simply ignored, whatever
// the value of skipLast.
func (m *Model) getRelatedModelInfo(path conditions.FieldName, skipLast ...bool) *Model {
	if path == nil {
		return m
	}
	var skip bool
	if len(skipLast) > 0 {
		skip = skipLast[0]
	}

	exprs := conditions.SplitFieldNames(path, conditions.ExprSep)
	fi := m.fields.MustGet(exprs[0].JSON())
	if fi.RelatedModel == nil || (len(exprs) == 1 && skip) {
		// The field is a non relational field, so we are already
		// on the related Model. Or we have only 1 exprs and we skip the last one.
		return m
	}
	if len(exprs) > 1 {
		return fi.RelatedModel.getRelatedModelInfo(conditions.JoinFieldNames(exprs[1:], conditions.ExprSep), skipLast...)
	}
	return fi.RelatedModel
}

// getRelatedFieldIfo returns the Field of the related field when
// following path. Path can be formed from field names or JSON names.
func (m *Model) GetRelatedFieldInfo(path conditions.FieldName) *Field {
	colExprs := conditions.SplitFieldNames(path, conditions.ExprSep)
	var rmi *Model
	num := len(colExprs)
	if len(colExprs) > 1 {
		rmi = m.getRelatedModelInfo(path, true)
	} else {
		rmi = m
	}
	fi := rmi.fields.MustGet(colExprs[num-1].JSON())
	return fi
}

// scanToFieldMap scans the connector query result r into the given FieldMap.
// Unlike slqx.MapScan, the returned interface{} values are of the type
// of the Model fields instead of the database types.
//
// substs is a map for substituting field names in the ColScanner if necessary (typically if length is over 64 chars).
// Keys are the alias used in the query, and values are '__' separated paths such as "user_id__profile_id__age"
func (m *Model) scanToFieldMap(r *sql.Rows, dest *FieldMap, substs map[string]string) error {
	columns, err := r.Columns()
	if err != nil {
		return err
	}

	// Step 1: We create a []interface{} which is in fact a []*interface{}
	// and we scan our DB row into it. This enables us to Get null values
	// without panic, since null values will map to nil.
	dbValues := make([]interface{}, len(columns))
	for i := range dbValues {
		dbValues[i] = new(interface{})
	}

	err = r.Scan(dbValues...)
	if err != nil {
		return err
	}

	// Step 2: We populate our dest FieldMap with these values
	for i, dbValue := range dbValues {
		colName := columns[i]
		if s, ok := substs[colName]; ok {
			colName = s
		}
		colName = strings.Replace(colName, conditions.SqlSep, conditions.ExprSep, -1)
		dbVal := reflect.ValueOf(dbValue).Elem().Interface()
		(*dest)[colName] = dbVal
	}

	// Step 3: We convert values with the type of the corresponding Field
	// if the value is not nil.
	m.convertValuesToFieldType(dest, false)
	return r.Err()
}

// convertValuesToFieldType converts all values of the given FieldMap to
// their type in the Model.
//
// If this method is used to convert values before writing to DB, you
// should set writeDB to true.
func (m *Model) convertValuesToFieldType(fMap *FieldMap, writeDB bool) {
	destVals := reflect.ValueOf(fMap).Elem()
	for colName, fMapValue := range *fMap {
		if val, ok := fMapValue.(bool); ok && !val {
			// Hack to manage client returning false instead of nil
			fMapValue = nil
		}
		fi := m.GetRelatedFieldInfo(m.FieldName(colName))
		fType := fi.structField.Type
		typedValue := reflect.New(fType).Interface()
		err := typesutils.Convert(fMapValue, typedValue, fi.IsRelationField())
		if err != nil {
			log.Panic(err.Error(), "model", m.name, "field", colName, "type", fType, "value", fMapValue)
		}
		destVals.SetMapIndex(reflect.ValueOf(colName), reflect.ValueOf(typedValue).Elem())
	}
	if writeDB {
		// Change zero values to NULL if writing to DB when applicable
		for colName, fMapValue := range *fMap {
			fi := m.GetRelatedFieldInfo(m.FieldName(colName))
			val := reflect.ValueOf(fMapValue)
			switch {
			case fi.FieldType.IsFKRelationType() && val.Kind() == reflect.Int64 && val.Int() == 0:
				val = reflect.ValueOf((*interface{})(nil))
				destVals.SetMapIndex(reflect.ValueOf(colName), val)
			}
		}
	}
}

// AddFields adds the given fields to the model.
func (m *Model) AddFields(fields map[string]FieldDefinition) {
	if m.fields.Model() == nil {
		m.fields.model = m
	}
	for name, field := range fields {
		if field == nil {
			log.Warn("models.Field is not defined", "model", m.name, "field", name)
			continue
		}
		log.Debug("Add declared Field: ", "model", m.name, "field", name)
		newField := field.DeclareField(m.fields, name)
		if _, exists := m.fields.Get(name); exists {
			log.Panic("models.Field already exists", "model", m.name, "field", name)
		}
		if newField.model == nil {
			newField.model = m
		}
		m.fields.add(newField)
	}
}

// IsMixin returns true if this is a mixin model.
func (m *Model) IsMixin() bool {
	if m.options&tools.MixinModel > 0 {
		return true
	}
	return false
}

// IsManual returns true if this is a manual model.
func (m *Model) IsManual() bool {
	if m.options&tools.ManualModel > 0 {
		return true
	}
	return false
}

// isSystem returns true if this is a system model.
func (m *Model) isSystem() bool {
	if m.options&tools.SystemModel > 0 {
		return true
	}
	return false
}

// isContext returns true if this is a context model.
func (m *Model) isContext() bool {
	if m.options&tools.ContextsModel > 0 {
		return true
	}
	return false
}

// IsM2MLink returns true if this is an M2M Link model.
func (m *Model) IsM2MLink() bool {
	if m.options&tools.Many2ManyLinkModel > 0 {
		return true
	}
	return false
}

// IsTransient returns true if this Model is transient
func (m *Model) IsTransient() bool {
	return m.options == tools.TransientModel
}

// hasParentField returns true if this model is recursive and has a Parent field.
func (m *Model) hasParentField() bool {
	_, parentExists := m.fields.Get("Parent")
	return parentExists
}

// Fields returns the fields collection of this model
func (m *Model) Fields() *FieldsCollection {
	return m.fields
}

// FieldNames returns the slice of all field's names for this model
func (m *Model) FieldNames() conditions.FieldNames {
	return m.fields.allFieldNames()
}

// Methods returns the methods collection of this model
func (m *Model) Methods() *MethodsCollection {
	return m.methods
}

// SetDefaultOrder sets the default order used by this model
// when no OrderBy() is specified in a query. When unspecified,
// default order is 'id asc'.
//
// Give the order fields in separate strings, such as
// model.SetDefaultOrder("Name desc", "date asc", "id")
func (m *Model) SetDefaultOrder(orders ...string) {
	m.defaultOrderStr = orders
}

// ordersFromStrings returns the given order by exprs as a slice of order structs
func (m *Model) ordersFromStrings(exprs []string) []orderPredicate {
	res := make([]orderPredicate, len(exprs))
	for i, o := range exprs {
		toks := strings.Split(o, " ")
		var desc bool
		if len(toks) > 1 && strings.ToLower(toks[1]) == "desc" {
			desc = true
		}
		res[i] = orderPredicate{field: m.FieldName(toks[0]), desc: desc}
	}
	return res
}

// JSONizeFieldName returns the json name of the given fieldName
// If fieldName is already the json name, returns it without modifying it.
// fieldName may be a dot separated path from this model.
// It panics if the path is invalid.
func (m *Model) JSONizeFieldName(fieldName string) string {
	return jsonizePath(m, fieldName)
}

// FieldName returns a FieldName for the field with the given name.
// name may be a dot separated path from this model.
// It returns nil if the name is empty and panics if the path is invalid.
func (m *Model) FieldName(name string) conditions.FieldName {
	if name == "" {
		return nil
	}
	jsonName := jsonizePath(m, name)
	return conditions.NewFieldName(name, jsonName)
}

// Field starts a condition on this model
func (m *Model) Field(name conditions.FieldName) *conditions.ConditionField {
	newExprs := conditions.SplitFieldNames(name, conditions.ExprSep)
	cp := conditions.ConditionField{}
	cp.Exprs = append(cp.Exprs, newExprs...)
	return &cp
}

// FieldsGet returns the definition of each field.
// The embedded fields are included.
//
// If no fields are given, then all fields are returned.
//
// The result map is indexed by the fields JSON names.
func (m *Model) FieldsGet(fields ...conditions.FieldName) map[string]*FieldInfo {
	if len(fields) == 0 {
		for n := range m.fields.registryByName {
			fields = append(fields, m.FieldName(n))
		}
	}
	res := make(map[string]*FieldInfo)
	for _, f := range fields {
		fInfo := m.fields.MustGet(f.Name())
		var relation string
		if fInfo.RelatedModel != nil {
			relation = fInfo.RelatedModel.name
		}
		var filter interface{}
		if fInfo.filter != nil {
			filter = fInfo.filter.Serialize()
		}
		_, translate := fInfo.contexts["lang"]
		res[fInfo.json] = &FieldInfo{
			Name:          fInfo.name,
			JSON:          fInfo.json,
			Help:          fInfo.help,
			Searchable:    true,
			Depends:       fInfo.depends,
			Sortable:      true,
			Type:          fInfo.FieldType,
			Store:         fInfo.isSettable(),
			String:        fInfo.description,
			Relation:      relation,
			Selection:     fInfo.selection,
			Domain:        filter,
			ReverseFK:     fInfo.JsonReverseFK,
			OnChange:      fInfo.OnChange != "",
			Translate:     translate,
			InvisibleFunc: fInfo.invisibleFunc,
			ReadOnly:      fInfo.isReadOnly(),
			ReadOnlyFunc:  fInfo.readOnlyFunc,
			Required:      fInfo.required,
			RequiredFunc:  fInfo.requiredFunc,
			DefaultFunc:   fInfo.defaultFunc,
			GoType:        fInfo.structField.Type,
			Index:         fInfo.index,
		}
	}
	return res
}

// FilteredOn adds a condition with a table join on the given field and
// filters the result with the given condition
func (m *Model) FilteredOn(field conditions.FieldName, condition *conditions.Condition) *conditions.Condition {
	res := conditions.Condition{Predicates: make([]conditions.ConditionPredicate, len(condition.Predicates))}
	i := 0
	for _, p := range condition.Predicates {
		p.Exprs = append([]conditions.FieldName{field}, p.Exprs...)
		res.Predicates[i] = p
		i++
	}
	return &res
}

// Create creates a new record in this model with the given data.
func (m *Model) Create(env Environment, data interface{}) *RecordCollection {
	return env.Pool(m.name).Call("Create", data).(conditions.RecordSet).Collection().(*RecordCollection)
}

// Search searches the database and returns records matching the given condition.
func (m *Model) Search(env Environment, cond conditions.Conditioner) *RecordCollection {
	return env.Pool(m.name).Call("Search", cond).(conditions.RecordSet).Collection().(*RecordCollection)
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (m *Model) Browse(env Environment, ids []int64) *RecordCollection {
	return env.Pool(m.name).Call("Browse", ids).(conditions.RecordSet).Collection().(*RecordCollection)
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search the given id.
func (m *Model) BrowseOne(env Environment, id int64) *RecordCollection {
	return env.Pool(m.name).Call("BrowseOne", id).(conditions.RecordSet).Collection().(*RecordCollection)
}

// AddSQLConstraint adds a table constraint in the database.
//   - name is an arbitrary name to reference this constraint. It will be appended by
//     the table name in the database, so there is only need to ensure that it is unique
//     in this model.
//   - sql is constraint definition to pass to the database.
//   - errorString is the text to display to the user when the constraint is violated
func (m *Model) AddSQLConstraint(name, sql, errorString string) {
	constraintName := fmt.Sprintf("%s_%s_mancon", name, m.tableName)
	m.sqlConstraints[constraintName] = sqlConstraint{
		name:        constraintName,
		sql:         sql,
		errorString: errorString,
	}
}

// RemoveSQLConstraint removes the sql constraint with the given name from the database.
func (m *Model) RemoveSQLConstraint(name string) {
	delete(m.sqlConstraints, fmt.Sprintf("%s_mancon", name))
}

// TableName return the connector table name
func (m *Model) TableName() string {
	return m.tableName
}

// Underlying returns the underlying Model data object, i.e. itself
func (m *Model) Underlying() *Model {
	return m
}

// InheritModel extends this Model by importing all fields and methods of mixInModel.
// MixIn methods and fields have a lower priority than those of the model and are
// overridden by the them when applicable.
func (m *Model) InheritModel(mixInModel Modeler) {
	m.mixins = append(m.mixins, mixInModel.Underlying())
}

func (m *Model) Bootstrapped(b bool) {
	m.Bootstrapped(b)
}

// CreateModel creates a new Model with the given name and options.
// You should not use this function directly. Use NewModel instead.
func CreateModel(name string, options tools.Option) *Model {
	mi := &Model{
		name:            name,
		options:         options,
		rulesRegistry:   newRecordRuleRegistry(),
		tableName:       strutils.SnakeCase(name),
		fields:          newFieldsCollection(),
		methods:         newMethodsCollection(),
		sqlConstraints:  make(map[string]sqlConstraint),
		sqlErrors:       make(map[string]string),
		defaultOrderStr: []string{"ID"},
	}
	return mi
}
