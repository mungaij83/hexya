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
		//
		tt := ProfileModel{
			Age: 10, Gender: "Test",
			Money:    7.0,
			Country:  "KE",
			UserName: "test",
			Action:   "Test",
		}
		log.Debug("Data:", "dd", tt, "isModel", tt.IsModel(), "isAbstract", !tt.IsTransient(), "Transient", tt.IsTransient())
		// Initialize profile
		// Limitations above:https://go101.org/generics/888-the-status-quo-of-go-custom-generics.html
		// Make it difficult to use the BaseModel, methods in base are not detected as implemented by the other models
		// Also some methods are implemented with interface type as argument or return typ for the reason above, maybe a work around is available.
		pp := ProfileRepository[ProfileModel, int64]{}
		// Save before should result in an error
		err := pp.Save(&tt)
		log.Debug("Should fail on clean Database:", "error", err)
		//So(err, ShouldNotBeNil)
		// Add to registry
		err = Registry.add(pp)
		So(err, ShouldBeNil)
		// Save after adding model to registry
		err = pp.Save(&tt)
		log.Debug("Failed to add profile", "error", err)
		So(err, ShouldBeNil)

		err = pp.registerExtension(&DefaultMixinExtension[ProfileModel]{})
		So(err, ShouldBeNil)
	})
}
