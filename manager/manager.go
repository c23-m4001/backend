package manager

import (
	"capstone/config"
	jwtInternal "capstone/internal/jwt"
)

type ManagerConfig int

const (
	LoadDefault ManagerConfig = 0
)

type Container struct {
	managerConfig ManagerConfig

	infrastructureManager InfrastructureManager
	jwt                   jwtInternal.Jwt
	repositoryManager     RepositoryManager
	useCaseManager        UseCaseManager
}

func (m Container) InfrastructureManager() InfrastructureManager {
	return m.infrastructureManager
}

func (m Container) RepositoryManager() RepositoryManager {
	return m.repositoryManager
}

func (m Container) UseCaseManager() UseCaseManager {
	return m.useCaseManager
}

func (m Container) Close() error {
	return m.InfrastructureManager().Close()
}

func NewContainer(managerConfig ManagerConfig) *Container {
	container := Container{
		managerConfig: managerConfig,
	}
	container.infrastructureManager = newInfrastructureManager()
	container.repositoryManager = newRepositoryManager(container.infrastructureManager)

	container.jwt = jwtInternal.NewJwt(
		config.GetJwtPrivateKeyFilePath(),
		config.GetJwtPublicKeyFilePath(),
	)

	container.useCaseManager = newUseCaseManager(
		container.infrastructureManager,
		container.repositoryManager,
		container.jwt,
	)

	return &container
}
