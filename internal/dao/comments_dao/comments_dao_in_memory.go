package comments_dao

import (
	"context"
	"errors"
	"slices"
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

	dao.IM.CommentsParentChild = append(dao.IM.CommentsParentChild, *in_memory.NewCommentParentChild(len(dao.IM.Comments)-1, len(dao.IM.Comments)-1, 0))
	return len(dao.IM.Comments) - 1, nil
}

func (dao *CommentDAOInMemory) AddReply(ctx context.Context, text string, userID int, parentCommentId int) (int, error) {
	dao.IM.Comments[len(dao.IM.Comments)] = *in_memory.NewComment(len(dao.IM.Comments), -1, userID, text)

	dao.IM.CommentsParentChild = append(dao.IM.CommentsParentChild, *in_memory.NewCommentParentChild(len(dao.IM.Comments)-1, len(dao.IM.Comments)-1, 0))

	newCommentParentChilds := make([]in_memory.CommentParentChild, 0)

	for _, v := range dao.IM.CommentsParentChild {
		if v.Children_id == parentCommentId {
			newCommentParentChilds = append(newCommentParentChilds, *in_memory.NewCommentParentChild(v.Parent_id, len(dao.IM.Comments)-1, v.Level+1))
		}
	}

	dao.IM.CommentsParentChild = append(dao.IM.CommentsParentChild, newCommentParentChilds...)

	return len(dao.IM.Comments) - 1, nil
}

func (dao *CommentDAOInMemory) GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {
	resp := make([]model.Comment, 0)

	selectedPostCommentsMap := make(map[int]struct{})

	for _, comment := range dao.IM.Comments {
		if comment.Post_id == postID {
			selectedPostCommentsMap[comment.Id] = struct{}{}
		}
	}

	for _, v := range dao.IM.CommentsParentChild {
		if _, ok := selectedPostCommentsMap[v.Parent_id]; ok {
			if v.Level >= startLevel && v.Level <= lastLevel {
				com := dao.IM.Comments[v.Children_id]

				resp = append(resp, model.Comment{
					Content: com.Text,
					Author: &model.User{
						ID: com.User_id,
					},
					Level: v.Level,
				})
			}
		}
	}

	slices.SortFunc(resp, func(a, b model.Comment) int {
		if a.Level < b.Level {
			return -1
		} else if a.Level > b.Level {
			return 1
		}

		return 0
	})

	if len(resp) > limit {
		return resp[:limit], nil
	}

	return resp, nil
}

func (dao *CommentDAOInMemory) GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]model.Comment, error) {
	resp := make([]model.Comment, 0)

	for _, comment := range dao.IM.CommentsParentChild {
		if comment.Parent_id == parentCommentID && comment.Level >= startLevel && comment.Level <= lastLevel {
			com, ok := dao.IM.Comments[comment.Children_id]
			if ok {
				resp = append(resp, model.Comment{
					ID:      com.Id,
					Content: com.Text,
					Author: &model.User{
						ID: com.User_id,
					},
					Level: comment.Level,
				})
			}
		}
	}

	slices.SortFunc(resp, func(a, b model.Comment) int {
		if a.Level < b.Level {
			return -1
		} else if a.Level > b.Level {
			return 1
		}

		return 0
	})

	if len(resp) > limit {
		return resp[:limit], nil
	}

	return resp, nil
}

func (dao *CommentDAOInMemory) DeleteComment(ctx context.Context, commentID int) (int, error) {
	comment, ok := dao.IM.Posts[commentID]

	if !ok {
		return 0, errors.New("bad comment id")
	}
	comment.Content = "Deleted comment"

	return commentID, nil
}
