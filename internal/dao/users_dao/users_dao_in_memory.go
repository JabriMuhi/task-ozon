package users_dao

import (
	"context"
	"errors"
	"taskOzon/graph/model"
	"taskOzon/pkg/db/in_memory"
)

type UserDAOInMemory struct {
	IM *in_memory.InMemory
}

func NewUserDaoInMemory(IM *in_memory.InMemory) *UserDAOInMemory {
	return &UserDAOInMemory{IM: IM}
}

func NewUserModel(id int, username string, password string) *model.User {
	return &model.User{
		ID:       id,
		Username: username,
		Password: password,
	}
}

func (dao *UserDAOInMemory) AddUser(ctx context.Context, username string, password string, email string) (int, error) {
	dao.IM.Users[len(dao.IM.Users)] = *in_memory.NewUser(len(dao.IM.Users), username, password, email)
	return len(dao.IM.Users) - 1, nil
}

func (dao *UserDAOInMemory) GetUser(ctx context.Context, userID int) (string, error) {
	user, ok := dao.IM.Users[userID]

	if !ok {
		return "", errors.New("bad user id")
	}

	return user.Username, nil
}

func (dao *UserDAOInMemory) DeleteUser(ctx context.Context, userID int) (int, error) {
	return 1, nil
}
