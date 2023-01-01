package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type ProfileModel struct {
	HexyaBaseModel
	Age      int16
	Gender   string `hexya:"type=selection;options=male:Male,female:Female"`
	Money    float64
	Country  string
	UserName string `hexya:"related=User.Name"`
	Action   string `hexya:"goType=actions.ActionRef"`
}

type ProfileRepository[T ProfileModel, K int64] struct {
	ModelRepository[T, K]
}

func TestModelDeclaration(t *testing.T) {
	Convey("Creating DataBase...", t, func() {
		pp := ProfileRepository[ProfileModel, int64]{}
		Registry.add(pp)
	})
}
