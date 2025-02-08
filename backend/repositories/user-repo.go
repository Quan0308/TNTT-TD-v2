package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/Quan0308/main-api/enums"
	"github.com/Quan0308/main-api/interfaces"
	"github.com/Quan0308/main-api/models"
	"github.com/google/uuid"
)

type UserRepositoryHandler struct {
	uow interfaces.UnitOfWork
}

func NewUserRepository(uow interfaces.UnitOfWork) interfaces.UserRepository {
	return &UserRepositoryHandler{uow: uow}
}

func (r *UserRepositoryHandler) GetCurrentUserByUid(ctx context.Context, uid string) (*models.CurrentUser, error) {
	type UserWithRoles struct {
		Id     uuid.UUID  `db:"id"`
		RoleId enums.Role `db:"role_id"`
	}

	var userWithRoles []UserWithRoles
	query := "SELECT users.id, ra.role_id FROM users LEFT JOIN role_assignments ra ON user_id = users.id WHERE firebase_uid = ?"

	err := r.uow.SelectContext(ctx, &userWithRoles, query, uid)

	if err != nil {
		log.Printf("error: %s", err)
		return nil, err
	}

	currentUser := make(map[uuid.UUID]*models.CurrentUser)

	for _, row := range userWithRoles {
		if _, exists := currentUser[row.Id]; !exists {
			currentUser[row.Id] = &models.CurrentUser{
				Id:    row.Id,
				Roles: []enums.Role{},
			}
		}
		if row.RoleId > 0 {
			currentUser[row.Id].Roles = append(currentUser[row.Id].Roles, row.RoleId)
		}
	}
	for _, user := range currentUser {
		return user, nil
	}

	return nil, errors.New("user not found")
}
