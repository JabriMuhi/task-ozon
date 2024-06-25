package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"taskOzon/graph/model"
	"taskOzon/internal/dao/comments_dao"
)

// Test AddComment method
func TestAddComment(t *testing.T) {
	mc := comments_dao.NewCommentCRUDMock(t)
	defer mc.MinimockFinish()

	service := CommentServiceImpl{commentCRUD: mc}

	ctx := context.Background()
	expectedID := 1
	mc.AddCommentMock.Expect(ctx, "Test comment", 1, 1).Return(expectedID, nil)

	commentID, err := service.AddComment(ctx, "Test comment", 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, commentID)

	mc.MinimockFinish()
}

// Test AddReply method
func TestAddReply(t *testing.T) {
	mc := comments_dao.NewCommentCRUDMock(t)
	defer mc.MinimockFinish()

	service := CommentServiceImpl{commentCRUD: mc}

	ctx := context.Background()
	expectedID := 1
	mc.AddReplyMock.Expect(ctx, "Test reply", 1, 1).Return(expectedID, nil)

	replyID, err := service.AddReply(ctx, "Test reply", 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, replyID)

	mc.MinimockFinish()
}

// Test GetPostComments method
func TestGetPostComments(t *testing.T) {
	mc := comments_dao.NewCommentCRUDMock(t)
	defer mc.MinimockFinish()

	service := CommentServiceImpl{commentCRUD: mc}

	ctx := context.Background()
	expectedComments := []model.Comment{
		{ID: 1, Content: "Comment 1"},
		{ID: 2, Content: "Comment 2"},
	}
	mc.GetPostCommentsMock.Expect(ctx, 1, 0, 10, 20).Return(expectedComments, nil)

	comments, err := service.GetPostComments(ctx, 1, 0, 10, 20)
	assert.NoError(t, err)
	assert.Equal(t, expectedComments, comments)

	mc.MinimockFinish()
}

// Test GetChildrenComments method
func TestGetChildrenComments(t *testing.T) {
	mc := comments_dao.NewCommentCRUDMock(t)
	defer mc.MinimockFinish()

	service := CommentServiceImpl{commentCRUD: mc}

	ctx := context.Background()
	expectedComments := []model.Comment{
		{ID: 1, Content: "Child Comment 1"},
		{ID: 2, Content: "Child Comment 2"},
	}
	mc.GetChildrenCommentsMock.Expect(ctx, 1, 0, 10, 20).Return(expectedComments, nil)

	comments, err := service.GetChildrenComments(ctx, 1, 0, 10, 20)
	assert.NoError(t, err)
	assert.Equal(t, expectedComments, comments)

	mc.MinimockFinish()
}

// Test DeleteComment method
func TestDeleteComment(t *testing.T) {
	mc := comments_dao.NewCommentCRUDMock(t)
	defer mc.MinimockFinish()

	service := CommentServiceImpl{commentCRUD: mc}

	ctx := context.Background()
	expectedID := 1
	mc.DeleteCommentMock.Expect(ctx, 1).Return(expectedID, nil)

	deletedID, err := service.DeleteComment(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, deletedID)

	mc.MinimockFinish()
}
