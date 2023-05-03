package manager

import "capstone/repository"

type RepositoryManager interface {
	CategoryRepository() repository.CategoryRepository
	UserAccessTokenRepository() repository.UserAccessTokenRepository
	UserRepository() repository.UserRepository
	WalletRepository() repository.WalletRepository
}

type repositoryManager struct {
	categoryRepository        repository.CategoryRepository
	userAccessTokenRepository repository.UserAccessTokenRepository
	userRepository            repository.UserRepository
	walletRepository          repository.WalletRepository
}

func (m *repositoryManager) CategoryRepository() repository.CategoryRepository {
	return m.categoryRepository
}

func (m *repositoryManager) UserAccessTokenRepository() repository.UserAccessTokenRepository {
	return m.userAccessTokenRepository
}

func (m *repositoryManager) UserRepository() repository.UserRepository {
	return m.userRepository
}

func (m *repositoryManager) WalletRepository() repository.WalletRepository {
	return m.walletRepository
}

func newRepositoryManager(infrastructureManager InfrastructureManager) RepositoryManager {
	db := infrastructureManager.GetDB()
	return &repositoryManager{
		categoryRepository:        repository.NewCategoryRepository(db),
		userAccessTokenRepository: repository.NewUserAccessTokenRepository(db),
		userRepository:            repository.NewUserRepository(db),
		walletRepository:          repository.NewWalletRepository(db),
	}
}
