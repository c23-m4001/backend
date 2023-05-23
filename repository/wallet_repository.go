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
	FetchByUserIds(ctx context.Context, userIds []string) ([]model.Wallet, error)
	Get(ctx context.Context, id string) (*model.Wallet, error)
	GetSumTotalAmountByUserId(ctx context.Context, userId string) (float64, error)
	IsExist(ctx context.Context, id string) (bool, error)
	IsExistByUserIds(ctx context.Context, userIds []string) (map[string]bool, error)

	// update
	Update(ctx context.Context, wallet *model.Wallet) error
	UpdateAmount(ctx context.Context, wallet *model.Wallet) error
	UpdateAddAmountById(ctx context.Context, walletId string, amount float64) error

	// delete
	Delete(ctx context.Context, wallet *model.Wallet) error
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
		From(model.WalletTableName)

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

func (r *walletRepository) FetchByUserIds(ctx context.Context, userIds []string) ([]model.Wallet, error) {
	stmt := stmtBuilder.Select("*").
		From(model.WalletTableName).
		Where(squirrel.Eq{"user_id": userIds})

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

func (r *walletRepository) IsExist(ctx context.Context, id string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("*").
			From(model.WalletTableName).
			Where(squirrel.Eq{"id": id}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, stmt)
}

func (r *walletRepository) IsExistByUserIds(ctx context.Context, userIds []string) (map[string]bool, error) {
	stmt := stmtBuilder.Select("DISTINCT user_id").
		From(model.WalletTableName).
		Where(squirrel.Eq{"user_id": userIds})

	fetchedUserIds := []string{}
	if err := fetch(r.db, &fetchedUserIds, stmt); err != nil {
		return nil, err
	}

	mapped := map[string]bool{}
	for _, fetchedUserId := range fetchedUserIds {
		mapped[fetchedUserId] = true
	}

	return mapped, nil
}

func (r *walletRepository) Update(ctx context.Context, wallet *model.Wallet) error {
	excludedColumns := []string{
		"total_amount",
		"user_id",
	}
	columns := extractColumnsFromBaseModel(wallet, excludedColumns)
	return defaultUpdate(r.db, ctx, wallet, columns, nil)
}

func (r *walletRepository) UpdateAmount(ctx context.Context, wallet *model.Wallet) error {
	return defaultUpdate(r.db, ctx, wallet, "total_amount", nil)
}

func (r *walletRepository) UpdateAddAmountById(ctx context.Context, walletId string, amount float64) error {
	params := map[string]interface{}{
		"amount": squirrel.Expr(fmt.Sprintf("amount + (%f)", amount)),
	}
	whereStmt := squirrel.Eq{
		"id": walletId,
	}

	return update(r.db, model.WalletTableName, params, whereStmt)
}

func (r *walletRepository) Delete(ctx context.Context, wallet *model.Wallet) error {
	return defaultDestroy(r.db, ctx, wallet, nil)
}

func (r *walletRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, model.WalletTableName)
}
