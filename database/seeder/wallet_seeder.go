package seeder

import (
	"capstone/data_type"
	"capstone/manager"
	"capstone/model"
	"context"
)

var (
	WalletOne = model.Wallet{
		Id:          "1d6dd823-328b-4ba4-b859-9841b2874373",
		UserId:      UserOne.Id,
		Name:        "Cash",
		TotalAmount: 0,
		LogoType:    data_type.WalletLogoTypeDefault,
	}
	WalletTwo = model.Wallet{
		Id:          "d2ddbc33-3b4b-4dc3-8eb3-a9709e95e725",
		UserId:      UserOne.Id,
		Name:        "Bank",
		TotalAmount: 0,
		LogoType:    data_type.WalletLogoTypeDefault,
	}
)

func WalletSeed(repositoryManager manager.RepositoryManager) {
	walletRepository := repositoryManager.WalletRepository()

	count, err := walletRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	if count > 0 {
		return
	}

	if err := walletRepository.InsertMany(context.Background(), []model.Wallet{
		WalletOne,
		WalletTwo,
	}); err != nil {
		panic(err)
	}
}
