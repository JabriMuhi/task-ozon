package comments_dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"taskOzon/graph/model"
)

type CommentDAO struct {
	DB *sql.DB
}

func NewCommentDao(DB *sql.DB) *CommentDAO {
	return &CommentDAO{DB: DB}
}

func (dao *CommentDAO) AddComment(ctx context.Context, text string, userID int, postID int) (int, error) {
	queryUserCheck := "SELECT username FROM users WHERE id = $1"
	var user model.User
	err := dao.DB.QueryRowContext(ctx, queryUserCheck, userID).Scan(&user.Username)
	if err != nil || user.Username == "Deleted user" {
		return 0, errors.New("bad user id or deleted user")
	}

	queryPostCheck := `SELECT p.id, title, content, author_id, comments_allowed, u.username FROM posts p INNER JOIN users u ON p.author_id = u.id WHERE p.id = $1`
	var post model.Post
	post.Author = &model.User{}
	err = dao.DB.QueryRowContext(ctx, queryPostCheck, postID).Scan(&post.ID, &post.Title, &post.Content, &post.Author.ID, &post.CommentsAllowed, &post.Author.Username)
	if err != nil || post.Title == "Deleted post" {
		return 0, errors.New("bad post id or deleted post")
	}

	query := `INSERT INTO comments (text, user_id, post_id) VALUES ($1, $2, $3) RETURNING id`
	var commentID int
	err = dao.DB.QueryRowContext(ctx, query, text, userID, postID).Scan(&commentID)
	if err != nil {
		return commentID, fmt.Errorf("error inserting comment: %v", err)
	}

	addChildQuery := `INSERT INTO comments_parent_childs_ids (parent_id, children_id, level) VALUES ($1, $1, 0)`
	_, err = dao.DB.Query(addChildQuery, commentID)
	if err != nil {
		return 0, err
	}

	return commentID, nil
}

func (dao *CommentDAO) AddReply(ctx context.Context, text string, userID int, parentCommentId int) (int, error) {

	queryUserCheck := "SELECT username FROM users WHERE id = $1"
	var user model.User

	err := dao.DB.QueryRowContext(ctx, queryUserCheck, userID).Scan(&user.Username)
	if err != nil || user.Username == "Deleted user" {
		return 0, errors.New("bad user id")
	}

	addQuery := `INSERT INTO comments (text, user_id) VALUES ($1, $2) RETURNING id`

	var commentID int

	err = dao.DB.QueryRowContext(ctx, addQuery, text, userID).Scan(&commentID)
	if err != nil {
		return commentID, fmt.Errorf("error inserting comment: %v", err)
	}

	addChildQuery := `INSERT INTO comments_parent_childs_ids (parent_id, children_id, level) VALUES ($1, $1, 0)`

	_, err = dao.DB.Query(addChildQuery, commentID)
	if err != nil {
		return 0, err
	}

	addParentsRelation := `INSERT INTO comments_parent_childs_ids (parent_id, children_id, level) SELECT cp.parent_id, $1, cp.level+1 FROM comments_parent_childs_ids cp WHERE children_id = $2;`

	_, err = dao.DB.Query(addParentsRelation, commentID, parentCommentId)
	if err != nil {
		return 0, err
	}

	return commentID, nil
}

func (dao *CommentDAO) GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {
	query := `SELECT c.user_id, c.text, p.level
FROM
    public.comments AS c
        JOIN public.comments_parent_childs_ids p ON c.id = p.children_id and p.level >= $1 and p.level <= $2
WHERE
    p.parent_id in (select id from comments coms where coms.post_id = $3) ORDER BY p.level LIMIT $4;`

	rows, err := dao.DB.Query(query, startLevel, lastLevel, postID, limit)
	if err != nil {
		return nil, err
	}

	var comments []model.Comment
	for rows.Next() {
		comment := model.Comment{Author: &model.User{}}

		err = rows.Scan(&comment.Author.ID, &comment.Content, &comment.Level)

		if err != nil {
			return nil, err
		}

		if comment.Content == "Deleted comment" {
			continue
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (dao *CommentDAO) GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {
	query := `SELECT c.id, c.text, c.user_id, p.level
FROM
    public.comments AS c
        JOIN public.comments_parent_childs_ids p ON c.id = p.children_id AND p.level >= $1 AND p.level <= $2
WHERE
    p.parent_id = $3
ORDER BY p.level
LIMIT $4;`

	rows, err := dao.DB.Query(query, startLevel, lastLevel, parentCommentID, limit)
	if err != nil {
		return nil, err
	}

	var comments []model.Comment
	for rows.Next() {
		comment := model.Comment{Author: &model.User{}}
		log.Printf(comment.Content)

		err = rows.Scan(&comment.ID, &comment.Content, &comment.Author.ID, &comment.Level)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (dao *CommentDAO) DeleteComment(ctx context.Context, commentID int) (int, error) {
	query := `UPDATE comments SET text = 'Deleted comment' WHERE id = $1 RETURNING id`

	var comment model.Comment
	err := dao.DB.QueryRowContext(ctx, query, commentID).Scan(&comment.ID)

	if err != nil {
		return commentID, fmt.Errorf("error deleting comment: %v", err)
	}

	return comment.ID, nil
}
