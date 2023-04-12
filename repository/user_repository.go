package repository

import (
	"capstone/infrastructure"
	"capstone/model"
	"context"

	"github.com/Masterminds/squirrel"
)

type UserRepository interface {
	// create
	Insert(ctx context.Context, user *model.User) error
	InsertMany(ctx context.Context, users []model.User) error

	// read
	Get(ctx context.Context, id string) (*model.User, error)

	// delete
	Truncate(ctx context.Context) error
}

type userRepository struct {
	db infrastructure.DBTX
}

func NewUserRepository(
	db infrastructure.DBTX,
) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) get(stmt squirrel.Sqlizer) (*model.User, error) {
	user := model.User{}
	if err := get(r.db, &user, stmt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Insert(ctx context.Context, user *model.User) error {
	return defaultInsert(r.db, ctx, user, "*")
}

func (r *userRepository) InsertMany(ctx context.Context, users []model.User) error {
	arr := []model.BaseModel{}
	for i := range users {
		arr = append(arr, &users[i])
	}
	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *userRepository) Get(ctx context.Context, id string) (*model.User, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(stmt)
}

func (r *userRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, model.UserTableName)
}
