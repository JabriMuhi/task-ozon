package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"taskOzon/internal/dao/users_dao"
)

// Test AddUser method
func TestAddUser(t *testing.T) {
	mc := users_dao.NewUserCRUDMock(t)
	defer mc.MinimockFinish()

	service := UserServiceImpl{userCRUD: mc}

	ctx := context.Background()
	expectedID := 1
	mc.AddUserMock.Expect(ctx, "testuser", "password123", "test@example.com").Return(expectedID, nil)

	userID, err := service.AddUser(ctx, "testuser", "password123", "test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, expectedID, userID)

	mc.MinimockFinish()
}

// Test GetUser method
func TestGetUser(t *testing.T) {
	mc := users_dao.NewUserCRUDMock(t)
	defer mc.MinimockFinish()

	service := UserServiceImpl{userCRUD: mc}

	ctx := context.Background()
	expectedUsername := "testuser"
	mc.GetUserMock.Expect(ctx, 1).Return(expectedUsername, nil)

	username, err := service.GetUser(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsername, username)

	mc.MinimockFinish()
}

// Test DeleteUser method
func TestDeleteUser(t *testing.T) {
	mc := users_dao.NewUserCRUDMock(t)
	defer mc.MinimockFinish()

	service := UserServiceImpl{userCRUD: mc}

	ctx := context.Background()
	expectedID := 1
	mc.DeleteUserMock.Expect(ctx, 1).Return(expectedID, nil)

	deletedID, err := service.DeleteUser(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, deletedID)

	mc.MinimockFinish()
}
