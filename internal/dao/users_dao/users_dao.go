package users_dao

import (
	"context"
)

type UserCRUD interface {
	AddUser(ctx context.Context, username string, password string, email string) (int, error)
	GetUser(ctx context.Context, userID int) (string, error)
	DeleteUser(ctx context.Context, userID int) (int, error)
}
