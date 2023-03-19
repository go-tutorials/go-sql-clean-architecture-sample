package user

import (
	"context"

	"go-service/internal/user/entity"
)

type UserService interface {
	Load(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (int64, error)
	Update(ctx context.Context, user *entity.User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
