package seeder

import (
	"capstone/manager"
	"capstone/model"
	"context"
)

var (
	CategoryOne = model.Category{
		Id:        "9df43422-caf6-42ec-bf18-8fa97213efab",
		UserId:    &UserOne.Id,
		Name:      "Makanan",
		IsGlobal:  false,
		IsExpense: true,
	}
	CategoryTwo = model.Category{
		Id:        "7de25382-f9a7-40ee-9f19-21157e724181",
		UserId:    &UserOne.Id,
		Name:      "Gaji",
		IsGlobal:  false,
		IsExpense: false,
	}
)

func CategorySeed(repositoryManager manager.RepositoryManager) {
	categoryRepository := repositoryManager.CategoryRepository()

	count, err := categoryRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	if count > 0 {
		return
	}

	if err := categoryRepository.InsertMany(context.Background(), []model.Category{
		CategoryOne,
		CategoryTwo,
	}); err != nil {
		panic(err)
	}
}
