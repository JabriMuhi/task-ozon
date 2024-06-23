package service

import (
	"context"
	"database/sql"
	"taskOzon/graph/model"
	"taskOzon/internal/dao"
	"taskOzon/internal/models"
)

type PostServiceImpl struct {
	postCRUD dao.PostCRUD
}

func InitPostService(db *sql.DB) *PostServiceImpl {
	return &PostServiceImpl{
		postCRUD: dao.NewPostDao(db),
	}
}

type PostService interface {
	AddPost(ctx context.Context, text string, authorID uint32) error
	GetPost(ctx context.Context, postID uint32) (model.Post, error)
	GetPosts(ctx context.Context, page uint32, itemsByPage uint32, strategy models.Strategy) ([]model.Post, error)
	EditPost(ctx context.Context, postID uint32, text string) error
	DeletePost(ctx context.Context, postID uint32) error
}

func (p *PostServiceImpl) GetPost(ctx context.Context, postID uint32) (model.Post, error) {
	return p.postCRUD.GetPost(ctx, postID)
}

func (p *PostServiceImpl) AddPost(ctx context.Context, text string, authorID uint32) error {
	return nil
}
func (p *PostServiceImpl) GetPosts(ctx context.Context, page uint32, itemsByPage uint32, strategy models.Strategy) ([]model.Post, error) {
	return []model.Post{}, nil
}
func (p *PostServiceImpl) EditPost(ctx context.Context, postID uint32, text string) error {
	return nil
}
func (p *PostServiceImpl) DeletePost(ctx context.Context, postID uint32) error {
	return nil
}
