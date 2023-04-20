package repository

import (
	"capstone/infrastructure"
	"capstone/model"
	"context"
)

type UserAccessTokenRepository interface {
	// create
	Insert(ctx context.Context, userAccessToken *model.UserAccessToken) error
	InsertMany(ctx context.Context, userAccessTokens []model.UserAccessToken) error

	// delete
	Truncate(ctx context.Context) error
}

type userAccessTokenRepository struct {
	db infrastructure.DBTX
}

func NewUserAccessTokenRepository(db infrastructure.DBTX) UserAccessTokenRepository {
	return &userAccessTokenRepository{}
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

func (r *userAccessTokenRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, model.UserAccessTokenTableName)
}
