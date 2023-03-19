package user

import (
	"context"
	"go-sample/internal/user/entity"
)

// type Reader interface {
// 	Load(ctx context.Context, id string) (*entity.User, error)

// }

// type Writer interface {
// 	Create(ctx context.Context, user *entity.User) (int64, error)
// 	Update(ctx context.Context, user *entity.User) (int64, error)
// 	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
// 	Delete(ctx context.Context, id string) (int64, error)
// }

// type Repository interface {
// 	Reader
// 	Writer
// }

type UserRepository interface {
	Load(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (int64, error)
	Update(ctx context.Context, user *entity.User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}