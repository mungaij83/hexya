package testmodule

import "github.com/hexya-erp/hexya/src/models"

type UserRepository[T UserModel, K int64] struct {
	models.ModelRepository[UserModel, int64]
}

type CommentRepository[T CommentModel, K int64] struct {
	models.ModelRepository[T, K]
}

type PostRepository[T PostModel, K int64] struct {
	models.ModelRepository[T, K]
}
type TagRepository[T TagModel, K int64] struct {
	models.ModelRepository[T, K]
}
