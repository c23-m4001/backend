package manager

import (
	geoIpInternal "capstone/internal/geoip"
	jwtInternal "capstone/internal/jwt"
	"capstone/use_case"
)

type UseCaseManager interface {
	AuthUseCase() use_case.AuthUseCase
}

type useCaseManager struct {
	authUseCase use_case.AuthUseCase
}

func (m *useCaseManager) AuthUseCase() use_case.AuthUseCase {
	return m.authUseCase
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
	}
}
