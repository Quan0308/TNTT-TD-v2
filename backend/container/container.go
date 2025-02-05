package container

import (
	"github.com/Quan0308/main-api/api"
	"github.com/Quan0308/main-api/core/base"
	"github.com/Quan0308/main-api/handlers"
	"github.com/Quan0308/main-api/interfaces"
	"github.com/Quan0308/main-api/repositories"
	"github.com/Quan0308/main-api/services"
	"github.com/jmoiron/sqlx"
)

type ContainerImpl struct {
	//Repositories
	uow      interfaces.UnitOfWork
	userRepo interfaces.UserRepository

	//External service
	firebaseAuthService interfaces.FirebaseAuthService
}

func NewContainer() interfaces.Container {
	return &ContainerImpl{}
}

func (c *ContainerImpl) Init(db *sqlx.DB) {
	c.uow = base.NewUnitOfWorkImpl(db)
	c.userRepo = repositories.NewUserRepository(c.uow)
	c.firebaseAuthService = services.NewFireBaseAuthServiceImpl()
}

func (c *ContainerImpl) GetAuthAPI() interfaces.AuthAPI {
	authHandler := handlers.NewAuthHandler(c.uow, c.userRepo, c.firebaseAuthService)
	return api.NewAuthAPI(authHandler, c.firebaseAuthService)
}
