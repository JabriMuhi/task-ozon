package service

import (
	"context"
	"database/sql"
	"taskOzon/internal/dao"
)

type UserServiceImpl struct {
	userCRUD dao.UserCRUD
}

func InitUserService(db *sql.DB) *UserServiceImpl {
	return &UserServiceImpl{
		userCRUD: dao.NewUserDao(db),
	}
}

type UserService interface {
	AddUser(ctx context.Context, username string, password string, email string) (int, error)
	GetUser(ctx context.Context, userID int) (string, error)
	DeleteUser(ctx context.Context, userID int) (int, error)
}

func (u *UserServiceImpl) AddUser(ctx context.Context, username string, password string, email string) (int, error) {
	return u.userCRUD.AddUser(ctx, username, password, email)
}

func (u *UserServiceImpl) GetUser(ctx context.Context, userID int) (string, error) {
	return u.userCRUD.GetUser(ctx, userID)
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, userID int) (int, error) {
	return u.userCRUD.DeleteUser(ctx, userID)
}
