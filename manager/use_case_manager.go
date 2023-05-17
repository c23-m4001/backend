package manager

import (
	geoIpInternal "capstone/internal/geoip"
	jwtInternal "capstone/internal/jwt"
	"capstone/use_case"
)

type UseCaseManager interface {
	AuthUseCase() use_case.AuthUseCase
	UserUseCase() use_case.UserUseCase
}

type useCaseManager struct {
	authUseCase use_case.AuthUseCase
	userUseCase use_case.UserUseCase
}

func (m *useCaseManager) AuthUseCase() use_case.AuthUseCase {
	return m.authUseCase
}

func (m *useCaseManager) UserUseCase() use_case.UserUseCase {
	return m.userUseCase
}

func newUseCaseManager(
	infrastructureManager InfrastructureManager,
	repositoryManager RepositoryManager,
	geoIp geoIpInternal.GeoIp,
	jwt jwtInternal.Jwt,
) UseCaseManager {
	return &useCaseManager{
		authUseCase: use_case.NewAuthUseCase(
			repositoryManager.UserAccessTokenRepository(),
			repositoryManager.UserRepository(),
			geoIp,
			jwt,
		),
		userUseCase: use_case.NewUserUseCase(
			repositoryManager.UserRepository(),
		),
	}
}
