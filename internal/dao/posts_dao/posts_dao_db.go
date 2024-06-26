package posts_dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"taskOzon/graph/model"
)

type PostDAO struct {
	DB *sql.DB
}

func NewPostDao(DB *sql.DB) *PostDAO {
	return &PostDAO{DB: DB}
}

func (dao *PostDAO) AddPost(ctx context.Context, title string, content string, authorID int, commentsAllowed bool) (int, error) {
	query := "SELECT username FROM users WHERE id = $1"

	var user model.User

	err := dao.DB.QueryRowContext(ctx, query, authorID).Scan(&user.Username)
	if err != nil {
		return 0, errors.New("bad user id or deleted user")
	}

	query = `INSERT INTO posts (title, content, author_id, comments_allowed) VALUES ($1, $2, $3, $4) RETURNING id`

	var postID int

	err = dao.DB.QueryRowContext(ctx, query, title, content, authorID, commentsAllowed).Scan(&postID)
	if err != nil {
		return postID, fmt.Errorf("error inserting post: %v", err)
	}

	return postID, nil
}

func (dao *PostDAO) GetPost(ctx context.Context, postID int) (*model.Post, error) {
	query := `SELECT p.id, title, content, author_id, comments_allowed, u.username FROM posts p INNER JOIN users u ON p.author_id = u.id WHERE p.id = $1`

	var post model.Post

	post.Author = &model.User{}

	err := dao.DB.QueryRowContext(ctx, query, postID).Scan(&post.ID, &post.Title, &post.Content, &post.Author.ID, &post.CommentsAllowed, &post.Author.Username)
	if err != nil || post.Title == "Deleted post" {
		return &post, fmt.Errorf("error fetching post: %v", err)
	}

	return &post, nil
}

func (dao *PostDAO) GetPosts(ctx context.Context, page int, itemsByPage int) ([]*model.Post, error) {
	offset := (page - 1) * itemsByPage

	query := `SELECT p.id, p.title, p.content, p.author_id, p.comments_allowed, u.username FROM posts p INNER JOIN users u ON p.author_id = u.id LIMIT $1 OFFSET $2`

	rows, err := dao.DB.QueryContext(ctx, query, itemsByPage, offset)
	if err != nil {
		return nil, fmt.Errorf("error fetching posts: %v", err)
	}

	defer rows.Close()

	var posts []*model.Post

	for rows.Next() {

		var post model.Post

		post.Author = &model.User{}

		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author.ID, &post.CommentsAllowed, &post.Author.Username); err != nil {
			return nil, fmt.Errorf("error scanning post: %v", err)
		}

		if post.Title == "Deleted post" {
			continue
		}

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating posts: %v", err)
	}

	return posts, nil
}

func (dao *PostDAO) ChangeCommentsAllowed(ctx context.Context, postID int, commentsAllowed bool) (int, error) {
	query := "UPDATE posts SET comments_allowed = $2 WHERE id = $1 and title != 'Deleted post'"

	_, err := dao.DB.ExecContext(ctx, query, postID, commentsAllowed)
	if err != nil {
		return postID, fmt.Errorf("error updating post: %v", err)
	}

	return postID, nil
}

func (dao *PostDAO) DeletePost(ctx context.Context, postID int) (int, error) {
	query := `UPDATE posts SET title = 'Deleted post' WHERE id = $1 RETURNING id`

	_, err := dao.DB.ExecContext(ctx, query, postID)
	if err != nil {
		return postID, fmt.Errorf("error deleting post: %v", err)
	}

	return postID, nil
}
