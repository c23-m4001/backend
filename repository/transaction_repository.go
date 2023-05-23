package repository

import (
	"capstone/data_type"
	"capstone/infrastructure"
	"capstone/model"
	"context"

	"github.com/Masterminds/squirrel"
)

type TransactionRepository interface {
	// create
	Insert(ctx context.Context, transaction *model.Transaction) error
	InsertMany(ctx context.Context, transactions []model.Transaction) error

	// read
	Count(ctx context.Context, options ...model.TransactionQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.TransactionQueryOption) ([]model.Transaction, error)
	Get(ctx context.Context, id string) (*model.Transaction, error)
	GetSumAmountFromPreviousDate(ctx context.Context, startingDate data_type.Date) (float64, error)
	IsExistByCategoryId(ctx context.Context, categoryId string) (bool, error)

	// update
	Update(ctx context.Context, transaction *model.Transaction) error

	// delete
	Delete(ctx context.Context, transaction *model.Transaction) error
	DeleteByWalletId(ctx context.Context, walletId string) error
	Truncate(ctx context.Context) error
}

type transactionRepository struct {
	db infrastructure.DBTX
}

func NewTransactionRepository(
	db infrastructure.DBTX,
) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) fetch(stmt squirrel.Sqlizer) ([]model.Transaction, error) {
	transactions := []model.Transaction{}
	if err := fetch(r.db, &transactions, stmt); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *transactionRepository) get(stmt squirrel.Sqlizer) (*model.Transaction, error) {
	transaction := model.Transaction{}
	if err := get(r.db, &transaction, stmt); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepository) prepareQuery(option model.TransactionQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(model.TransactionTableName)

	if option.CategoryId != nil {
		stmt = stmt.Where(squirrel.Eq{"category_id": option.CategoryId})
	}

	if option.WalletId != nil {
		stmt = stmt.Where(squirrel.Eq{"wallet_id": option.WalletId})
	}

	if option.UserId != nil {
		stmt = stmt.Where(squirrel.Eq{"user_id": option.UserId})
	}

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		stmt = stmt.Where(squirrel.ILike{"name": phrase})
	}

	if option.StartDate.DateP() != nil {
		stmt = stmt.Where(squirrel.GtOrEq{"date": option.StartDate})
	}

	if option.EndDate.DateP() != nil {
		stmt = stmt.Where(squirrel.Lt{"date": option.EndDate})
	}

	stmt = option.Prepare(stmt)

	return stmt
}

func (r *transactionRepository) Insert(ctx context.Context, transaction *model.Transaction) error {
	return defaultInsert(r.db, ctx, transaction, "*")
}

func (r *transactionRepository) InsertMany(ctx context.Context, transactions []model.Transaction) error {
	arr := []model.BaseModel{}
	for i := range transactions {
		arr = append(arr, &transactions[i])
	}
	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *transactionRepository) Count(ctx context.Context, options ...model.TransactionQueryOption) (int, error) {
	option := model.TransactionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.SetDefault()
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, stmt)
}

func (r *transactionRepository) Fetch(ctx context.Context, options ...model.TransactionQueryOption) ([]model.Transaction, error) {
	option := model.TransactionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.SetDefault()

	stmt := r.prepareQuery(option)

	return r.fetch(stmt)
}

func (r *transactionRepository) Get(ctx context.Context, id string) (*model.Transaction, error) {
	stmt := stmtBuilder.Select("*").
		From(model.TransactionTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(stmt)
}

func (r *transactionRepository) GetSumAmountFromPreviousDate(ctx context.Context, startingDate data_type.Date) (float64, error) {
	stmt := stmtBuilder.Select().
		Column(squirrel.ConcatExpr("SUM(", squirrel.Case().When("is_expense = true", "-1 * amount").Else("amount"), ")")).
		From(model.WalletTableName).
		Where(squirrel.Lt{"date": startingDate})

	sumTotal := 0.0
	if err := get(r.db, &sumTotal, stmt); err != nil {
		return 0, err
	}

	return sumTotal, nil
}

func (r *transactionRepository) IsExistByCategoryId(ctx context.Context, categoryId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("1").
			From(model.TransactionTableName).
			Where(squirrel.Eq{"category_id": categoryId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, stmt)
}

func (r *transactionRepository) Update(ctx context.Context, transaction *model.Transaction) error {
	excludedColumns := []string{
		"total_amount",
		"user_id",
	}
	columns := extractColumnsFromBaseModel(transaction, excludedColumns)
	return defaultUpdate(r.db, ctx, transaction, columns, nil)
}

func (r *transactionRepository) Delete(ctx context.Context, transaction *model.Transaction) error {
	return defaultDestroy(r.db, ctx, transaction, nil)
}

func (r *transactionRepository) DeleteByWalletId(ctx context.Context, walletId string) error {
	whereStmt := squirrel.Eq{
		"wallet_id": walletId,
	}

	return destroy(r.db, model.TransactionTableName, whereStmt)
}

func (r *transactionRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, model.TransactionTableName)
}
