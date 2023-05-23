package manager

import (
	geoIpInternal "capstone/internal/geoip"
	jwtInternal "capstone/internal/jwt"
	"capstone/use_case"
)

type UseCaseManager interface {
	AuthUseCase() use_case.AuthUseCase
	CategoryUseCase() use_case.CategoryUseCase
	TransactionUseCase() use_case.TransactionUseCase
	UserUseCase() use_case.UserUseCase
	WalletUseCase() use_case.WalletUseCase
}

type useCaseManager struct {
	authUseCase        use_case.AuthUseCase
	categoryUseCase    use_case.CategoryUseCase
	transactionUseCase use_case.TransactionUseCase
	userUseCase        use_case.UserUseCase
	walletUseCase      use_case.WalletUseCase
}

func (m *useCaseManager) AuthUseCase() use_case.AuthUseCase {
	return m.authUseCase
}

func (m *useCaseManager) CategoryUseCase() use_case.CategoryUseCase {
	return m.categoryUseCase
}

func (m *useCaseManager) TransactionUseCase() use_case.TransactionUseCase {
	return m.transactionUseCase
}

func (m *useCaseManager) UserUseCase() use_case.UserUseCase {
	return m.userUseCase
}

func (m *useCaseManager) WalletUseCase() use_case.WalletUseCase {
	return m.walletUseCase
}

func newUseCaseManager(
	infrastructureManager InfrastructureManager,
	repositoryManager RepositoryManager,
	geoIp geoIpInternal.GeoIp,
	jwt jwtInternal.Jwt,
) UseCaseManager {
	baseUseCase := use_case.NewBaseUseCase(
		repositoryManager.CategoryRepository(),
		repositoryManager.TransactionRepository(),
		repositoryManager.WalletRepository(),
	)

	return &useCaseManager{
		authUseCase: use_case.NewAuthUseCase(
			repositoryManager.UserAccessTokenRepository(),
			repositoryManager.UserRepository(),
			geoIp,
			jwt,
		),
		categoryUseCase: use_case.NewCategoryUseCase(
			repositoryManager.CategoryRepository(),
			baseUseCase,
		),
		transactionUseCase: use_case.NewTransactionUseCase(
			repositoryManager.CompositeRepository(),
			repositoryManager.TransactionRepository(),
			baseUseCase,
		),
		userUseCase: use_case.NewUserUseCase(
			repositoryManager.UserRepository(),
		),
		walletUseCase: use_case.NewWalletUseCase(
			repositoryManager.CompositeRepository(),
			repositoryManager.WalletRepository(),
			baseUseCase,
		),
	}
}
