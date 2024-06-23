package service

import (
	"context"
	"taskOzon/graph/model"
)

type CommentCRUD interface {
	AddComment(ctx context.Context, text string, authorID uint32, postId uint32) error
	AddReply(ctx context.Context, text string, authorID uint32, parentCommentId uint32) error
	GetPostComments(ctx context.Context, postID uint32, startLevel uint32, lastLevel uint32, limit uint32) ([]model.Comment, error)
	GetChildrenComments(ctx context.Context, parentCommentID uint32, startLevel uint32, lastLevel uint32, limit uint32) ([]model.Comment, error)
	EditComment(ctx context.Context, commentID uint32, text string) error
	DeleteComment(ctx context.Context, commentID uint32) error
}
