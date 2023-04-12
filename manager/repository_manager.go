package manager

import "capstone/repository"

type RepositoryManager interface {
	UserRepository() repository.UserRepository
}

type repositoryManager struct {
	userRepository repository.UserRepository
}

func (m *repositoryManager) UserRepository() repository.UserRepository {
	return m.userRepository
}

func NewRepositoryManager(infrastructureManager *infrastructureManager) RepositoryManager {
	db := infrastructureManager.GetDB()
	return &repositoryManager{
		userRepository: repository.NewUserRepository(db),
	}
}
