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

package tests

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/loader"
	"github.com/hexya-erp/hexya/src/models/security"
	testmodule "github.com/hexya-erp/hexya/src/tests/testgomodels"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConditions(t *testing.T) {
	Convey("Testing SQL building for queries", t, func() {
		So(loader.SimulateInNewEnvironment(security.SuperUserID, func(env loader.Environment) {
			rp := models.GetRepository[testmodule.UserModel]()
			rs, err := rp.Search(rp.Condition().ID().Equals(1).Underlying())
			So(err, ShouldBeNil)
			Convey("Simple query", func() {
				So(rs, ShouldNotBeNil)
			})
			Convey("Simple query with args inflation", func() {
				rs2, err := rp.Search(rp.Condition().FieldName("username").Equals("test"))
				So(err, ShouldBeNil)
				So(rs2, ShouldNotBeNil)
			})
			Convey("Check WHERE clause with additionnal filter", func() {
				rs, err = rp.Search(rp.Condition().FilteredOn(rp.GetFieldName("age"), rp.Condition().FieldName("age").GreaterOrEqual(12)))
				So(err, ShouldNotBeNil)
				So(rs, ShouldNotBeNil)
			})
			//Convey("Check full query with all conditions", func() {
			//	rs = rp.Search(rp.Condition().FilteredOn(rp..Age().GreaterOrEqual(12)).Or().Name().ILike("John"))
			//	c2 := q.User().Name().Like("jane").Or().ProfileFilteredOn(q.Profile().Money().Lower(1234.56))
			//	rs = rs.Search(c2)
			//	rs.Load()
			//	So(func() { rs.Load() }, ShouldNotPanic)
			//})
		}), ShouldBeNil)
	})
}
