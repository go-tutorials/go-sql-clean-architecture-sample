package usecase

import (
	"context"
	"database/sql"

	"go-sample/internal/user"
	"go-sample/internal/user/entity"
)

type UserUsecase struct {
	db         *sql.DB
	repository user.UserRepository
}

func NewUserService(db *sql.DB, repository user.UserRepository) *UserUsecase {
	return &UserUsecase{db: db, repository: repository}
}

func (s *UserUsecase) Load(ctx context.Context, id string) (*entity.User, error) {
	return s.repository.Load(ctx, id)
}

func (s *UserUsecase) Create(ctx context.Context, user *entity.User) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Create(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

func (s *UserUsecase) Update(ctx context.Context, user *entity.User) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Update(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

func (s *UserUsecase) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Patch(ctx, user)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}
	err = tx.Commit()
	return res, err
}

func (s *UserUsecase) Delete(ctx context.Context, id string) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "tx", tx)
	res, err := s.repository.Delete(ctx, id)
	if err != nil {
		er := tx.Rollback()
		if er != nil {
			return -1, er
		}
		return -1, err
	}
	err = tx.Commit()
	return res, err
}
