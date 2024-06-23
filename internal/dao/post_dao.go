package dao

import (
	"context"
	"database/sql"
	"fmt"
	"taskOzon/graph/model"
	"taskOzon/internal/models"
)

type PostCRUD interface {
	AddPost(ctx context.Context, title string, content string, authorID uint32, commentsAllowed bool) (model.Post, error)
	GetPost(ctx context.Context, postID uint32) (model.Post, error)
	GetPosts(ctx context.Context, page uint32, itemsByPage uint32, strategy models.Strategy) ([]model.Post, error)
	EditPost(ctx context.Context, postID uint32, title string, content string, commentsAllowed bool) error
	DeletePost(ctx context.Context, postID uint32) error
}

type PostDAO struct {
	DB *sql.DB
}

func NewPostDao(DB *sql.DB) *PostDAO {
	return &PostDAO{DB: DB}
}

func (dao *PostDAO) AddPost(ctx context.Context, title string, content string, authorID uint32, commentsAllowed bool) (model.Post, error) {
	query := `INSERT INTO posts (title, content, author_id, comments_allowed) VALUES ($1, $2, $3, $4) RETURNING id, title, content, author_id, comments_allowed`
	var post model.Post
	err := dao.DB.QueryRowContext(ctx, query, title, content, authorID, commentsAllowed).Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.CommentsAllowed)
	if err != nil {
		return post, fmt.Errorf("error inserting post: %v", err)
	}
	return post, nil
}

func (dao *PostDAO) GetPost(ctx context.Context, postID uint32) (model.Post, error) {
	query := `SELECT p.id, title, content, author_id, comments_allowed, u.username FROM posts p INNER JOIN users u ON p.author_id = u.id WHERE p.id = $1`
	var post model.Post
	post.Author = &model.User{}
	err := dao.DB.QueryRowContext(ctx, query, postID).Scan(&post.ID, &post.Title, &post.Content, &post.Author.ID, &post.CommentsAllowed, &post.Author.Username)
	if err != nil {
		return post, fmt.Errorf("error fetching post: %v", err)
	}
	return post, nil
}

func (dao *PostDAO) GetPosts(ctx context.Context, page uint32, itemsByPage uint32, strategy models.Strategy) ([]model.Post, error) {
	offset := (page - 1) * itemsByPage
	query := `SELECT id, title, content, author_id, comments_allowed FROM posts LIMIT $1 OFFSET $2`
	rows, err := dao.DB.QueryContext(ctx, query, itemsByPage, offset)
	if err != nil {
		return nil, fmt.Errorf("error fetching posts: %v", err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.CommentsAllowed); err != nil {
			return nil, fmt.Errorf("error scanning post: %v", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating posts: %v", err)
	}
	return posts, nil
}

func (dao *PostDAO) EditPost(ctx context.Context, postID uint32, title string, content string, commentsAllowed bool) error {
	query := `UPDATE posts SET title = $1, content = $2, comments_allowed = $3 WHERE id = $4`
	_, err := dao.DB.ExecContext(ctx, query, title, content, commentsAllowed, postID)
	if err != nil {
		return fmt.Errorf("error updating post: %v", err)
	}
	return nil
}

func (dao *PostDAO) DeletePost(ctx context.Context, postID uint32) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := dao.DB.ExecContext(ctx, query, postID)
	if err != nil {
		return fmt.Errorf("error deleting post: %v", err)
	}
	return nil
}
