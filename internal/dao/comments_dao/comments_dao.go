package comments_dao

import (
	"context"
	"taskOzon/graph/model"
)

type CommentCRUD interface {
	AddComment(ctx context.Context, text string, userID int, postID int) (int, error)
	AddReply(ctx context.Context, text string, userID int, parentCommentID int) (int, error)
	GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error)
	GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error)
	DeleteComment(ctx context.Context, commentID int) (int, error)
}
