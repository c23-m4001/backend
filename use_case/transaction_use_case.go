package use_case

import (
	"capstone/delivery/dto_request"
	"capstone/delivery/dto_response"
	"capstone/model"
	"capstone/repository"
	"capstone/util"
	"context"
)

type TransactionUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.TransactionCreateRequest) model.Transaction

	// read
	Fetch(ctx context.Context, request dto_request.TransactionFetchRequest) ([]model.Transaction, int)
	Get(ctx context.Context, request dto_request.TransactionGetRequest) model.Transaction

	// update
	Update(ctx context.Context, request dto_request.TransactionUpdateRequest) model.Transaction

	// delete
	Delete(ctx context.Context, request dto_request.TransactionDeleteRequest)
}

type transactionUseCase struct {
	transactionRepository repository.TransactionRepository
	baseUseCase           *baseUseCase
}

func NewTransactionUseCase(
	transactionRepository repository.TransactionRepository,
	baseUseCase *baseUseCase,
) TransactionUseCase {
	return &transactionUseCase{
		transactionRepository: transactionRepository,
		baseUseCase:           baseUseCase,
	}
}

func (u *transactionUseCase) mustLoadTransactionData(transaction *model.Transaction) {

}

func (u *transactionUseCase) Create(ctx context.Context, request dto_request.TransactionCreateRequest) model.Transaction {
	currentUser := model.MustGetUserCtx(ctx)

	u.baseUseCase.mustGetCategory(ctx, request.CategoryId, panicIsNotPath)
	u.baseUseCase.mustGetWallet(ctx, request.WalletId, panicIsNotPath)

	transaction := model.Transaction{
		Id:         util.NewUuid(),
		CategoryId: request.CategoryId,
		WalletId:   request.WalletId,
		UserId:     currentUser.Id,
		Name:       request.Name,
		Amount:     request.Amount,
		Date:       request.Date,
	}

	// TODO: change wallet amount too
	panicIfErr(
		u.transactionRepository.Insert(ctx, &transaction),
	)

	u.mustLoadTransactionData(&transaction)

	return transaction
}

func (u *transactionUseCase) Fetch(ctx context.Context, request dto_request.TransactionFetchRequest) ([]model.Transaction, int) {
	queryOption := model.TransactionQueryOption{
		QueryOption: model.NewBasicQueryOption(request.Limit, request.Page, model.Sorts(request.Sorts)),
		CategoryId:  request.CategoryId,
		UserId:      util.StringP(model.MustGetUserCtx(ctx).Id),
		WalletId:    request.WalletId,
		Phrase:      request.Phrase,
	}

	transactions, err := u.transactionRepository.Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.transactionRepository.Count(ctx, queryOption)
	panicIfErr(err)

	return transactions, total
}

func (u *transactionUseCase) Get(ctx context.Context, request dto_request.TransactionGetRequest) model.Transaction {
	transaction := u.baseUseCase.mustGetTransaction(ctx, request.TransactionId, panicIsPath)

	if transaction.UserId != model.MustGetUserCtx(ctx).Id {
		panic(dto_response.NewForbiddenResponse("FORBIDDEN"))
	}

	u.mustLoadTransactionData(&transaction)

	return transaction
}

func (u *transactionUseCase) Update(ctx context.Context, request dto_request.TransactionUpdateRequest) model.Transaction {
	transaction := u.baseUseCase.mustGetTransaction(ctx, request.TransactionId, panicIsPath)

	if transaction.UserId != model.MustGetUserCtx(ctx).Id {
		panic(dto_response.NewForbiddenResponse("FORBIDDEN"))
	}

	transaction.Name = request.Name
	transaction.Amount = request.Amount
	transaction.Date = request.Date

	// TODO: change wallet amount too
	panicIfErr(
		u.transactionRepository.Update(ctx, &transaction),
	)

	u.mustLoadTransactionData(&transaction)

	return transaction
}

func (u *transactionUseCase) Delete(ctx context.Context, request dto_request.TransactionDeleteRequest) {
	transaction := u.baseUseCase.mustGetTransaction(ctx, request.TransactionId, panicIsPath)

	if transaction.UserId != model.MustGetUserCtx(ctx).Id {
		panic(dto_response.NewForbiddenResponse("FORBIDDEN"))
	}

	// TODO: change wallet amount too
	panicIfErr(
		u.transactionRepository.Delete(ctx, &transaction),
	)
}
