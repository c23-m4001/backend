package manager

import "capstone/repository"

type RepositoryManager interface {
	UserAccessTokenRepository() repository.UserAccessTokenRepository
	UserRepository() repository.UserRepository
}

type repositoryManager struct {
	userAccessTokenRepository repository.UserAccessTokenRepository
	userRepository            repository.UserRepository
}

func (m *repositoryManager) UserAccessTokenRepository() repository.UserAccessTokenRepository {
	return m.userAccessTokenRepository
}

func (m *repositoryManager) UserRepository() repository.UserRepository {
	return m.userRepository
}

func newRepositoryManager(infrastructureManager InfrastructureManager) RepositoryManager {
	db := infrastructureManager.GetDB()
	return &repositoryManager{
		userAccessTokenRepository: repository.NewUserAccessTokenRepository(db),
		userRepository:            repository.NewUserRepository(db),
	}
}
