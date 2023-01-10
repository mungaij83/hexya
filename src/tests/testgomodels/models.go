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
	"github.com/hexya-erp/hexya/src/models"
	"time"
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
	models.HexyaBaseModel
	Name          string `hexya:"String=Name;Help=The user's username;unique;NoCopy"`
	DecoratedName string `hexya:"display_name=Decorated Name"`
	Email         string `hexya:"help=The user's email address;size=100;index=true"`
	Password      string `hexya:"noCopy"`
	Status        int    `json:"status_json" hexya:"goType=int16;default=12"`
	IsStaff       bool   `hexya:"display_name=isStaffString;help=IsStaffHelp"`
	IsActive      bool   `hexya:"display_name=Active User"`
	ProfileId     int64
	Profile       ProfileModel `hexya:"one2one=id;onDelete=SetNull" gorm:"foreignKey:Id;references=ProfileId"`
	Age           int          `hexya:"depends=Profile;Profile.Age;Stored=true;goType=int16"`
	Posts         []PostModel  `hexya:"one2many=Id;ReverseFK=User;copy=true;" gorm:"foreignKey:UserId"`
	PMoney        float64      `hexya:"related:Profile.Money"`
	Resume        ResumeModel  `hexya:"embed=true" gorm:"embed"`
	LastPostId    int64
	LastPost      *PostModel `hexya:"many2one=Id" gorm:"foreignKey;references:LastPostId"`
	Email2        string     `hexya:"help=Email user"`
	IsPremium     bool       `hexya:"display_name=isPremiumString;help=isPremiumHelp"`
	Nums          int
	Size          float64
	Education     string `hexya:"display_name=Educational Background"`
}

func (UserModel) TableName() string {
	return "user"
}
func (UserModel) ModelName() string {
	return "user"
}

type ProfileModel struct {
	models.HexyaBaseModel
	Age        int16
	Gender     string `hexya:"type=selection;options=male:Male,female:Female"`
	Money      float64
	UserId     int64
	User       *UserModel `hexya:"ReverseFK=Profile" gorm:"foreignKey;references=UserId"`
	BestPostId int64
	BestPost   PostModel `hexya:"many2one=Id" gorm:"foreignKey;references=BestPostId"`
	Country    string
	UserName   string `hexya:"related=User.Name"`
	Action     string `hexya:"goType=actions.ActionRef"`
}

type PostModel struct {
	models.HexyaBaseModel
	UserId           int64
	User             *UserModel     `hexya:"many2one=Id" gorm:"foreignKey:UserId"`
	Title            string         `hexya:"required"`
	Content          string         `hexya:"type=html"`
	Abstract         string         `hexya:"type=text"`
	Tags             []TagModel     `hexya:"many2many=Id" gorm:"foreignKey:PostId"`
	Comments         []CommentModel `hexya:"many2many=Id;ReverseFK=Post" gorm:"foreignKey:PostId"`
	LastRead         time.Time      `hexya:"type=datetime"`
	FirstCommentText string
	FirstTagName     string `hexya:"related=Tags.Name"`
	WriterMoney      float64
}
type CommentModel struct {
	models.HexyaBaseModel
	PostId      int64
	Post        PostModel `hexya:"many2one=ID" gorm:"foreignKey:PostId"`
	WriterMoney float64   `hexya:"related=PostWriter.PMoney"`
	Text        string
}

type TagModel struct {
	models.HexyaBaseModel
	Name         string
	BestPostId   int64
	BestPost     PostModel `hexya:"many2one=id" gorm:"foreignKey:BestPostId"`
	PostId       int64
	Posts        []PostModel `hexya:"many2many=id;ReverseFK=PostId" gorm:"many2many:PostId"`
	ParentPostId int64
	Parent       *TagModel `hexya:"many2one=id" gorm:"foreignKey:ParentPostId"`
	Description  string
	Rate         float64
}

type ResumeModel struct {
	Education  string
	Experience string `hexya:"translate"`
	Leisure    string `hexya:"type=text"`
	Other      string `hexya:"computed=true"`
}

type AddressModelMixin struct {
	Street string
	Zip    string
	City   string
}

type ActiveModelMixin struct {
	Active bool
}

type UserViewModel struct {
	Name string
	City string
}
