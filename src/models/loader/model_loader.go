package loader

import (
	"errors"
	"fmt"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/loader/generic"

	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/tools/logging"
	"github.com/hexya-erp/hexya/src/tools/nbutils"
	"reflect"
	"strconv"
	"strings"
)

var (
	log         logging.Logger
	modelLoader ModelLoader
)

func init() {
	log = logging.GetLogger("model_loader")
	modelLoader = ModelLoader{}
}

type TagData struct {
	Value     interface{}
	JSON      string
	Required  bool
	Type      string
	Translate bool
	Index     bool
	Size      int
	NoCopy    bool
	Unique    bool
	Precision int8
	Scale     int8
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
	// Override name of the table from the type interfaces
	var tmpName string
	switch v := data.(type) {
	case models.ManualModelType:
		tmpName = v.ExistingTableName()
		break
	case models.TableModel:
		tmpName = v.TableName()
		break
	case models.MixinModelType:
		tmpName = v.ParentTableName()
		break
	case models.TransientTableModel:
		tmpName = v.BaseMixin()
		break
	}
	// Override table name
	if len(tmpName) > 2 {
		tableName = tmpName
	}
	return tableName
}
func (ml ModelLoader) detectTableType(data interface{}) (string, models.Option) {
	// Determine the type or name of the model
	var optionSetting models.Option
	tableName := ml.detectTableName(data)
	// Override name of the table from the type interfaces
	switch v := data.(type) {
	case models.TableModel:
		optionSetting = 0
		log.Info("Loading model: ", v.TableName())
		break
	case models.MixinModelType:
		optionSetting = models.MixinModel
		break
	case models.TransientTableModel:
		optionSetting = models.TransientModel
		break
	default:
		optionSetting = models.ManualModel
	}
	return tableName, optionSetting
}

// Model  type is determined with priority as seen in the case statement
// Implementing multiple types will result in other types being less prioritized as per the list below
func (ml ModelLoader) LoadBaseModel(data interface{}) (*models.Model, error) {
	fieldDefinitions, err := ml.LoadModel(data)
	if err != nil {
		return nil, err
	}
	tableName, option := ml.detectTableType(data)
	var mdl *models.Model
	switch option {
	case models.MixinModel:
		mdl = models.NewMixinModel(tableName)
		break
	case models.ManualModel:
		mdl = models.NewManualModel(tableName)
		break
	case models.TransientModel:
		mdl = models.NewTransientModel(tableName)
		break
	default:
		mdl = models.NewModel(tableName)
		break
	}
	// Add Fields and sorting order
	mdl.AddFields(fieldDefinitions)
	switch v := data.(type) {
	case models.OrderedTableModel:
		orders := v.OrderFields()
		if len(orders) > 0 {
			mdl.SetDefaultOrder(orders...)
		}
		break
	}
	return mdl, nil
}

func (ml ModelLoader) LoadModel(data interface{}) (map[string]models.FieldDefinition, error) {
	modelFields := make(map[string]models.FieldDefinition)
	fieldTypes := reflect.TypeOf(data)
	log.Info("Number of fields: %v", fieldTypes.NumField())
	for i := 0; i < fieldTypes.NumField(); i++ {
		f := fieldTypes.Field(i)
		log.Info("Field: %v -> %v", f.Name, f.Type.Name())
		k, v, ferr := ml.GetFieldDetails(f, &data)
		if ferr != nil {
			if v == nil {
				log.Warn("Ignoring unknown field type: %v", ferr.Error())
			} else {
				return nil, ferr
			}
		}
		modelFields[k] = v
	}
	return modelFields, nil
}

func (ml ModelLoader) GetFieldOptions(name string, data interface{}) map[string]string {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Failed to get options from struct: ", err)
			panic(fmt.Sprintf("Field name not available: %v -> %v", name, err))
		}

	}()
	log.Info("Get options from field: ", name)
	structVal := reflect.ValueOf(data)
	log.Info("Get options from field: ", structVal)
	field := structVal.FieldByName(name)
	if field.IsValid() {
		// Todo: Check the enum wrapper
		f := field.Interface()
		if wrapper, ok := f.(models.EnumWrapper); ok {
			return wrapper.Values()
		}
	}
	return map[string]string{}
}
func (ml ModelLoader) GetFieldDetails(f reflect.StructField, modelData interface{}) (string, models.FieldDefinition, error) {
	data := ml.GetFieldTags(f)
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
	case "int", "int64", "int32":
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

// Selection tag should have values: a,b:c,d:e,f; where : is the key-value pair separator
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
		vl.Stored = ml.GetBooleanTag(tagMap, "readOnly")
		vl.NoCopy = ml.GetBooleanTag(tagMap, "noCopy")
	}
	return vl
}

func NewTypedModel(modelRef interface{}) *models.Model {
	model, err := modelLoader.LoadBaseModel(modelRef)
	if err != nil {
		log.Error("Failed to load model: %v", err)
	}
	model.InheritModel(models.Registry.MustGet("ModelMixin"))
	return model
}

func NewModelSet(modelRef interface{}) generic.ModelDefinition {

}