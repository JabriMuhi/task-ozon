package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"taskOzon/graph/model"
	"taskOzon/internal/dao/posts_dao"
)

// Test AddPost method
func TestAddPost(t *testing.T) {
	mc := posts_dao.NewPostCRUDMock(t)
	defer mc.MinimockFinish()

	service := PostServiceImpl{postCRUD: mc}

	ctx := context.Background()
	expectedID := 1
	mc.AddPostMock.Expect(ctx, "Test title", "Test content", 1, true).Return(expectedID, nil)

	postID, err := service.AddPost(ctx, "Test title", "Test content", 1, true)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, postID)

	mc.MinimockFinish()
}

// Test GetPost method
func TestGetPost(t *testing.T) {
	mc := posts_dao.NewPostCRUDMock(t)
	defer mc.MinimockFinish()

	service := PostServiceImpl{postCRUD: mc}

	ctx := context.Background()
	expectedPost := &model.Post{ID: 1, Title: "Test title", Content: "Test content"}
	mc.GetPostMock.Expect(ctx, 1).Return(expectedPost, nil)

	post, err := service.GetPost(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedPost, post)

	mc.MinimockFinish()
}

// Test GetPosts method
func TestGetPosts(t *testing.T) {
	mc := posts_dao.NewPostCRUDMock(t)
	defer mc.MinimockFinish()

	service := PostServiceImpl{postCRUD: mc}

	ctx := context.Background()
	expectedPosts := []*model.Post{
		{ID: 1, Title: "Post 1", Content: "Content 1"},
		{ID: 2, Title: "Post 2", Content: "Content 2"},
	}
	mc.GetPostsMock.Expect(ctx, 1, 10).Return(expectedPosts, nil)

	posts, err := service.GetPosts(ctx, 1, 10)
	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, posts)

	mc.MinimockFinish()
}

// Test ChangeCommentsAllowed method
func TestChangeCommentsAllowed(t *testing.T) {
	mc := posts_dao.NewPostCRUDMock(t)
	defer mc.MinimockFinish()

	service := PostServiceImpl{postCRUD: mc}

	ctx := context.Background()
	expectedID := 1
	mc.ChangeCommentsAllowedMock.Expect(ctx, 1, false).Return(expectedID, nil)

	changedID, err := service.ChangeCommentsAllowed(ctx, 1, false)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, changedID)

	mc.MinimockFinish()
}

// Test DeletePost method
func TestDeletePost(t *testing.T) {
	mc := posts_dao.NewPostCRUDMock(t)
	defer mc.MinimockFinish()

	service := PostServiceImpl{postCRUD: mc}

	ctx := context.Background()
	expectedID := 1
	mc.DeletePostMock.Expect(ctx, 1).Return(expectedID, nil)

	deletedID, err := service.DeletePost(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, deletedID)

	mc.MinimockFinish()
}
