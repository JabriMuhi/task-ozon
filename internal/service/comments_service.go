package service

import (
	"context"
	"database/sql"
	"taskOzon/graph/model"
	"taskOzon/internal/dao"
)

type CommentServiceImpl struct {
	commentCRUD dao.CommentCRUD
}

func InitCommentService(db *sql.DB) *CommentServiceImpl {
	return &CommentServiceImpl{
		commentCRUD: dao.NewCommentDao(db),
	}
}

type CommentService interface {
	AddComment(ctx context.Context, text string, userID int, postID int) (int, error)
	AddReply(ctx context.Context, text string, userID int, parentCommentID int) (int, error)
	GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error)
	GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error)
	EditComment(ctx context.Context, commentID int, text string) (int, error)
	DeleteComment(ctx context.Context, commentID int) (int, error)
}

func (c *CommentServiceImpl) AddComment(ctx context.Context, text string, userID int, postID int) (int, error) {
	return c.commentCRUD.AddComment(ctx, text, userID, postID)
}

func (c *CommentServiceImpl) AddReply(ctx context.Context, text string, userID int, parentCommentID int) (int, error) {
	return c.commentCRUD.AddReply(ctx, text, userID, parentCommentID)
}

func (c *CommentServiceImpl) GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {
	return c.commentCRUD.GetPostComments(ctx, postID, startLevel, lastLevel, limit)
}

func (c *CommentServiceImpl) GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {
	return nil, nil
}
func (c *CommentServiceImpl) EditComment(ctx context.Context, commentID int, text string) (int, error) {
	return 1, nil
}
func (c *CommentServiceImpl) DeleteComment(ctx context.Context, commentID int) (int, error) {
	return 1, nil
}
