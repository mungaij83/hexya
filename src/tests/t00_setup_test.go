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
	testmodule "github.com/hexya-erp/hexya/src/tests/testgomodels"
	"log"
	"testing"

	_ "github.com/hexya-erp/hexya/src/tests/testmodule"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	RunTests(m, "tests", func() {
		err := models.RegisterModel(testmodule.ProfileRepository[testmodule.ProfileModel, int64]{})
		if err != nil {
			log.Printf("Error registering user: %v", err)
		}
		err = models.RegisterModel(testmodule.UserRepository[testmodule.UserModel, int64]{})
		if err != nil {
			log.Printf("Error registering user: %v", err)
		}
		err = models.RegisterModel(testmodule.TagRepository[testmodule.TagModel, int64]{})
		if err != nil {
			log.Printf("Error registering tags:%v", err)
		}
		err = models.RegisterModel(testmodule.PostRepository[testmodule.PostModel, int64]{})
		if err != nil {
			log.Printf("Error registering posts: %v", err)
		}
		err = models.RegisterModel(testmodule.CommentRepository[testmodule.CommentModel, int64]{})
		if err != nil {
			log.Printf("Error registering comments:%v", err)
		}

	})
}
