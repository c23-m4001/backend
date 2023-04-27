package use_case

import (
	"capstone/model"
	"capstone/repository"
	"context"
)

type baseUseCase struct {
	walletRepository repository.WalletRepository
}

func NewBaseUseCase(
	walletRepository repository.WalletRepository,
) *baseUseCase {
	return &baseUseCase{
		walletRepository: walletRepository,
	}
}

func (u *baseUseCase) mustGetWallet(ctx context.Context, walletId string, isPath bool) model.Wallet {
	wallet, err := u.walletRepository.Get(ctx, walletId)
	panicIfRepositoryError(err, "Wallet data not found", isPath)
	return *wallet
}
