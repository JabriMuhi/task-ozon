package repository

import (
	"context"
	"database/sql"
)

type CommentsRepo struct {
	tx *sql.Tx
}

func NewCommentsRepo(tx *sql.Tx) *CommentsRepo {
	return &CommentsRepo{tx: tx}
}

func (cr *CommentsRepo) AddComment(_ context.Context, text string, authorID uint32, postId uint32) error {
	var query = "insert into comments (post_id, user_id, text) values ($1, $2, $3);"

	_, err := cr.tx.Query(query, postId, authorID, text)

	return err
}

//type CommentCRUD interface {
//	AddComment(ctx context.Context, text string, authorID uint32) error
//	GetPostComments(ctx context.Context, postID uint32, startLevel uint32, lastLevel uint32, limit uint32) ([]model.Comment, error)
//	GetChildrenComments(ctx context.Context, parentCommentID uint32, startLevel uint32, lastLevel uint32, limit uint32) ([]model.Comment, error)
//	EditComment(ctx context.Context, commentID uint32, text string) error
//	DeleteComment(ctx context.Context, commentID uint32) error
//}
