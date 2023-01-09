// Copyright 2016 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testllmodule

import (
	"fmt"
	"github.com/hexya-erp/hexya/src/models/conditions"
	"github.com/hexya-erp/hexya/src/models/loader"
	"log"

	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/hexya/src/models/types"
)

func declareModels() {
	user := loader.NewModel("User")
	profile := loader.NewModel("Profile")
	post := loader.NewModel("Post")
	tag := loader.NewModel("Tag")
	cv := loader.NewModel("Resume")
	addressMI := loader.NewMixinModel("AddressMixIn")
	activeMI := loader.NewMixinModel("ActiveMixIn")
	viewModel := loader.NewManualModel("UserView")

	user.NewMethod("PrefixedUser",
		func(rc *loader.RecordCollection, prefix string) []string {
			var res []string
			for _, u := range rc.Records() {
				res = append(res, fmt.Sprintf("%s: %s", prefix, u.Get(loader.Name)))
			}
			return res
		})

	user.Methods().MustGet("PrefixedUser").Extend(
		func(rc *loader.RecordCollection, prefix string) []string {
			res := rc.Super().Call("PrefixedUser", prefix).([]string)
			for i, u := range rc.Records() {
				email := u.Get(rc.Model().FieldName("Email")).(string)
				res[i] = fmt.Sprintf("%s %s", res[i], rc.Call("DecorateEmail", email))
			}
			return res
		})

	user.NewMethod("DecorateEmail",
		func(rc *loader.RecordCollection, email string) string {
			if rc.Env().Context().HasKey("use_square_brackets") {
				return fmt.Sprintf("[%s]", email)
			}
			return fmt.Sprintf("<%s>", email)
		})

	user.Methods().MustGet("DecorateEmail").Extend(
		func(rc *loader.RecordCollection, email string) string {
			if rc.Env().Context().HasKey("use_double_square") {
				rc = rc.
					Call("WithContext", "use_square_brackets", true).(*loader.RecordCollection).
					WithContext("fake_key", true)
			}
			res := rc.Super().Call("DecorateEmail", email).(string)
			return fmt.Sprintf("[%s]", res)
		})

	user.NewMethod("OnChangeName",
		func(rc *loader.RecordCollection) *conditions.ModelData {
			res := make(loader.FieldMap)
			res["DecoratedName"] = rc.Call("PrefixedUser", "User").([]string)[0]
			return conditions.NewModelDataFromRS(rc, res)
		})

	user.NewMethod("ComputeDecoratedName",
		func(rc *loader.RecordCollection) *conditions.ModelData {
			res := make(loader.FieldMap)
			res["DecoratedName"] = rc.Call("PrefixedUser", "User").([]string)[0]
			return conditions.NewModelDataFromRS(rc, res)
		})

	user.NewMethod("ComputeAge",
		func(rc *loader.RecordCollection) *conditions.ModelData {
			res := make(loader.FieldMap)
			res["Age"] = rc.Get(rc.Model().FieldName("Profile")).(*loader.RecordCollection).Get(rc.Model().FieldName("Age")).(int16)
			return conditions.NewModelDataFromRS(rc, res)
		})

	user.NewMethod("InverseSetAge",
		func(rc *loader.RecordCollection, age int16) {
			rc.Get(rc.Model().FieldName("Profile")).(*loader.RecordCollection).Set(rc.Model().FieldName("Age"), age)
		})

	user.NewMethod("UpdateCity",
		func(rc *loader.RecordCollection, value string) {
			rc.Get(rc.Model().FieldName("Profile")).(*loader.RecordCollection).Set(rc.Model().FieldName("City"), value)
		})

	activeMI.NewMethod("IsActivated",
		func(rc *loader.RecordCollection) bool {
			return rc.Get(rc.Model().FieldName("Active")).(bool)
		})

	addressMI.NewMethod("SayHello",
		func(rc *loader.RecordCollection) string {
			return "Hello !"
		})

	addressMI.NewMethod("PrintAddress",
		func(rc *loader.RecordCollection) string {
			return fmt.Sprintf("%s, %s %s", rc.Get(rc.Model().FieldName("Street")), rc.Get(rc.Model().FieldName("Zip")), rc.Get(rc.Model().FieldName("City")))
		})

	profile.NewMethod("PrintAddress",
		func(rc *loader.RecordCollection) string {
			res := rc.Super().Call("PrintAddress").(string)
			return fmt.Sprintf("%s, %s", res, rc.Get(rc.Model().FieldName("Country")))
		})

	addressMI.Methods().MustGet("PrintAddress").Extend(
		func(rc *loader.RecordCollection) string {
			res := rc.Super().Call("PrintAddress").(string)
			return fmt.Sprintf("<%s>", res)
		})

	profile.Methods().MustGet("PrintAddress").Extend(
		func(rc *loader.RecordCollection) string {
			res := rc.Super().Call("PrintAddress").(string)
			return fmt.Sprintf("[%s]", res)
		})

	post.Methods().MustGet("Create").Extend(
		func(rc *loader.RecordCollection, data conditions.RecordData) *loader.RecordCollection {
			res := rc.Super().Call("Create", data).(conditions.RecordSet).Collection()
			return res
		})

	post.Methods().MustGet("WithContext").Extend(
		func(rc *loader.RecordCollection, key string, value interface{}) *loader.RecordCollection {
			return rc.Super().Call("WithContext", key, value).(*loader.RecordCollection)
		})

	tag.NewMethod("CheckRate",
		func(rc *loader.RecordCollection) {
			if rc.Get(rc.Model().FieldName("Rate")).(float32) < 0 || rc.Get(rc.Model().FieldName("Rate")).(float32) > 10 {
				log.Panic("Tag rate must be between 0 and 10")
			}
		})

	tag.NewMethod("CheckNameDescription",
		func(rc *loader.RecordCollection) {
			if rc.Get(rc.Model().FieldName("Name")).(string) == rc.Get(rc.Model().FieldName("Description")).(string) {
				log.Panic("Tag name and description must be different")
			}
		})

	// Because we run without pool, we need to declare our CRUD mixin methods
	for _, methName := range []string{"Load", "Create", "Write", "Unlink"} {
		tag.AddEmptyMethod(methName)
	}
	tag.Methods().AllowAllToGroup(security.GroupEveryone)

	user.AddFields(map[string]loader.FieldDefinition{
		"Name": fields.Char{String: "Name", Help: "The user's username", Unique: true,
			NoCopy: true, OnChange: user.Methods().MustGet("OnChangeName")},
		"DecoratedName": fields.Char{Compute: user.Methods().MustGet("ComputeDecoratedName")},
		"Email":         fields.Char{Help: "The user's email address", Size: 100, Index: true},
		"Password":      fields.Char{NoCopy: true},
		"Status": fields.Integer{JSON: "status_json", GoType: new(int16),
			Default: loader.DefaultValue(int16(12))},
		"IsStaff":  fields.Boolean{},
		"IsActive": fields.Boolean{},
		"Profile":  fields.Many2One{RelationModel: models.Registry.MustGet("Profile")},
		"Age": fields.Integer{Compute: user.Methods().MustGet("ComputeAge"),
			Inverse: user.Methods().MustGet("InverseSetAge"),
			Depends: []string{"Profile", "Profile.Age"}, Stored: true, GoType: new(int16)},
		"Posts":     fields.One2Many{RelationModel: models.Registry.MustGet("Post"), ReverseFK: "User"},
		"PMoney":    fields.Float{Related: "Profile.Money"},
		"LastPost":  fields.Many2One{RelationModel: models.Registry.MustGet("Post"), Embed: true},
		"Email2":    fields.Char{},
		"IsPremium": fields.Boolean{},
		"Nums":      fields.Integer{GoType: new(int)},
		"Size":      fields.Float{},
		"Education": fields.Text{String: "Educational Background"},
	})

	profile.AddFields(map[string]loader.FieldDefinition{
		"Age":      fields.Integer{GoType: new(int16)},
		"Gender":   fields.Selection{Selection: types.Selection{"male": "Male", "female": "Female"}},
		"Money":    fields.Float{},
		"User":     fields.Many2One{RelationModel: models.Registry.MustGet("User")},
		"BestPost": fields.One2One{RelationModel: models.Registry.MustGet("Post")},
		"City":     fields.Char{},
		"Country":  fields.Char{},
	})

	post.AddFields(map[string]loader.FieldDefinition{
		"User":            fields.Many2One{RelationModel: models.Registry.MustGet("User")},
		"Title":           fields.Char{},
		"Content":         fields.HTML{},
		"Tags":            fields.Many2Many{RelationModel: models.Registry.MustGet("Tag")},
		"BestPostProfile": fields.Rev2One{RelationModel: models.Registry.MustGet("Profile"), ReverseFK: "BestPost"},
		"Abstract":        fields.Text{},
		"Attachment":      fields.Binary{},
		"LastRead":        fields.Date{},
	})

	post.Methods().MustGet("Create").Extend(
		func(rc *loader.RecordCollection, data conditions.RecordData) *loader.RecordCollection {
			res := rc.Super().Call("Create", data).(*loader.RecordCollection)
			return res
		})

	tag.AddFields(map[string]loader.FieldDefinition{
		"Name":        fields.Char{Constraint: tag.Methods().MustGet("CheckNameDescription")},
		"BestPost":    fields.Many2One{RelationModel: models.Registry.MustGet("Post")},
		"Posts":       fields.Many2Many{RelationModel: models.Registry.MustGet("Post")},
		"Parent":      fields.Many2One{RelationModel: models.Registry.MustGet("Tag")},
		"Description": fields.Char{Constraint: tag.Methods().MustGet("CheckNameDescription")},
		"Rate":        fields.Float{Constraint: tag.Methods().MustGet("CheckRate"), GoType: new(float32)},
	})

	cv.AddFields(map[string]loader.FieldDefinition{
		"Education":  fields.Text{},
		"Experience": fields.Text{Translate: true},
		"Leisure":    fields.Text{},
	})

	addressMI.AddFields(map[string]loader.FieldDefinition{
		"Street": fields.Char{GoType: new(string)},
		"Zip":    fields.Char{},
		"City":   fields.Char{},
	})
	profile.InheritModel(addressMI)

	activeMI.AddFields(map[string]loader.FieldDefinition{
		"Active": fields.Boolean{},
	})

	models.Registry.MustGet("CommonMixin").InheritModel(activeMI)

	viewModel.AddFields(map[string]loader.FieldDefinition{
		"Name": fields.Char{},
		"City": fields.Char{},
	})
}
