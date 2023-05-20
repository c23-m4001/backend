package repository

import (
	"capstone/infrastructure"
	"capstone/model"
	"context"

	"github.com/Masterminds/squirrel"
)

type UserAccessTokenRepository interface {
	// create
	Insert(ctx context.Context, userAccessToken *model.UserAccessToken) error
	InsertMany(ctx context.Context, userAccessTokens []model.UserAccessToken) error

	// read
	FetchNLatestByUserId(ctx context.Context, userId string, n int) ([]model.UserAccessToken, error)
	Get(ctx context.Context, id string) (*model.UserAccessToken, error)
	IsExist(ctx context.Context, id string) (bool, error)

	// delete
	Truncate(ctx context.Context) error
}

type userAccessTokenRepository struct {
	db infrastructure.DBTX
}

func NewUserAccessTokenRepository(db infrastructure.DBTX) UserAccessTokenRepository {
	return &userAccessTokenRepository{
		db: db,
	}
}

func (r *userAccessTokenRepository) fetch(stmt squirrel.SelectBuilder) ([]model.UserAccessToken, error) {
	userAccessTokens := []model.UserAccessToken{}
	if err := fetch(r.db, &userAccessTokens, stmt); err != nil {
		return nil, err
	}

	return userAccessTokens, nil
}

func (r *userAccessTokenRepository) get(stmt squirrel.SelectBuilder) (*model.UserAccessToken, error) {
	userAccessToken := model.UserAccessToken{}
	if err := get(r.db, &userAccessToken, stmt); err != nil {
		return nil, err
	}

	return &userAccessToken, nil
}

func (r *userAccessTokenRepository) Insert(ctx context.Context, userAccessToken *model.UserAccessToken) error {
	return defaultInsert(r.db, ctx, userAccessToken, "*")
}

func (r *userAccessTokenRepository) InsertMany(ctx context.Context, userAccessTokens []model.UserAccessToken) error {
	arr := []model.BaseModel{}
	for i := range userAccessTokens {
		arr = append(arr, &userAccessTokens[i])
	}
	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *userAccessTokenRepository) FetchNLatestByUserId(ctx context.Context, userId string, n int) ([]model.UserAccessToken, error) {

	stmt := stmtBuilder.Select("*").
		From(model.UserAccessTokenTableName).
		Where(squirrel.Eq{"user_id": userId}).
		Limit(uint64(n)).
		OrderBy("created_at DESC")

	return r.fetch(stmt)
}

func (r *userAccessTokenRepository) Get(ctx context.Context, id string) (*model.UserAccessToken, error) {
	stmt := stmtBuilder.Select("*").
		From(model.UserAccessTokenTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(stmt)
}

func (r *userAccessTokenRepository) IsExist(ctx context.Context, id string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.UserAccessTokenTableName).
			Where(squirrel.Eq{"id": id}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, stmt)
}

func (r *userAccessTokenRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, model.UserAccessTokenTableName)
}
