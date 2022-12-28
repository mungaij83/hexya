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

package testmodule

import (
	"fmt"
	"github.com/hexya-erp/hexya/src/models/loader"
	"log"
	"time"

	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

var (
	// IsStaffHelp exported
	IsStaffHelp   = "This is a var help message"
	isPremiumHelp = "This the IsPremium Help message"
)

const (
	isStaffString        = "Is Staff"
	isPremiumDescription = "This is a const description string"
)

const isPremiumString = isPremiumDescription

type UserModel struct {
	Name          string `hexya:"String=Name;Help=The user's username;unique;NoCopy"`
	DecoratedName string
	Email         string `hexya:"help=The user's email address;size=100;index=true"`
	Password      string `hexya:"noCopy"`
	Status        int    `json:"status_json" hexya:"goType=int16;default=12"`
	IsStaff       bool   `hexya:"display_name=isStaffString;help=IsStaffHelp"`
	IsActive      bool
	Profile       ProfileModel `hexya:"one2one=id;onDelete=SetNull"`
	Age           int          `hexya:"depends=Profile;Profile.Age;Stored=true;goType=int16"`
	Posts         []PostModel  `hexya:"one2many=Id;ReverseFK=User;copy=true"`
	PMoney        float64      `hexya:"related:Profile.Money"`
	Resume        ResumeModel  `hexya:"embed=true"`
	LastPost      *PostModel   `hexya:"many2one=Id"`
	Email2        string
	IsPremium     bool `hexya:"display_name=isPremiumString;help=isPremiumHelp"`
	Nums          int
	Size          float64
	Education     string `hexya:"display_name=Educational Background"`
}

// DecorateEmail decorates the email of the given user
func user_DecorateEmail(_ m.UserSet, email string) string {
	return fmt.Sprintf("<%s>", email)
}

func user_OnChangeName(rs m.UserSet) m.UserData {
	return h.User().NewData().SetDecoratedName(rs.PrefixedUser("User")[0])
}

func user_ComputeDecoratedName(rs m.UserSet) m.UserData {
	return h.User().NewData().SetDecoratedName(rs.PrefixedUser("User")[0])
}

func user_ComputeAge(rs m.UserSet) m.UserData {
	return h.User().NewData().SetAge(rs.Profile().Age())
}

func user_PrefixedUser(rs m.UserSet, prefix string) []string {
	var res []string
	for _, u := range rs.Records() {
		res = append(res, fmt.Sprintf("%s: %s", prefix, u.Name()))
	}
	return res
}

func user_ext_DecorateEmail(rs m.UserSet, email string) string {
	res := rs.Super().DecorateEmail(email)
	return fmt.Sprintf("[%s]", res)
}

func user_RecursiveMethod(rs m.UserSet, depth int, result string) string {
	if depth == 0 {
		return result
	}
	return rs.RecursiveMethod(depth-1, fmt.Sprintf("%s, recursion %d", result, depth))
}

func user_ext_RecursiveMethod(rs m.UserSet, depth int, result string) string {
	result = "> " + result + " <"
	sup := rs.Super().RecursiveMethod(depth, result)
	return sup
}

func user_SubSetSuper(rs m.UserSet) string {
	var res string
	for _, rec := range rs.Records() {
		res += rec.Name()
	}
	return res
}

func user_ext_SubSetSuper(rs m.UserSet) string {
	userJane := h.User().Search(rs.Env(), q.User().Email().Equals("jane.smith@example.com"))
	userJohn := h.User().Search(rs.Env(), q.User().Email().Equals("jsmith2@example.com"))
	users := h.User().NewSet(rs.Env())
	users = users.Union(userJane)
	users = users.Union(userJohn)
	return users.Super().SubSetSuper()
}

func user_InverseSetAge(rs m.UserSet, age int16) {
	rs.Profile().SetAge(age)
}

func user_ext_PrefixedUser(rs m.UserSet, prefix string) []string {
	res := rs.Super().PrefixedUser(prefix)
	for i, u := range rs.Records() {
		res[i] = fmt.Sprintf("%s %s", res[i], rs.DecorateEmail(u.Email()))
	}
	return res
}

func user_UpdateCity(rs m.UserSet, value string) {
	rs.Profile().SetCity(value)
}

func user_Aggregates(rs m.UserSet, fieldNames ...models.FieldName) []m.UserGroupAggregateRow {
	return rs.Super().Aggregates(fieldNames...)
}

type ProfileModel struct {
	Age      int16
	Gender   string `hexya:"type=selection;options=male:Male,female:Female"`
	Money    float64
	User     *UserModel `hexya:"ReverseFK=Profile"`
	BestPost PostModel  `hexya:"many2one=Id"`
	Country  string
	UserName string `hexya:"related=User.Name"`
	Action   string `hexya:"goType=actions.ActionRef"`
}

func profile_PrintAddress(rs m.ProfileSet) string {
	res := rs.Super().PrintAddress()
	return fmt.Sprintf("%s, %s", res, rs.Country())
}

func profile_ext_PrintAddress(rs m.ProfileSet) string {
	res := rs.Super().PrintAddress()
	return fmt.Sprintf("[%s]", res)
}

func post_Create(rs m.PostSet, data m.PostData) m.PostSet {
	res := rs.Super().Create(data)
	return res
}

func post_Search(rs m.PostSet, cond q.PostCondition) m.PostSet {
	res := rs.Super().Search(cond)
	return res
}

type PostModel struct {
	User             *UserModel     `hexya:"many2one=Id"`
	Title            string         `hexya:"required"`
	Content          string         `hexya:"type=html"`
	Abstract         string         `hexya:"type=text"`
	Tags             []TagModel     `hexya:"many2many=Id"`
	Comments         []CommentModel `hexya:"many2many=Id;ReverseFK=Post"`
	LastRead         time.Time      `hexya:"type=datetime"`
	FirstCommentText string
	FirstTagName     string `hexya:"related=Tags.Name"`
	WriterMoney      float64
}
type CommentModel struct {
	Post        PostModel `hexya:"many2one=ID"`
	WriterMoney float64   `hexya:"related=PostWriter.PMoney"`
	Text        string
}

type TagModel struct {
	Name        string
	BestPost    PostModel   `hexya:"many2one=id"`
	Posts       []PostModel `hexya:"many2many=id"`
	Parent      *TagModel   `hexya:"many2one=id"`
	Description string
	Rate        float64
}

func tag_CheckNameDescription(rs m.TagSet) {
	if rs.Name() == rs.Description() {
		log.Panic("Tag name and description must be different")
	}
}

func tag_CheckRate(rs m.TagSet) {
	if rs.Rate() < 0 || rs.Rate() > 10 {
		log.Panic("Tag rate must be between 0 and 10")
	}
}

type ResumeModel struct {
	Education  string
	Experience string `hexya:"translate"`
	Leisure    string `hexya:"type=text"`
	Other      string `hexya:"computed=true"`
}

func resume_Create(rs m.ResumeSet, data m.ResumeData) m.ResumeSet {
	return rs.Super().Create(data)
}

func resume_ComputeOther(_ m.ResumeSet) m.ResumeData {
	return h.Resume().NewData().SetOther("Other information")
}

type AddressModelMixin struct {
	Street string
	Zip    string
	City   string
}

func addressMixIn_SayHello(_ m.AddressMixInSet) string {
	return "Hello !"
}

func addressMixIn_PrintAddress(rs m.AddressMixInSet) string {
	return fmt.Sprintf("%s, %s %s", rs.Street(), rs.Zip(), rs.City())
}

func addressMixIn_ext_PrintAddress(rs m.AddressMixInSet) string {
	res := rs.Super().PrintAddress()
	return fmt.Sprintf("<%s>", res)
}

type ActiveModelMixin struct {
	Active bool
}

func activeMixIn_IsActivated(rs m.ActiveMixInSet) bool {
	return rs.Active()
}

type UserViewModel struct {
	Name string
	City string
}

func init() {
	user := loader.NewModelDefinition(UserModel{})

	user.Fields().Experience().SetString("Professional Experience")

	user.NewMethod("OnChangeName", user_OnChangeName)
	user.NewMethod("ComputeDecoratedName", user_ComputeDecoratedName)
	user.NewMethod("ComputeAge", user_ComputeAge)
	user.NewMethod("PrefixedUser", user_PrefixedUser)
	user.NewMethod("DecorateEmail", user_DecorateEmail)
	user.NewMethod("RecursiveMethod", user_RecursiveMethod)
	user.NewMethod("SubSetSuper", user_SubSetSuper)
	user.NewMethod("InverseSetAge", user_InverseSetAge)
	user.NewMethod("UpdateCity", user_UpdateCity)
	user.Methods().DecorateEmail().Extend(user_ext_DecorateEmail)
	user.Methods().RecursiveMethod().Extend(user_ext_RecursiveMethod)
	user.Methods().SubSetSuper().Extend(user_ext_SubSetSuper)
	user.Methods().PrefixedUser().Extend(user_ext_PrefixedUser)
	user.Methods().Aggregates().Extend(user_Aggregates)

	profile := loader.NewModelDefinition(ProfileModel{})
	profile.InheritModel(h.AddressMixIn())

	profile.AddFields(fields_Profile)
	profile.Fields().Zip().SetString("Zip Code")

	profile.Methods().PrintAddress().Extend(profile_PrintAddress)
	profile.Methods().PrintAddress().Extend(profile_ext_PrintAddress)

	post := loader.NewModelDefinition(PostModel{})

	post.Methods().Create().Extend(post_Create)
	post.Methods().Search().Extend(post_Search)

	comment := loader.NewModelDefinition(CommentModel{})
	print(comment.IsManual())
	tag := loader.NewModelDefinition(TagModel{})
	tag.SetDefaultOrder("Name DESC", "ID ASC")

	tag.NewMethod("CheckNameDescription", tag_CheckNameDescription).AllowGroup(security.GroupEveryone)
	tag.NewMethod("CheckRate", tag_CheckRate)

	resume := loader.NewModelDefinition(ResumeModel{})

	resume.Methods().Create().Extend(resume_Create)
	resume.NewMethod("ComputeOther", resume_ComputeOther)

	addressMI2 := loader.NewModelDefinition(AddressModelMixin{})
	addressMI2.NewMethod("SayHello", addressMixIn_SayHello)
	addressMI2.NewMethod("PrintAddress", addressMixIn_PrintAddress)
	addressMI2.Methods().PrintAddress().Extend(addressMixIn_ext_PrintAddress)

	activeMixin1 := loader.NewModelDefinition(ActiveModelMixin{})

	// Chained declaration
	activeMI2 := activeMixin1
	activeMI2.NewMethod("IsActivated", activeMixIn_IsActivated)

	manualModel := loader.NewModelDefinition(UserViewModel{})
	print(manualModel.IsManual())
}
