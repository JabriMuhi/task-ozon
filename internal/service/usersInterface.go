package service

import (
	"context"
)

type UserCRUD interface {
	AddUser(ctx context.Context, username string, password string, email string) error
	DeleteUser(ctx context.Context, userID uint32) error
}
