package use_case

import (
	"capstone/model"
	"capstone/repository"
	"context"
)

type UserUseCase interface {
	GetMe(ctx context.Context) model.User
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(
	userRepository repository.UserRepository,
) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) GetMe(ctx context.Context) model.User {
	return model.MustGetUserCtx(ctx)
}
