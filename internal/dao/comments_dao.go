package dao

import (
	"context"
	"database/sql"
	"fmt"
	"taskOzon/graph/model"
)

type CommentCRUD interface {
	AddComment(ctx context.Context, text string, userID int, postID int) (int, error)
	AddReply(ctx context.Context, text string, userID int, parentCommentID int) (int, error)
	GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error)
	GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error)
	EditComment(ctx context.Context, commentID int, text string) (int, error)
	DeleteComment(ctx context.Context, commentID int) (int, error)
}

type CommentDAO struct {
	DB *sql.DB
}

func NewCommentDao(DB *sql.DB) *CommentDAO {
	return &CommentDAO{DB: DB}
}

func (dao *CommentDAO) AddComment(ctx context.Context, text string, userID int, postID int) (int, error) {
	query := `INSERT INTO comments (text, user_id, post_id) VALUES ($1, $2, $3) RETURNING id`

	var commentID int

	err := dao.DB.QueryRowContext(ctx, query, text, userID, postID).Scan(&commentID)
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
	addQuery := `INSERT INTO comments (text, user_id) VALUES ($1, $2) RETURNING id`

	var commentID int

	err := dao.DB.QueryRowContext(ctx, addQuery, text, userID).Scan(&commentID)
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
	query := `SELECT c.user_id, c.text
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

		err = rows.Scan(&comment.Author.ID, &comment.Content)

		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (dao *CommentDAO) GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {
	query := `SELECT c.user_id, c.text
FROM
    public.comments AS c
        JOIN public.comments_parent_childs_ids p ON c.id = p.children_id AND p.level >= $1 AND p.level <= $2
WHERE
    p.parent_id = $3
ORDER BY p.level
LIMIT $4;`

	rows, err := dao.DB.QueryContext(ctx, query, startLevel, lastLevel, parentCommentID, limit)
	if err != nil {
		return nil, fmt.Errorf("error querying children comments: %v", err)
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		comment := model.Comment{Author: &model.User{}}
		err := rows.Scan(&comment.Author.ID, &comment.Content)
		if err != nil {
			return nil, fmt.Errorf("error scanning comment: %v", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows iteration: %v", err)
	}

	return comments, nil
}
func (dao *CommentDAO) EditComment(ctx context.Context, commentID int, text string) (int, error) {
	return 1, nil
}
func (dao *CommentDAO) DeleteComment(ctx context.Context, commentID int) (int, error) {
	return 1, nil
}
