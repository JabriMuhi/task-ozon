package service

import (
	"context"
	"database/sql"
	"taskOzon/graph/model"
	"taskOzon/internal/dao"
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
	AddPost(ctx context.Context, title string, content string, authorID int, commentsAllowed bool) (int, error)
	GetPost(ctx context.Context, postID int) (*model.Post, error)
	GetPosts(ctx context.Context, page int, itemsByPage int /*, strategy models.Strategy*/) ([]*model.Post, error)
	ChangeCommentsAllowed(ctx context.Context, postID int, commentsAllowed bool) (int, error)
	DeletePost(ctx context.Context, postID int) (int, error)
}

func (p *PostServiceImpl) GetPost(ctx context.Context, postID int) (*model.Post, error) {
	return p.postCRUD.GetPost(ctx, postID)
}

func (p *PostServiceImpl) AddPost(ctx context.Context, title string, content string, authorID int, commentsAllowed bool) (int, error) {
	return p.postCRUD.AddPost(ctx, title, content, authorID, commentsAllowed)
}
func (p *PostServiceImpl) GetPosts(ctx context.Context, page int, itemsByPage int /*, strategy models.Strategy*/) ([]*model.Post, error) {
	return p.postCRUD.GetPosts(ctx, page, itemsByPage)
}
func (p *PostServiceImpl) ChangeCommentsAllowed(ctx context.Context, postID int, commentsAllowed bool) (int, error) {
	return p.postCRUD.ChangeCommentsAllowed(ctx, postID, commentsAllowed)
}
func (p *PostServiceImpl) DeletePost(ctx context.Context, postID int) (int, error) {
	return p.postCRUD.DeletePost(ctx, postID)
}
