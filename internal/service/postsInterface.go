package service

import (
	"context"
	"taskOzon/graph/model"
	"taskOzon/internal/models"
)

type PostCRUD interface {
	AddPost(ctx context.Context, text string, authorID uint32) error
	GetPost(ctx context.Context, postID uint32, level uint32) (model.Post, error)
	GetPosts(ctx context.Context, page uint32, itemsByPage uint32, strategy models.Strategy) ([]model.Post, error)
	EditPost(ctx context.Context, postID uint32, text string) error
	DeletePost(ctx context.Context, postID uint32) error
}
