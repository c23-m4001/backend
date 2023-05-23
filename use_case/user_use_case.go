package use_case

import (
	"capstone/loader"
	"capstone/model"
	"capstone/repository"
	"context"

	"golang.org/x/sync/errgroup"
)

type UserUseCase interface {
	GetMe(ctx context.Context) model.User
}

type userUseCase struct {
	userRepository   repository.UserRepository
	walletRepository repository.WalletRepository
}

func NewUserUseCase(
	userRepository repository.UserRepository,
	walletRepository repository.WalletRepository,
) UserUseCase {
	return &userUseCase{
		userRepository:   userRepository,
		walletRepository: walletRepository,
	}
}

func (u *userUseCase) mustLoadUserMeData(ctx context.Context, user *model.User) {
	walletLoader := loader.NewWalletloader(u.walletRepository)

	panicIfErr(
		await(func(group *errgroup.Group) {
			group.Go(walletLoader.UserFn(user))
		}),
	)

}

func (u *userUseCase) GetMe(ctx context.Context) model.User {
	user := model.MustGetUserCtx(ctx)

	u.mustLoadUserMeData(ctx, &user)

	return user
}
