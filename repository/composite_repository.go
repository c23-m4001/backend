package repository

import (
	"capstone/model"
	"context"

	"github.com/jmoiron/sqlx"
)

type CompositeRepository interface {
	InsertTransactionAndUpdateWalletAmount(ctx context.Context, transaction *model.Transaction, walelt *model.Wallet) error
	UpdateTransactionAndUpdateWalletAmount(ctx context.Context, transaction *model.Transaction, walelt *model.Wallet) error
	DeleteTransactionAndUpdateWalletAmount(ctx context.Context, transaction *model.Transaction, walelt *model.Wallet) error

	// wallet
	DeleteWalletAndTransactions(ctx context.Context, wallet *model.Wallet) error
}

type compositeRepository struct {
	db *sqlx.DB
}

func NewCompositeRepository(
	db *sqlx.DB,
) CompositeRepository {
	return &compositeRepository{
		db: db,
	}
}

func (r *compositeRepository) InsertTransactionAndUpdateWalletAmount(ctx context.Context, transaction *model.Transaction, wallet *model.Wallet) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return translateSqlError(err)
	}

	walletRepository := NewWalletRepository(tx)
	transactionRepository := NewTransactionRepository(tx)

	if err := transactionRepository.Insert(ctx, transaction); err != nil {
		tx.Rollback()
		return err
	}

	if err := walletRepository.UpdateAmount(ctx, wallet); err != nil {
		tx.Rollback()
		return err
	}

	return translateSqlError(tx.Commit())
}

func (r *compositeRepository) UpdateTransactionAndUpdateWalletAmount(ctx context.Context, transaction *model.Transaction, wallet *model.Wallet) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return translateSqlError(err)
	}

	walletRepository := NewWalletRepository(tx)
	transactionRepository := NewTransactionRepository(tx)

	if err := transactionRepository.Update(ctx, transaction); err != nil {
		tx.Rollback()
		return err
	}

	if err := walletRepository.UpdateAmount(ctx, wallet); err != nil {
		tx.Rollback()
		return err
	}

	return translateSqlError(tx.Commit())
}

func (r *compositeRepository) DeleteTransactionAndUpdateWalletAmount(ctx context.Context, transaction *model.Transaction, wallet *model.Wallet) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return translateSqlError(err)
	}

	walletRepository := NewWalletRepository(tx)
	transactionRepository := NewTransactionRepository(tx)

	if err := transactionRepository.Delete(ctx, transaction); err != nil {
		tx.Rollback()
		return err
	}

	if err := walletRepository.UpdateAmount(ctx, wallet); err != nil {
		tx.Rollback()
		return err
	}

	return translateSqlError(tx.Commit())
}

func (r *compositeRepository) DeleteWalletAndTransactions(ctx context.Context, wallet *model.Wallet) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return translateSqlError(err)
	}

	walletRepository := NewWalletRepository(tx)
	transactionRepository := NewTransactionRepository(tx)

	if err := transactionRepository.DeleteByWalletId(ctx, wallet.Id); err != nil {
		tx.Rollback()
		return err
	}

	if err := walletRepository.Delete(ctx, wallet); err != nil {
		tx.Rollback()
		return err
	}

	return translateSqlError(tx.Commit())
}
