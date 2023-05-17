package use_case

import (
	"capstone/delivery/dto_request"
	"capstone/delivery/dto_response"
	"capstone/model"
	"capstone/repository"
	"capstone/util"
	"context"
)

type WalletUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.WalletCreateRequest) model.Wallet

	// read
	Fetch(ctx context.Context, request dto_request.WalletFetchRequest) ([]model.Wallet, int)
	Get(ctx context.Context, request dto_request.WalletGetRequest) model.Wallet

	// update
	Update(ctx context.Context, request dto_request.WalletUpdateRequest) model.Wallet
}

type walletUseCase struct {
	walletRepository repository.WalletRepository

	baseUseCase *baseUseCase
}

func NewWalletUseCase(
	walletRepository repository.WalletRepository,
	baseUseCase *baseUseCase,
) WalletUseCase {
	return &walletUseCase{
		walletRepository: walletRepository,

		baseUseCase: baseUseCase,
	}
}

func (u *walletUseCase) Create(ctx context.Context, request dto_request.WalletCreateRequest) model.Wallet {
	currentUser := model.MustGetUserCtx(ctx)

	wallet := model.Wallet{
		Id:          util.NewUuid(),
		UserId:      currentUser.Id,
		Name:        request.Name,
		TotalAmount: 0,
		LogoType:    request.LogoType,
	}

	panicIfErr(
		u.walletRepository.Insert(ctx, &wallet),
	)

	return wallet
}

func (u *walletUseCase) Fetch(ctx context.Context, request dto_request.WalletFetchRequest) ([]model.Wallet, int) {
	queryOption := model.WalletQueryOption{
		QueryOption: model.NewBasicQueryOption(request.Limit, request.Page, model.Sorts(request.Sorts)),
		Phrase:      request.Phrase,
	}

	wallets, err := u.walletRepository.Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.walletRepository.Count(ctx, queryOption)
	panicIfErr(err)

	return wallets, total
}

func (u *walletUseCase) Get(ctx context.Context, request dto_request.WalletGetRequest) model.Wallet {
	wallet := u.baseUseCase.mustGetWallet(ctx, request.WalletId, panicIsPath)

	return wallet
}

func (u *walletUseCase) Update(ctx context.Context, request dto_request.WalletUpdateRequest) model.Wallet {
	currentUser := model.MustGetUserCtx(ctx)
	wallet := u.baseUseCase.mustGetWallet(ctx, request.WalletId, panicIsNotPath)

	if wallet.UserId != currentUser.Id {
		panic(dto_response.NewForbiddenResponse("Forbidden to access this wallet"))
	}

	wallet.Name = request.Name
	wallet.LogoType = request.LogoType

	panicIfErr(
		u.walletRepository.Update(ctx, &wallet),
	)

	return wallet
}
