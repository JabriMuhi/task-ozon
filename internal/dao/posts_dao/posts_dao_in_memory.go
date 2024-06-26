package posts_dao

import (
	"context"
	"errors"
	"taskOzon/graph/model"
	"taskOzon/pkg/db/in_memory"
)

type PostDAOInMemory struct {
	IM *in_memory.InMemory
}

func NewPostDaoInMemory(IM *in_memory.InMemory) *PostDAOInMemory {
	return &PostDAOInMemory{IM: IM}
}

func (dao *PostDAOInMemory) AddPost(ctx context.Context, title string, content string, authorID int, commentsAllowed bool) (int, error) {
	author, ok := dao.IM.Users[authorID]

	if !ok || author.Username == "Deleted user" {
		return 0, errors.New("bad user id or deleted user")
	}

	dao.IM.Posts[len(dao.IM.Posts)] = *in_memory.NewPost(len(dao.IM.Posts), title, content, authorID, commentsAllowed)

	return len(dao.IM.Posts) - 1, nil
}

func (dao *PostDAOInMemory) GetPost(ctx context.Context, postID int) (*model.Post, error) {
	post, ok := dao.IM.Posts[postID]

	if !ok || post.Title == "Deleted post" {
		return nil, errors.New("bad post id or deleted post")
	}

	author, ok := dao.IM.Users[post.Author_id]

	if !ok || author.Username == "Deleted user" {
		return nil, errors.New("bad author id or deleted user")
	}

	return &model.Post{
		ID:      post.Id,
		Title:   post.Title,
		Content: post.Content,
		Author: &model.User{
			ID:       author.Id,
			Username: author.Username,
		},
		CommentsAllowed: post.Comments_allowed,
	}, nil

}

func (dao *PostDAOInMemory) GetPosts(ctx context.Context, page int, itemsByPage int) ([]*model.Post, error) {
	offset := (page - 1) * itemsByPage

	response := make([]*model.Post, 0)

	for i := offset; i < offset+itemsByPage; i++ {
		post, err := dao.GetPost(ctx, i)

		if err != nil {
			return response, nil
		}

		if post.Title == "Deleted post" {
			continue
		}

		response = append(response, post)
	}

	return response, nil
}

func (dao *PostDAOInMemory) ChangeCommentsAllowed(ctx context.Context, postID int, commentsAllowed bool) (int, error) {
	post, ok := dao.IM.Posts[postID]

	if !ok || post.Title == "Deleted post" {
		return 0, errors.New("bad post id")
	}

	post.Comments_allowed = commentsAllowed

	return post.Id, nil
}

func (dao *PostDAOInMemory) DeletePost(ctx context.Context, postID int) (int, error) {
	post, ok := dao.IM.Posts[postID]

	if !ok {
		return 0, errors.New("bad post id")
	}

	post.Title = "Deleted post"

	return post.Id, nil
}
