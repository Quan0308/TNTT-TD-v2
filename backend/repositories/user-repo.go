package repositories

import (
	"context"

	"github.com/Quan0308/main-api/interfaces"
	"github.com/Quan0308/main-api/models"
)

type UserRepositoryHandler struct {
	uow interfaces.UnitOfWork
}

func NewUserRepository(uow interfaces.UnitOfWork) interfaces.UserRepository {
	return &UserRepositoryHandler{uow: uow}
}

func (r *UserRepositoryHandler) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE email = ?"

	err := r.uow.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
