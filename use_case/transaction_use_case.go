package use_case

import (
	"capstone/delivery/dto_request"
	"capstone/delivery/dto_response"
	"capstone/loader"
	"capstone/model"
	"capstone/repository"
	"capstone/util"
	"context"

	"golang.org/x/sync/errgroup"
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
	compositeRepository   repository.CompositeRepository
	transactionRepository repository.TransactionRepository
	baseUseCase           *baseUseCase
}

func NewTransactionUseCase(
	compositeRepository repository.CompositeRepository,
	transactionRepository repository.TransactionRepository,
	baseUseCase *baseUseCase,
) TransactionUseCase {
	return &transactionUseCase{
		compositeRepository:   compositeRepository,
		transactionRepository: transactionRepository,
		baseUseCase:           baseUseCase,
	}
}

func (u *transactionUseCase) mustLoadTransactionData(transaction *model.Transaction) {
	categoryLoader := loader.NewCategoryloader(u.baseUseCase.categoryRepository)

	panicIfErr(
		await(func(group *errgroup.Group) {
			group.Go(categoryLoader.TransactionFn(transaction))
		}),
	)
}

func (u *transactionUseCase) mustLoadTransactionsData(transactions []model.Transaction) {
	categoryLoader := loader.NewCategoryloader(u.baseUseCase.categoryRepository)

	panicIfErr(
		await(func(group *errgroup.Group) {
			for i := range transactions {
				group.Go(categoryLoader.TransactionFn(&transactions[i]))
			}
		}),
	)
}

func (u *transactionUseCase) Create(ctx context.Context, request dto_request.TransactionCreateRequest) model.Transaction {
	currentUser := model.MustGetUserCtx(ctx)

	category := u.baseUseCase.mustGetCategory(ctx, request.CategoryId, panicIsNotPath)
	wallet := u.baseUseCase.mustGetWallet(ctx, request.WalletId, panicIsNotPath)

	transaction := model.Transaction{
		Id:         util.NewUuid(),
		CategoryId: request.CategoryId,
		WalletId:   request.WalletId,
		UserId:     currentUser.Id,
		Name:       request.Name,
		Amount:     request.Amount,
		Date:       request.Date,
	}

	if category.IsExpense {
		wallet.TotalAmount -= request.Amount
	} else {
		wallet.TotalAmount += request.Amount
	}

	panicIfErr(
		u.compositeRepository.InsertTransactionAndUpdateWalletAmount(ctx, &transaction, &wallet),
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

	u.mustLoadTransactionsData(transactions)

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

	category := u.baseUseCase.mustGetCategory(ctx, transaction.CategoryId, true)
	wallet := u.baseUseCase.mustGetWallet(ctx, transaction.WalletId, true)

	previousAmount := transaction.Amount

	transaction.Name = request.Name
	transaction.Amount = request.Amount
	transaction.Date = request.Date

	if category.IsExpense {
		wallet.TotalAmount -= (request.Amount - previousAmount)
	} else {
		wallet.TotalAmount += (request.Amount - previousAmount)
	}

	panicIfErr(
		u.compositeRepository.UpdateTransactionAndUpdateWalletAmount(ctx, &transaction, &wallet),
	)

	u.mustLoadTransactionData(&transaction)

	return transaction
}

func (u *transactionUseCase) Delete(ctx context.Context, request dto_request.TransactionDeleteRequest) {
	transaction := u.baseUseCase.mustGetTransaction(ctx, request.TransactionId, panicIsPath)

	if transaction.UserId != model.MustGetUserCtx(ctx).Id {
		panic(dto_response.NewForbiddenResponse("FORBIDDEN"))
	}

	category := u.baseUseCase.mustGetCategory(ctx, transaction.CategoryId, true)
	wallet := u.baseUseCase.mustGetWallet(ctx, transaction.WalletId, true)

	if category.IsExpense {
		wallet.TotalAmount += transaction.Amount
	} else {
		wallet.TotalAmount -= transaction.Amount
	}

	panicIfErr(
		u.compositeRepository.DeleteTransactionAndUpdateWalletAmount(ctx, &transaction, &wallet),
	)
}
