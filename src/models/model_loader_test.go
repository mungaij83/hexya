package models

import (
	"github.com/hexya-erp/hexya/src/models/types"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var mLoader ModelLoader

type TestModel struct {
	Name         string    `json:"name" hexya:"display_name=Name;index;required;translate"`
	CreateOn     time.Time `json:"create_on" hexya:"display_name=Create On"`
	ProductCount int64     `json:"product_count" hexya:"display_name=# Products;help=The number of products under this category (Does not consider the children categories)"`
}

type TestModel2 struct {
	HexyaBaseModel
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
	Gender       GenderEnum `json:"gender" hexya:"type=selection;display_name=Gender;translate"`
	CreateOn     time.Time  `json:"create_on" hexya:"display_name=Create On"`
	Model2Id     int64
	Model2       TestModel2 `hexya:"one2many=Model2Id" gorm:"foreignKey=Model2Id"`
	ModelId      int64
	Model        *TestModel `hexya:"one2many=ModelId" gorm:"foreignKey=ModelId"`
	ProductCount int64      `json:"product_count" hexya:"display_name=# Products;help=The number of products under this category (Does not consider the children categories)"`
}

func (TestModel2) ParentTableName() string {
	return "test_model"
}

func LoadMain(t *testing.M) {
	mLoader = ModelLoader{}
	t.Run()
}

func TestModelLoader_LoadBaseModel(t *testing.T) {
	Convey("Load model", t, func() {
		mdl, err := mLoader.LoadBaseModel(TestModel{})
		So(err, ShouldBeNil)
		So(mdl, ShouldNotBeNil)
		t.Logf("Data: %+v", mdl)
		mdl, err = mLoader.LoadBaseModel(TestModel2{})
		So(err, ShouldBeNil)
		So(mdl, ShouldNotBeNil)
		So(mdl.Fields().MustGet("CreateDate"), ShouldNotBeNil)
		So(mdl.Fields().MustGet("CreateUID"), ShouldNotBeNil)
		So(mdl.IsMixin(), ShouldBeFalse)
		t.Logf("Data 2: %+v", mdl)
		mdl, err = mLoader.LoadBaseModel(TestModel3{})
		So(err, ShouldBeNil)
		f, ok := mdl.Fields().Get("Model2")
		So(ok, ShouldBeTrue)
		So(len(f.ReverseFK) > 0, ShouldBeTrue)
		f, ok = mdl.Fields().Get("Model")
		So(ok, ShouldBeTrue)
		So(len(f.ReverseFK) > 0, ShouldBeTrue)

		So(mdl.FieldName("Model2"), ShouldNotBeNil)
		So(mdl, ShouldNotBeNil)
		_, ok = mdl.Fields().Get("Gender")
		So(ok, ShouldBeTrue)
	})
}
func TestModelLoader_LoadModel(t *testing.T) {
	Convey("Load model details", t, func() {
		modelFields, err := mLoader.LoadModel(TestModel{})
		So(err, ShouldBeNil)
		So(modelFields, ShouldNotBeNil)
		t.Logf("Data: %+v", modelFields)
	})

}
