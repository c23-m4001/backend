package manager

import (
	jwtInternal "capstone/internal/jwt"
)

type UseCaseManager interface {
}

type useCaseManager struct {
}

func newUseCaseManager(
	infrastructureManager InfrastructureManager,
	repositoryManager RepositoryManager,
	jwt jwtInternal.Jwt,
) UseCaseManager {
	return &useCaseManager{}
}
