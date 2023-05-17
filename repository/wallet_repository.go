package repository

import (
	"capstone/infrastructure"
	"capstone/model"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type WalletRepository interface {
	// create
	Insert(ctx context.Context, wallet *model.Wallet) error
	InsertMany(ctx context.Context, wallets []model.Wallet) error

	// read
	Count(ctx context.Context, options ...model.WalletQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.WalletQueryOption) ([]model.Wallet, error)
	FetchByUserId(ctx context.Context, userId string) ([]model.Wallet, error)
	Get(ctx context.Context, id string) (*model.Wallet, error)
	GetSumTotalAmountByUserId(ctx context.Context, userId string) (float64, error)

	// update
	Update(ctx context.Context, wallet *model.Wallet) error

	// delete
	Truncate(ctx context.Context) error
}

type walletRepository struct {
	db infrastructure.DBTX
}

func NewWalletRepository(
	db infrastructure.DBTX,
) WalletRepository {
	return &walletRepository{
		db: db,
	}
}

func (r *walletRepository) fetch(stmt squirrel.Sqlizer) ([]model.Wallet, error) {
	wallets := []model.Wallet{}
	if err := fetch(r.db, &wallets, stmt); err != nil {
		return nil, err
	}

	return wallets, nil
}

func (r *walletRepository) get(stmt squirrel.Sqlizer) (*model.Wallet, error) {
	wallet := model.Wallet{}
	if err := get(r.db, &wallet, stmt); err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepository) prepareQuery(option model.WalletQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s w", model.WalletTableName))

	if option.UserId != nil {
		stmt = stmt.Where(squirrel.Eq{"user_id": option.UserId})
	}

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.ILike{"name": phrase})
	}

	stmt = option.Prepare(stmt)

	return stmt
}

func (r *walletRepository) Insert(ctx context.Context, wallet *model.Wallet) error {
	return defaultInsert(r.db, ctx, wallet, "*")
}

func (r *walletRepository) InsertMany(ctx context.Context, wallets []model.Wallet) error {
	arr := []model.BaseModel{}
	for i := range wallets {
		arr = append(arr, &wallets[i])
	}
	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *walletRepository) Count(ctx context.Context, options ...model.WalletQueryOption) (int, error) {
	option := model.WalletQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.SetDefault()
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, stmt)
}

func (r *walletRepository) Fetch(ctx context.Context, options ...model.WalletQueryOption) ([]model.Wallet, error) {
	option := model.WalletQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.SetDefault()

	stmt := r.prepareQuery(option)

	return r.fetch(stmt)
}

func (r *walletRepository) FetchByUserId(ctx context.Context, userId string) ([]model.Wallet, error) {
	stmt := stmtBuilder.Select("*").
		From(model.WalletTableName).
		Where(squirrel.Eq{"user_id": userId})

	return r.fetch(stmt)
}

func (r *walletRepository) Get(ctx context.Context, id string) (*model.Wallet, error) {
	stmt := stmtBuilder.Select("*").
		From(model.WalletTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(stmt)
}

func (r *walletRepository) GetSumTotalAmountByUserId(ctx context.Context, userId string) (float64, error) {
	stmt := stmtBuilder.Select("SUM(total_amount)").
		From(model.WalletTableName).
		Where(squirrel.Eq{"user_id": userId})

	sumTotal := 0.0
	if err := get(r.db, &sumTotal, stmt); err != nil {
		return 0, err
	}

	return sumTotal, nil
}

func (r *walletRepository) Update(ctx context.Context, wallet *model.Wallet) error {
	excludedColumns := []string{
		"total_amount",
		"user_id",
	}
	columns := extractColumnsFromBaseModel(wallet, excludedColumns)
	return defaultUpdate(r.db, ctx, wallet, columns, nil)
}

func (r *walletRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, model.WalletTableName)
}
