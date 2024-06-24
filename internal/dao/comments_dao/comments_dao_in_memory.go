package comments_dao

import (
	"context"
	"errors"
	"taskOzon/graph/model"
	"taskOzon/pkg/db/in_memory"
)

type CommentDAOInMemory struct {
	IM *in_memory.InMemory
}

func NewCommentDaoInMemory(IM *in_memory.InMemory) *CommentDAOInMemory {
	return &CommentDAOInMemory{IM: IM}
}

func (dao *CommentDAOInMemory) AddComment(ctx context.Context, text string, userID int, postID int) (int, error) {
	dao.IM.Comments[len(dao.IM.Comments)] = *in_memory.NewComment(len(dao.IM.Comments), postID, userID, text)

	dao.IM.CommentsParentChild[len(dao.IM.CommentsParentChild)] = *in_memory.NewCommentParentChild(len(dao.IM.Comments)-1, len(dao.IM.Comments)-1, 0)
	return len(dao.IM.Comments) - 1, nil
}

func (dao *CommentDAOInMemory) AddReply(ctx context.Context, text string, userID int, parentCommentId int) (int, error) {
	dao.IM.Comments[len(dao.IM.Comments)] = *in_memory.NewComment(len(dao.IM.Comments), -1, userID, text)

	dao.IM.CommentsParentChild[len(dao.IM.CommentsParentChild)] = *in_memory.NewCommentParentChild(len(dao.IM.Comments)-1, len(dao.IM.Comments)-1, 0)

	return 1, nil
}

func (dao *CommentDAOInMemory) GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {

	return []model.Comment{}, nil
}

func (dao *CommentDAOInMemory) GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {

	return []model.Comment{}, nil
}

func (dao *CommentDAOInMemory) DeleteComment(ctx context.Context, commentID int) (int, error) {
	comment, ok := dao.IM.Posts[commentID]

	if !ok {
		return 0, errors.New("bad comment id")
	}
	comment.Content = "Deleted comment"

	return 1, nil
}
