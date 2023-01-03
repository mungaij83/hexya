package models

import (
	"errors"
	"fmt"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/tools"
	"github.com/hexya-erp/hexya/src/tools/nbutils"
	"reflect"
	"strconv"
	"strings"
)

var (
	modelLoader ModelLoader
)

func init() {
	modelLoader = ModelLoader{}
}

type TagData struct {
	Many2One  string
	Many2Many string
	One2Many  string
	Value     interface{}
	JSON      string
	Required  bool
	Type      string
	Translate bool
	Index     bool
	Size      int
	NoCopy    bool
	goType    string
	Unique    bool
	Precision int8
	Scale     int8
	Depends   []string
	Options   map[string]string
	ReadOnly  bool
	Stored    bool
	Help      string
	Message   string
}

type ModelLoader struct {
}

func (ml ModelLoader) detectTableName(data interface{}) string {
	fieldTypes := reflect.TypeOf(data)
	tableName := ""
	if strings.HasSuffix(fieldTypes.Name(), "MixinModel") {
		tableName = strings.TrimSuffix(fieldTypes.Name(), "MixinModel")
	} else if strings.HasSuffix(fieldTypes.Name(), "ManualModel") {
		tableName = strings.TrimSuffix(fieldTypes.Name(), "ManualModel")
	} else if strings.HasSuffix(fieldTypes.Name(), "TransientModel") {
		tableName = strings.TrimSuffix(fieldTypes.Name(), "TransientModel")
	} else {
		tableName = strings.TrimSuffix(fieldTypes.Name(), "Model")
	}
	return tableName
}
func (ml ModelLoader) detectTableType(data interface{}) (string, tools.Option) {
	// Determine the type or name of the model
	var optionSetting tools.Option
	tableName := ml.detectTableName(data)
	// Override name of the table from the type interfaces
	switch v := data.(type) {
	case HexyaBaseModel:
		optionSetting = 0
		log.Info("Loading model: ", v.IsModel())
		break
	case HexyaAbstractModel:
		optionSetting = tools.MixinModel
		break
	case HexyaTransientModel:
		optionSetting = tools.TransientModel
		break
	default:
		optionSetting = tools.ManualModel
	}
	return tableName, optionSetting
}

// Model  type is determined with priority as seen in the case statement
// Implementing multiple types will result in other types being less prioritized as per the list below
func (ml ModelLoader) LoadBaseModel(data interface{}) (*loader.Model, error) {
	fieldDefinitions, err := ml.LoadModel(data)
	if err != nil {
		return nil, err
	}
	tableName, option := ml.detectTableType(data)
	var mdl *loader.Model
	switch option {
	case tools.MixinModel:
		mdl = NewMixinModel(tableName)
		break
	case tools.ManualModel:
		mdl = NewManualModel(tableName)
		break
	case tools.TransientModel:
		mdl = NewTransientModel(tableName)
		break
	default:
		mdl = NewModel(tableName)
		break
	}
	// Add Fields and sorting order
	mdl.AddFields(fieldDefinitions)
	switch v := data.(type) {
	case OrderedTableModel:
		orders := v.OrderFields()
		if len(orders) > 0 {
			mdl.SetDefaultOrder(orders...)
		}
		break
	}
	return mdl, nil
}

func (ml ModelLoader) LoadModel(data interface{}) (map[string]loader.FieldDefinition, error) {
	modelFields := make(map[string]loader.FieldDefinition)
	fieldTypes := reflect.TypeOf(data)
	if fieldTypes.Kind() == reflect.Ptr {
		fieldTypes = fieldTypes.Elem()
	}
	log.Info("Number of fields:", "fields", fieldTypes.NumField())
	for i := 0; i < fieldTypes.NumField(); i++ {
		f := fieldTypes.Field(i)
		log.Info("Field: %v -> %v", f.Name, f.Type.Name())
		// Load field methods
		ferr := ml.LoadAndDetectEmbeddedFields(f, modelFields, data)
		if ferr != nil {
			log.Warn("Ignoring unknown field type:", "FieldError", ferr.Error())
		}
	}
	return modelFields, nil
}

func (ml ModelLoader) LoadAndDetectEmbeddedFields(f reflect.StructField, fields map[string]loader.FieldDefinition, data interface{}) error {
	var err error
	// Load a base field
	if f.Type.Kind() != reflect.Struct || f.Type.Name() == "Time" {
		k, v, ferr := ml.GetFieldDetails(f, &data)
		if ferr != nil {
			if v == nil {
				log.Warn("Ignoring unknown field type:", "FieldError", ferr.Error(), "kind", f.Type.Kind(), "name", f.Name)
			} else {
				return ferr
			}
		}
		fields[k] = v
		return nil
	}

	enumWrapper := reflect.TypeOf((*loader.EnumWrapper)(nil)).Elem()
	if f.Type.Implements(enumWrapper) {
		log.Debug("Loading enum field: ", "type", f.Type.Name(), "kind", f.Type.Kind())
		k, v, ferr := ml.GetFieldDetails(f, &data)
		if ferr != nil {
			return ferr
		}
		fields[k] = v
		return nil
	}

	// Parse an embedded struct
	log.Debug("Parsing embedded struct", "type", f.Type.Name(), "name", f.Name)
	ff := f.Type
	for i := 0; i < ff.NumField(); i++ {
		f := ff.Field(i)
		log.Info("Field: %v -> %v", f.Name, f.Type.Name())
		if f.Type.Kind() == reflect.Struct {
			err = ml.LoadAndDetectEmbeddedFields(f, fields, data)
			if err != nil {
				log.Warn("Ignoring unknown field type:", "EmbeddedStruct", err)
			}
		}
		k, v, ferr := ml.GetFieldDetails(f, &data)
		if ferr != nil {
			if v == nil {
				log.Warn("Ignoring unknown field type:", "FieldError", ferr.Error())
			} else {
				return ferr
			}
		}
		fields[k] = v
	}

	return nil
}
func (ml ModelLoader) GetFieldOptions(name string, data interface{}) map[string]string {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Failed to get options from struct: ", "optError", err)
			panic(fmt.Sprintf("Field name not available: %v -> %v", name, err))
		}

	}()
	log.Info("Get options from field: ", "fieldName", name)
	structVal := reflect.ValueOf(data)
	log.Info("Get options from field: ", "structValue", structVal)
	field := structVal.FieldByName(name)
	if field.IsValid() {
		// Todo: Check the enum wrapper
		f := field.Interface()
		if wrapper, ok := f.(loader.EnumWrapper); ok {
			return wrapper.Values()
		}
	}
	return map[string]string{}
}

func (ml ModelLoader) ParseRelatedFields(f reflect.StructField, data TagData) (string, loader.FieldDefinition, error) {
	if len(data.Many2One) > 0 {
		ff := fields.One2Many{
			ReverseFK: data.Many2One,
			Stored:    data.Stored,
			ReadOnly:  data.Stored,
			Index:     data.Index,
		}
		return f.Name, ff, nil
	} else if len(data.Many2Many) > 0 {
		ff := fields.Many2Many{
			M2MOurField: data.Many2One,
			Stored:      data.Stored,
			ReadOnly:    data.Stored,
			Index:       data.Index,
		}
		return f.Name, ff, nil
	} else if len(data.Many2One) > 0 {
		ff := fields.Many2One{
			Related:  data.Many2One,
			Stored:   data.Stored,
			ReadOnly: data.Stored,
			Index:    data.Index,
		}
		return f.Name, ff, nil
	}
	return f.Name, nil, errors.New(fmt.Sprintf("unknown relation field type: %v", data.Type))
}
func (ml ModelLoader) GetFieldDetails(f reflect.StructField, modelData interface{}) (string, loader.FieldDefinition, error) {
	data := ml.GetFieldTags(f)
	// Parse relationship first
	fname, rval, err := ml.ParseRelatedFields(f, data)
	if err == nil && rval != nil {
		return fname, rval, nil
	} else {
		log.Debug("Non relationship field: ", "error", err)
	}
	// Parse other fields
	switch f.Type.Name() {
	case "float64", "float32":
		val := fields.Float{
			JSON:     data.JSON,
			String:   data.Message,
			Index:    data.Index,
			Required: data.Required,
			Help:     data.Help,
			Unique:   data.Unique,
			Stored:   data.Stored,
			ReadOnly: data.ReadOnly,
			NoCopy:   data.NoCopy,
			Depends:  data.Depends,
		}
		ml.GetFloatField(&val, data)
		return f.Name, val, nil
	case "string":
		if strings.Compare(data.Type, "html") == 0 {
			val := fields.HTML{
				JSON:     data.JSON,
				String:   data.Message,
				Index:    data.Index,
				Size:     data.Size,
				Required: data.Required,
				Help:     data.Help,
				Unique:   data.Unique,
				Stored:   data.Stored,
				ReadOnly: data.ReadOnly,
				NoCopy:   data.NoCopy,
			}
			return f.Name, val, nil
		} else if strings.Compare(data.Type, "text") == 0 {
			val := fields.Text{
				JSON:     data.JSON,
				String:   data.Message,
				Index:    data.Index,
				Size:     data.Size,
				Required: data.Required,
				Help:     data.Help,
				Unique:   data.Unique,
				Stored:   data.Stored,
				ReadOnly: data.ReadOnly,
				NoCopy:   data.NoCopy,
			}
			return f.Name, val, nil
		} else if strings.Compare(data.Type, "selection") == 0 {
			// Check for field wrapper of a selection
			if len(data.Options) == 0 {
				data.Options = ml.GetFieldOptions(f.Name, data)
			}
			// Create a selection
			val := fields.Selection{
				JSON:      data.JSON,
				String:    data.Message,
				Index:     data.Index,
				Selection: data.Options,
				Required:  data.Required,
				Help:      data.Help,
				Unique:    data.Unique,
				Stored:    data.Stored,
				ReadOnly:  data.ReadOnly,
				NoCopy:    data.NoCopy,
			}
			return f.Name, val, nil
		} else {
			val := fields.Char{
				JSON:     data.JSON,
				String:   data.Message,
				Index:    data.Index,
				Required: data.Required,
				Help:     data.Help,
				Unique:   data.Unique,
				Stored:   data.Stored,
				ReadOnly: data.ReadOnly,
				NoCopy:   data.NoCopy,
			}
			ml.GetStringField(&val, data)
			return f.Name, val, nil
		}
	case "int", "int64", "int32", "int16", "uint", "uint16", "uint32", "uint64":
		val := fields.Integer{
			JSON:     data.JSON,
			String:   data.Message,
			Index:    data.Index,
			Required: data.Required,
			Help:     data.Help,
			Unique:   data.Unique,
			Stored:   data.Stored,
			ReadOnly: data.ReadOnly,
			NoCopy:   data.NoCopy,
		}
		ml.GetIntegerField(&val, data)
		return f.Name, val, nil
	case "bool":
		val := fields.Boolean{
			JSON:     data.JSON,
			String:   data.Message,
			Index:    data.Index,
			Required: data.Required,
			Help:     data.Help,
			Unique:   data.Unique,
			Stored:   data.Stored,
			ReadOnly: data.ReadOnly,
			NoCopy:   data.NoCopy,
		}
		return f.Name, val, nil
	case "byte":
		val := fields.Binary{
			JSON:     data.JSON,
			String:   data.Message,
			Index:    data.Index,
			Required: data.Required,
			Help:     data.Help,
			Unique:   data.Unique,
			Stored:   data.Stored,
			ReadOnly: data.ReadOnly,
			NoCopy:   data.NoCopy,
		}
		return f.Name, val, nil
	case "Time":
		if strings.Compare(data.Type, "date") == 0 {
			val := fields.Date{
				JSON:     data.JSON,
				String:   data.Message,
				Index:    data.Index,
				Required: data.Required,
				Help:     data.Help,
				Unique:   data.Unique,
				Stored:   data.Stored,
				ReadOnly: data.ReadOnly,
				NoCopy:   data.NoCopy,
			}
			return f.Name, val, nil
		} else {
			val := fields.DateTime{
				JSON:     data.JSON,
				String:   data.Message,
				Index:    data.Index,
				Required: data.Required,
				Help:     data.Help,
				Unique:   data.Unique,
				Stored:   data.Stored,
				ReadOnly: data.ReadOnly,
				NoCopy:   data.NoCopy,
			}
			return f.Name, val, nil
		}

	case "EnumWrapper":
		data.Options = ml.GetFieldOptions(f.Name, data)
		// Create a selection
		val := fields.Selection{
			JSON:      data.JSON,
			String:    data.Message,
			Index:     data.Index,
			Selection: data.Options,
			Required:  data.Required,
			Help:      data.Help,
			Unique:    data.Unique,
			Stored:    data.Stored,
			ReadOnly:  data.ReadOnly,
			NoCopy:    data.NoCopy,
		}
		return f.Name, val, nil
	default:
		if strings.Compare(data.Type, "selection") == 0 {
			// Check for field wrapper of a selection
			if len(data.Options) == 0 {
				data.Options = ml.GetFieldOptions(f.Name, data)
			}
			// Create a selection
			val := fields.Selection{
				JSON:      data.JSON,
				String:    data.Message,
				Index:     data.Index,
				Selection: data.Options,
				Required:  data.Required,
				Help:      data.Help,
				Unique:    data.Unique,
				Stored:    data.Stored,
				ReadOnly:  data.ReadOnly,
				NoCopy:    data.NoCopy,
			}
			return f.Name, val, nil
		}
	}
	return f.Name, nil, errors.New(fmt.Sprintf("Unknown field type: %s", f.Type.Name()))
}

func (ml ModelLoader) GetFloatField(val *fields.Float, data TagData) {
	if data.Precision > 0 && data.Scale >= 0 {
		val.Digits = nbutils.Digits{
			Precision: data.Precision,
			Scale:     data.Scale,
		}
	} else {
		val.Digits = nbutils.Digits{
			Precision: 15,
			Scale:     2,
		}
	}
}

func (ml ModelLoader) GetStringField(val *fields.Char, data TagData) {
	val.Required = data.Required
}
func (ml ModelLoader) GetIntegerField(val *fields.Integer, data TagData) {
	val.Required = data.Required
}
func (ml ModelLoader) GetBooleanTag(f map[string]string, key string) bool {
	tagValue := false
	value, ok := f[key]
	if ok {
		if strings.Compare(value, "false") == 0 {
			tagValue = false
		} else {
			tagValue = true
		}
	} else {
		tagValue = false
	}
	return tagValue
}

func (ml ModelLoader) GetIntegerTag(f map[string]string, key string) int {
	tagValue := 0
	value, ok := f[key]
	if ok {
		tagValue, _ = strconv.Atoi(value)
	}
	return tagValue
}

// GetSelectionTag selection tags should have values: a,b:c,d:e,f; where : is the key-value pair separator
func (ml ModelLoader) GetSelectionTag(f map[string]string, key string) map[string]string {
	var tagValue map[string]string
	value := ml.GetStringTag(f, key, "")
	if len(value) > 0 {
		tagValue = make(map[string]string)
		dataParts := strings.Split(value, ":")
		for _, part := range dataParts {
			dd := strings.Split(part, ",")
			if len(dd) == 1 {
				tagValue[part] = strings.TrimSpace(part)
			} else {
				tagValue[dd[0]] = strings.TrimSpace(strings.Join(dd[1:], ", "))
			}
		}
	} else {
		tagValue = nil
	}
	return tagValue
}
func (ml ModelLoader) GetArrayTag(f map[string]string, key string) []string {
	var tagValue []string
	data := ml.GetStringTag(f, key, "")
	if len(data) > 0 {
		parts := strings.Split(data, ",")
		tagValue = append(tagValue, parts...)
	}
	return tagValue
}

func (ml ModelLoader) GetStringTag(f map[string]string, key string, deflt string) string {
	var tagValue string
	value, ok := f[key]
	if ok {
		tagValue = value
	} else {
		tagValue = deflt
	}
	return tagValue
}

func (ml ModelLoader) GetFieldTags(f reflect.StructField) TagData {
	vl := TagData{}
	// Parse Json setting
	value, ok := f.Tag.Lookup("json")
	if ok {
		value = strings.Split(value, ",")[0]
	} else {
		value = f.Name
	}
	vl.JSON = value
	// Parse hexya values
	tagStr, ok := f.Tag.Lookup("hexya")
	if ok {
		// Load extra tags from the configurations
		data := strings.Split(tagStr, ";")
		tagMap := make(map[string]string)
		for _, d := range data {
			st := strings.Split(d, "=")
			if len(st) == 1 {
				tagMap[st[0]] = ""
			} else if len(st) > 0 {
				tagMap[st[0]] = strings.Join(st[1:], "=")
			}
		}
		// Interprate each supported value
		vl.Required = ml.GetBooleanTag(tagMap, "required")
		vl.Scale = int8(ml.GetIntegerTag(tagMap, "scale"))
		vl.Precision = int8(ml.GetIntegerTag(tagMap, "precision"))
		vl.Size = ml.GetIntegerTag(tagMap, "size")
		vl.Message = ml.GetStringTag(tagMap, "display_name", f.Name)
		vl.Help = ml.GetStringTag(tagMap, "help", "")
		vl.Type = ml.GetStringTag(tagMap, "type", "")
		vl.Translate = ml.GetBooleanTag(tagMap, "translate")
		vl.Index = ml.GetBooleanTag(tagMap, "index")
		vl.Options = ml.GetSelectionTag(tagMap, "options")
		vl.Unique = ml.GetBooleanTag(tagMap, "unique")
		vl.Stored = ml.GetBooleanTag(tagMap, "stored")
		vl.ReadOnly = ml.GetBooleanTag(tagMap, "readOnly")
		vl.NoCopy = ml.GetBooleanTag(tagMap, "noCopy")
		vl.Depends = ml.GetArrayTag(tagMap, "depends")
		vl.goType = ml.GetStringTag(tagMap, "goType", "")
		// Relationship
		vl.One2Many = ml.GetStringTag(tagMap, "one2many", "")
		vl.Many2One = ml.GetStringTag(tagMap, "many2one", "")
		vl.Many2Many = ml.GetStringTag(tagMap, "many2many", "")
	}
	return vl
}

func NewTypedModel[M any](modelRef interface{}) ModelDefinition[M] {
	model, err := modelLoader.LoadBaseModel(modelRef)
	if err != nil {
		log.Error("Failed to load model: %v", err)
	}
	return ModelDefinition[M]{model}
}

func NewModelSet[M loader.Model](modelRef *loader.Model) ModelDefinition[M] {
	return NewModelSet[M](modelRef)
}
