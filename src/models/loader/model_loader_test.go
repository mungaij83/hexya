package loader

import (
	"github.com/hexya-erp/hexya/src/models/types"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var loader ModelLoader

type TestModel struct {
	Name         string    `json:"name" hexya:"display_name=Name;index;required;translate"`
	CreateOn     time.Time `json:"create_on" hexya:"display_name=Create On"`
	ProductCount int64     `json:"product_count" hexya:"display_name=# Products;help=The number of products under this category (Does not consider the children categories)"`
}

type TestModel2 struct {
	Name         string    `json:"name" hexya:"display_name=Name;index;required;translate"`
	CreateOn     time.Time `json:"create_on" hexya:"display_name=Create On"`
	ProductCount int64     `json:"product_count" hexya:"display_name=# Products;help=The number of products under this category (Does not consider the children categories)"`
}

type GenderEnum struct {
}

func (GenderEnum) Values() types.Selection {
	return types.Selection{
		"M": "Male",
		"F": "Female",
		"O": "Others",
	}
}

type TestModel3 struct {
	Name         string     `json:"name" hexya:"display_name=Name;index;required;translate"`
	Gender       GenderEnum `json:"name" hexya:"type=selection;display_name=Gender;translate"`
	CreateOn     time.Time  `json:"create_on" hexya:"display_name=Create On"`
	ProductCount int64      `json:"product_count" hexya:"display_name=# Products;help=The number of products under this category (Does not consider the children categories)"`
}

func (TestModel2) ParentTableName() string {
	return "test_model"
}

func TestMain(t *testing.M) {
	loader = ModelLoader{}
	t.Run()
}
func TestModelLoader_LoadBaseModel(t *testing.T) {
	Convey("Load model", t, func() {
		mdl, err := loader.LoadBaseModel(TestModel{})
		So(err, ShouldBeNil)
		So(mdl, ShouldNotBeNil)
		t.Logf("Data: %+v", mdl)
		mdl, err = loader.LoadBaseModel(TestModel2{})
		So(err, ShouldBeNil)
		So(mdl, ShouldNotBeNil)
		So(mdl.IsMixin(), ShouldBeTrue)
		t.Logf("Data 2: %+v", mdl)
		mdl, err = loader.LoadBaseModel(TestModel3{})
		So(err, ShouldBeNil)
		So(mdl, ShouldNotBeNil)
		_, ok := mdl.Fields().Get("Gender")
		So(ok, ShouldBeTrue)
	})
}
func TestModelLoader_LoadModel(t *testing.T) {
	Convey("Load model details", t, func() {
		modelFields, err := loader.LoadModel(TestModel{})
		So(err, ShouldBeNil)
		So(modelFields, ShouldNotBeNil)
		t.Logf("Data: %+v", modelFields)
	})

}
