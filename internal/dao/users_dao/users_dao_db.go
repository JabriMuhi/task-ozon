package users_dao

import (
	"context"
	"database/sql"
	"fmt"
	"taskOzon/graph/model"
)

type UserDAO struct {
	DB *sql.DB
}

func NewUserDao(DB *sql.DB) *UserDAO {
	return &UserDAO{DB: DB}
}

func (dao *UserDAO) AddUser(ctx context.Context, username string, password string, email string) (int, error) {
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"

	var userID int
	err := dao.DB.QueryRowContext(ctx, query, username, password, email).Scan(&userID)
	if err != nil {
		return userID, fmt.Errorf("error inserting user: %v", err)
	}
	return userID, nil
}

func (dao *UserDAO) GetUser(ctx context.Context, userID int) (string, error) {
	query := "SELECT username FROM users WHERE id = $1"
	var user model.User
	err := dao.DB.QueryRowContext(ctx, query, userID).Scan(&user.Username)
	if err != nil {
		return user.Username, fmt.Errorf("error fetching user: %v", err)
	}
	return user.Username, nil
}

func (dao *UserDAO) DeleteUser(ctx context.Context, userID int) (int, error) {
	query := "DELETE FROM users WHERE id = $1"
	_, err := dao.DB.ExecContext(ctx, query, userID)
	if err != nil {
		return userID, fmt.Errorf("error deleting user: %v", err)
	}
	return userID, nil
}
