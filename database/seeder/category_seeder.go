package seeder

import (
	"capstone/data_type"
	"capstone/manager"
	"capstone/model"
	"context"
)

var (
	// Category Default
	CategoryDefaultOne = model.Category{
		Id:        "413be78a-290c-4cde-a4d1-5d769a5c9448",
		UserId:    nil,
		Name:      "Makanan & Minuman",
		IsGlobal:  true,
		IsExpense: true,
		LogoType:  data_type.CategoryLogoTypeFoodBeverage,
	}
	CategoryDefaultTwo = model.Category{
		Id:        "7f36b362-ba9e-446a-811d-e2ebc72971da",
		UserId:    nil,
		Name:      "Transportasi",
		IsGlobal:  true,
		IsExpense: true,
		LogoType:  data_type.CategoryLogoTypeTransportation,
	}
	CategoryDefaultThree = model.Category{
		Id:        "decedad9-8350-4a64-8a8b-82379aea3d4f",
		UserId:    nil,
		Name:      "Gaji",
		IsGlobal:  true,
		IsExpense: false,
		LogoType:  data_type.CategoryLogoTypeOther,
	}
	CategoryDefaultFour = model.Category{
		Id:        "709d4a02-f906-4d72-8228-32b99197db91",
		UserId:    nil,
		Name:      "Transfer Keluar",
		IsGlobal:  true,
		IsExpense: false,
		LogoType:  data_type.CategoryLogoTypeOther,
	}
	CategoryDefaultFive = model.Category{
		Id:        "9358754d-3917-4589-8098-b22ecda1f588",
		UserId:    nil,
		Name:      "Transfer Masuk",
		IsGlobal:  true,
		IsExpense: true,
		LogoType:  data_type.CategoryLogoTypeOther,
	}

	// Category User
	CategoryOne = model.Category{
		Id:        "9df43422-caf6-42ec-bf18-8fa97213efab",
		UserId:    &UserOne.Id,
		Name:      "Makanan",
		IsGlobal:  false,
		IsExpense: true,
		LogoType:  data_type.CategoryLogoTypeFoodBeverage,
	}
	CategoryTwo = model.Category{
		Id:        "7de25382-f9a7-40ee-9f19-21157e724181",
		UserId:    &UserOne.Id,
		Name:      "Gaji",
		IsGlobal:  false,
		IsExpense: false,
		LogoType:  data_type.CategoryLogoTypeOther,
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
		CategoryDefaultOne,
		CategoryDefaultTwo,
		CategoryDefaultThree,
		CategoryDefaultFour,
		CategoryDefaultFive,
		CategoryOne,
		CategoryTwo,
	}); err != nil {
		panic(err)
	}
}
