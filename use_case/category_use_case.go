package use_case

import (
	"capstone/delivery/dto_request"
	"capstone/delivery/dto_response"
	"capstone/model"
	"capstone/repository"
	"capstone/util"
	"context"
)

type CategoryUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.CategoryCreateRequest) model.Category

	// read
	Fetch(ctx context.Context, request dto_request.CategoryFetchRequest) ([]model.Category, int)
	Get(ctx context.Context, request dto_request.CategoryGetRequest) model.Category

	// update
	Update(ctx context.Context, request dto_request.CategoryUpdateRequest) model.Category

	// delete
	Delete(ctx context.Context, request dto_request.CategoryDeleteRequest)

	// option
	OptionForTransactionForm(ctx context.Context, request dto_request.CategoryOptionForTransactionFormRequest) ([]model.Category, int)
}

type categoryUseCase struct {
	categoryRepository repository.CategoryRepository

	baseUseCase *baseUseCase
}

func NewCategoryUseCase(
	categoryRepository repository.CategoryRepository,
	baseUseCase *baseUseCase,
) CategoryUseCase {
	return &categoryUseCase{
		categoryRepository: categoryRepository,

		baseUseCase: baseUseCase,
	}
}

func (u *categoryUseCase) mustValidateAllowDelete(ctx context.Context, categoryId string) {
	isExist, err := u.categoryRepository.IsExist(ctx, categoryId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestResponse("CATEGORY.IN_USED"))
	}
}

func (u *categoryUseCase) Create(ctx context.Context, request dto_request.CategoryCreateRequest) model.Category {
	currentUser := model.MustGetUserCtx(ctx)

	category := model.Category{
		Id:        util.NewUuid(),
		UserId:    &currentUser.Id,
		Name:      request.Name,
		IsGlobal:  false,
		IsExpense: request.IsExpense,
	}

	panicIfErr(
		u.categoryRepository.Insert(ctx, &category),
	)

	return category
}

func (u *categoryUseCase) Fetch(ctx context.Context, request dto_request.CategoryFetchRequest) ([]model.Category, int) {
	currentUser := model.MustGetUserCtx(ctx)

	queryOption := model.CategoryQueryOption{
		QueryOption:   model.NewBasicQueryOption(request.Limit, request.Page, model.Sorts(request.Sorts)),
		IncludeGlobal: util.BoolP(true),
		IsExpense:     request.IsExpense,
		UserId:        &currentUser.Id,
		Phrase:        request.Phrase,
	}

	categories, err := u.categoryRepository.Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.categoryRepository.Count(ctx, queryOption)
	panicIfErr(err)

	return categories, total
}

func (u *categoryUseCase) Get(ctx context.Context, request dto_request.CategoryGetRequest) model.Category {
	category := u.baseUseCase.mustGetCategory(ctx, request.CategoryId, panicIsPath)

	if !category.IsGlobal && *category.UserId != model.MustGetUserCtx(ctx).Id {
		panic(dto_response.NewForbiddenResponse("FORBIDDEN"))
	}

	return category
}

func (u *categoryUseCase) Update(ctx context.Context, request dto_request.CategoryUpdateRequest) model.Category {
	category := u.baseUseCase.mustGetCategory(ctx, request.CategoryId, panicIsPath)

	if category.IsGlobal {
		panic(dto_response.NewBadRequestResponse("CATEGORY.IS_DEFAULT_CATEGORY"))
	}

	if !category.IsGlobal && *category.UserId != model.MustGetUserCtx(ctx).Id {
		panic(dto_response.NewForbiddenResponse("FORBIDDEN"))
	}

	category.Name = request.Name
	category.IsExpense = request.IsExpense

	panicIfErr(
		u.categoryRepository.Update(ctx, &category),
	)

	return category
}

func (u *categoryUseCase) Delete(ctx context.Context, request dto_request.CategoryDeleteRequest) {
	category := u.baseUseCase.mustGetCategory(ctx, request.CategoryId, panicIsPath)

	if category.IsGlobal {
		panic(dto_response.NewBadRequestResponse("CATEGORY.IS_DEFAULT_CATEGORY"))
	}

	if !category.IsGlobal && *category.UserId != model.MustGetUserCtx(ctx).Id {
		panic(dto_response.NewForbiddenResponse("FORBIDDEN"))
	}

	u.mustValidateAllowDelete(ctx, request.CategoryId)

	panicIfErr(
		u.categoryRepository.Delete(ctx, &category),
	)
}

func (u *categoryUseCase) OptionForTransactionForm(ctx context.Context, request dto_request.CategoryOptionForTransactionFormRequest) ([]model.Category, int) {
	currentUser := model.MustGetUserCtx(ctx)

	queryOption := model.CategoryQueryOption{
		QueryOption:   model.NewBasicQueryOption(request.Limit, request.Page, model.Sorts(request.Sorts)),
		IncludeGlobal: util.BoolP(true),
		IsExpense:     request.IsExpense,
		UserId:        &currentUser.Id,
		Phrase:        request.Phrase,
	}

	categories, err := u.categoryRepository.Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.categoryRepository.Count(ctx, queryOption)
	panicIfErr(err)

	return categories, total
}
