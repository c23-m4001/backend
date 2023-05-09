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

	// delete
	Delete(ctx context.Context, request dto_request.WalletDeleteRequest)
}

type walletUseCase struct {
	compositeRepository repository.CompositeRepository
	walletRepository    repository.WalletRepository

	baseUseCase *baseUseCase
}

func NewWalletUseCase(
	compositeRepository repository.CompositeRepository,
	walletRepository repository.WalletRepository,
	baseUseCase *baseUseCase,
) WalletUseCase {
	return &walletUseCase{
		compositeRepository: compositeRepository,
		walletRepository:    walletRepository,

		baseUseCase: baseUseCase,
	}
}

func (u *walletUseCase) mustValidateWalletOwnedByCurrentUser(ctx context.Context, wallet model.Wallet) {
	currentUser := model.MustGetUserCtx(ctx)

	if wallet.UserId != currentUser.Id {
		panic(dto_response.NewForbiddenResponse("WALLET.FORBIDDEN_ACCESS"))
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
		UserId:      util.StringP(model.MustGetUserCtx(ctx).Id),
	}

	wallets, err := u.walletRepository.Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.walletRepository.Count(ctx, queryOption)
	panicIfErr(err)

	return wallets, total
}

func (u *walletUseCase) Get(ctx context.Context, request dto_request.WalletGetRequest) model.Wallet {
	wallet := u.baseUseCase.mustGetWallet(ctx, request.WalletId, panicIsPath)

	u.mustValidateWalletOwnedByCurrentUser(ctx, wallet)

	return wallet
}

func (u *walletUseCase) Update(ctx context.Context, request dto_request.WalletUpdateRequest) model.Wallet {
	wallet := u.baseUseCase.mustGetWallet(ctx, request.WalletId, panicIsPath)

	u.mustValidateWalletOwnedByCurrentUser(ctx, wallet)

	wallet.Name = request.Name
	wallet.LogoType = request.LogoType

	panicIfErr(
		u.walletRepository.Update(ctx, &wallet),
	)

	return wallet
}

func (u *walletUseCase) Delete(ctx context.Context, request dto_request.WalletDeleteRequest) {
	wallet := u.baseUseCase.mustGetWallet(ctx, request.WalletId, panicIsPath)

	u.mustValidateWalletOwnedByCurrentUser(ctx, wallet)

	panicIfErr(
		u.compositeRepository.DeleteWalletAndTransactions(ctx, &wallet),
	)
}
