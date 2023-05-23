package loader

import (
	"capstone/constant"
	"capstone/model"
	"capstone/repository"
	"context"

	"github.com/graph-gophers/dataloader"
)

type CategoryLoader struct {
	loader dataloader.Loader
}

func (l *CategoryLoader) load(id string) (*model.Category, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.Category), nil
}

func (l *CategoryLoader) TransactionFn(transaction *model.Transaction) func() error {
	return func() error {
		category, err := l.load(transaction.CategoryId)
		if err != nil {
			return err
		}

		transaction.Category = category

		return nil
	}
}

func NewCategoryloader(
	categoryRepository repository.CategoryRepository,
) *CategoryLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		categories, err := categoryRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		categoryById := map[string]model.Category{}
		for _, category := range categories {
			categoryById[category.Id] = category
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var category *model.Category
			if v, ok := categoryById[k.String()]; ok {
				category = &v
			}

			result := &dataloader.Result{Data: category, Error: nil}

			if category == nil {
				result.Error = constant.ErrNoData
			}

			results[idx] = result
		}
		return results
	}

	return &CategoryLoader{
		loader: newDataloader(batchFn),
	}
}
