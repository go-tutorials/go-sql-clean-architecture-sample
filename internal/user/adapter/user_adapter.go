package adapter

import (
	"context"
	"database/sql"
	"reflect"

	q "github.com/core-go/sql"

	"go-service/internal/user/entity"
)

type UserAdapter struct {
	keys          []string
	jsonColumnMap map[string]string
	DB            *sql.DB
}

func NewUserAdapter(db *sql.DB) *UserAdapter {
	userType := reflect.TypeOf(entity.User{})
	keys, _ := q.FindPrimaryKeys(userType)
	jsonColumnMap := q.MakeJsonColumnMap(userType)
	return &UserAdapter{keys: keys, jsonColumnMap: jsonColumnMap, DB: db}
}

func (r *UserAdapter) Load(ctx context.Context, id string) (*entity.User, error) {
	var users []entity.User
	query := `
		select id, username, email, phone, date_of_birth
		from users where id = $1 limit 1`
	err := q.Select(ctx, r.DB, &users, query, id)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		return &users[0], nil
	}
	return nil, nil
}

func (r *UserAdapter) Create(ctx context.Context, user *entity.User) (int64, error) {
	query, args := q.BuildToInsert("users", user, q.BuildParam)
	tx := q.GetTx(ctx)
	res, err := tx.ExecContext(ctx, query, args...)
	return RowsAffected(res, err)
}

func (r *UserAdapter) Update(ctx context.Context, user *entity.User) (int64, error) {
	tx := q.GetTx(ctx)
	query, args := q.BuildToUpdate("users", user, q.BuildParam)
	res, err := tx.ExecContext(ctx, query, args...)
	return RowsAffected(res, err)
}

func (r *UserAdapter) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	colMap := q.JSONToColumns(user, r.jsonColumnMap)
	query, args := q.BuildToPatch("users", colMap, r.keys, q.BuildParam)
	tx := q.GetTx(ctx)
	res, err := tx.ExecContext(ctx, query, args...)
	return RowsAffected(res, err)
}

func (r *UserAdapter) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from users where id = ?"
	tx := q.GetTx(ctx)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, id)
	return RowsAffected(res, err)
}

func RowsAffected(res sql.Result, err error) (int64, error) {
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}
