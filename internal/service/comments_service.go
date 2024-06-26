package service

import (
	"context"
	"database/sql"
	"taskOzon/graph/model"
	"taskOzon/internal/dao/comments_dao"
	"taskOzon/pkg/db/in_memory"
)

type CommentServiceImpl struct {
	commentCRUD     comments_dao.CommentCRUD
	subscriptionMap map[int][]chan<- model.Comment
}

func InitCommentService(db *sql.DB, subMap map[int][]chan<- model.Comment) *CommentServiceImpl {
	return &CommentServiceImpl{
		commentCRUD:     comments_dao.NewCommentDao(db),
		subscriptionMap: subMap,
	}
}

func InitCommentServiceInMemory(im *in_memory.InMemory, subMap map[int][]chan<- model.Comment) *CommentServiceImpl {
	return &CommentServiceImpl{
		commentCRUD:     comments_dao.NewCommentDaoInMemory(im),
		subscriptionMap: subMap,
	}
}

type CommentService interface {
	AddComment(ctx context.Context, text string, userID int, postID int) (int, error)
	AddReply(ctx context.Context, text string, userID int, parentCommentID int) (int, error)
	GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error)
	GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error)
	DeleteComment(ctx context.Context, commentID int) (int, error)
	NewSubscriber(ctx context.Context, postID int) (chan model.Comment, error)
}

func (c *CommentServiceImpl) AddComment(ctx context.Context, text string, userID int, postID int) (int, error) {
	resp, err := c.commentCRUD.AddComment(ctx, text, userID, postID)
	channels, ok := c.subscriptionMap[postID]
	if ok {
		for _, ch := range channels {
			com := model.Comment{
				ID:      resp,
				PostID:  postID,
				Content: text,
				Author:  &model.User{ID: userID},
			}

			ch <- com
		}
	}
	return resp, err
}

func (c *CommentServiceImpl) AddReply(ctx context.Context, text string, userID int, parentCommentID int) (int, error) {
	return c.commentCRUD.AddReply(ctx, text, userID, parentCommentID)
}

func (c *CommentServiceImpl) GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {
	return c.commentCRUD.GetPostComments(ctx, postID, startLevel, lastLevel, limit)
}

func (c *CommentServiceImpl) GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {
	return c.commentCRUD.GetChildrenComments(ctx, parentCommentID, startLevel, lastLevel, limit)
}

func (c *CommentServiceImpl) DeleteComment(ctx context.Context, commentID int) (int, error) {
	return c.commentCRUD.DeleteComment(ctx, commentID)
}
func (c *CommentServiceImpl) NewSubscriber(ctx context.Context, postID int) (chan model.Comment, error) {
	_, ok := c.subscriptionMap[postID]

	if !ok {
		c.subscriptionMap[postID] = make([]chan<- model.Comment, 0)
	}
	ch := make(chan model.Comment)
	c.subscriptionMap[postID] = append(c.subscriptionMap[postID], ch)

	return ch, nil
}
