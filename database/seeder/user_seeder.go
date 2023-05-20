package seeder

import (
	"capstone/manager"
	"capstone/model"
	"context"
)

var (
	UserOne = model.User{
		Id:       "7768da67-cee5-4e76-93d6-38ecf3f4675c",
		Name:     "John Doe",
		Email:    "johndoe@gmail.com",
		Password: "$2a$10$npm3kVkfujp/3VOkQ1vNKuluF5IS44H5zWgusRVmKau3stTdWYGfG", // password
	}
	UserTwo = model.User{
		Id:       "a811e5f0-ca53-43d6-b32e-06cd57b56183",
		Name:     "Jane Doe",
		Email:    "janedoe@gmail.com",
		Password: "$2a$10$Ef/tzwpeDuw5w8X5Sqe2PObWdTwVBsYbKzd/yyCTUF9t/lsTqYuWq", // 123456
	}
)

func UserSeed(repositoryManager manager.RepositoryManager) {
	userRepository := repositoryManager.UserRepository()

	count, err := userRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	if count > 0 {
		return
	}

	if err := userRepository.InsertMany(context.Background(), []model.User{
		UserOne,
		UserTwo,
	}); err != nil {
		panic(err)
	}
}
