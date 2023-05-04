package use_case

import (
	"capstone/model"
	"capstone/repository"
	"context"
)

type baseUseCase struct {
	categoryRepository repository.CategoryRepository
	walletRepository   repository.WalletRepository
}

func NewBaseUseCase(
	categoryRepository repository.CategoryRepository,
	walletRepository repository.WalletRepository,
) *baseUseCase {
	return &baseUseCase{
		categoryRepository: categoryRepository,
		walletRepository:   walletRepository,
	}
}

func (u *baseUseCase) mustGetWallet(ctx context.Context, walletId string, isPath bool) model.Wallet {
	wallet, err := u.walletRepository.Get(ctx, walletId)
	panicIfRepositoryError(err, "Wallet data not found", isPath)
	return *wallet
}

func (u *baseUseCase) mustGetCategory(ctx context.Context, categoryId string, isPath bool) model.Category {
	category, err := u.categoryRepository.Get(ctx, categoryId)
	panicIfRepositoryError(err, "Category data not found", isPath)
	return *category
}
