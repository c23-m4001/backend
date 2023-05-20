package seeder

import (
	"capstone/manager"
	"capstone/model"
	"fmt"
)

var Seeders = map[string]func(repositoryManager manager.RepositoryManager){
	model.CategoryTableName: CategorySeed,
	model.UserTableName:     UserSeed,
	model.WalletTableName:   WalletSeed,
}

func Seed(repositoryManager manager.RepositoryManager, tableName string) {
	if seed, exist := Seeders[tableName]; exist {
		seed(repositoryManager)
	} else {
		fmt.Printf("Seeder for table `%s` not found\n", tableName)
	}
}

func SeedAll(repositoryManager manager.RepositoryManager) {
	seedOrders := []string{
		model.UserTableName,

		model.CategoryTableName,
		model.WalletTableName,
	}

	for _, tableName := range seedOrders {
		seed, ok := Seeders[tableName]
		if !ok {
			panic(fmt.Errorf("table name %s not found", tableName))
		}
		seed(repositoryManager)
	}
}
