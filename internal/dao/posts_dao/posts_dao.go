package posts_dao

import (
	"context"
	"taskOzon/graph/model"
)

type PostCRUD interface {
	AddPost(ctx context.Context, title string, content string, authorID int, commentsAllowed bool) (int, error)
	GetPost(ctx context.Context, postID int) (*model.Post, error)
	GetPosts(ctx context.Context, page int, itemsByPage int /*, strategy models.Strategy*/) ([]*model.Post, error)
	ChangeCommentsAllowed(ctx context.Context, postID int, commentsAllowed bool) (int, error)
	DeletePost(ctx context.Context, postID int) (int, error)
}
