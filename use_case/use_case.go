package use_case

import (
	"capstone/model"
	"capstone/repository"
	"context"
)

type baseUseCase struct {
	categoryRepository    repository.CategoryRepository
	transactionRepository repository.TransactionRepository
	walletRepository      repository.WalletRepository
}

func NewBaseUseCase(
	categoryRepository repository.CategoryRepository,
	transactionRepository repository.TransactionRepository,
	walletRepository repository.WalletRepository,
) *baseUseCase {
	return &baseUseCase{
		categoryRepository:    categoryRepository,
		transactionRepository: transactionRepository,
		walletRepository:      walletRepository,
	}
}

func (u *baseUseCase) mustGetWallet(ctx context.Context, walletId string, isPath bool) model.Wallet {
	wallet, err := u.walletRepository.Get(ctx, walletId)
	panicIfRepositoryError(err, "Wallet data not found", isPath)
	return *wallet
}

func (u *baseUseCase) mustGetTransaction(ctx context.Context, transactionId string, isPath bool) model.Transaction {
	transaction, err := u.transactionRepository.Get(ctx, transactionId)
	panicIfRepositoryError(err, "Transaction data not found", isPath)
	return *transaction
}

func (u *baseUseCase) mustGetCategory(ctx context.Context, categoryId string, isPath bool) model.Category {
	category, err := u.categoryRepository.Get(ctx, categoryId)
	panicIfRepositoryError(err, "Category data not found", isPath)
	return *category
}

func (u *baseUseCase) mustAddWalletAmount(ctx context.Context, walletId string, amount float64) {
	panicIfErr(
		u.walletRepository.UpdateAddAmountById(ctx, walletId, amount),
	)
}
