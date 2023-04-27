package manager

import "capstone/repository"

type RepositoryManager interface {
	UserAccessTokenRepository() repository.UserAccessTokenRepository
	UserRepository() repository.UserRepository
	WalletRepository() repository.WalletRepository
}

type repositoryManager struct {
	userAccessTokenRepository repository.UserAccessTokenRepository
	userRepository            repository.UserRepository
	walletRepository          repository.WalletRepository
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
		userAccessTokenRepository: repository.NewUserAccessTokenRepository(db),
		userRepository:            repository.NewUserRepository(db),
		walletRepository:          repository.NewWalletRepository(db),
	}
}
