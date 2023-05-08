package production_seeder

import (
	"capstone/manager"
	"fmt"
)

var Seeders = map[string]func(repositoryManager manager.RepositoryManager){}

func Seed(repositoryManager manager.RepositoryManager, tableName string) {
	if seed, exist := Seeders[tableName]; exist {
		seed(repositoryManager)
	} else {
		fmt.Printf("Seeder for table `%s` not found\n", tableName)
	}
}

func SeedAll(repositoryManager manager.RepositoryManager) {
	seedOrders := []string{}

	for _, tableName := range seedOrders {
		seed, ok := Seeders[tableName]
		if !ok {
			panic(fmt.Errorf("table name %s not found", tableName))
		}
		seed(repositoryManager)
	}
}
