package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"taskOzon/graph/model"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link model.Link
	var user model.User
	link.Address = input.Address
	link.Title = input.Title
	user.Username = "test"
	link.User = &user
	return &link, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// AddPost is the resolver for the addPost field.
func (r *mutationResolver) AddPost(ctx context.Context, title string, content string, commentsAllowed bool, userID int) (int, error) {
	//postService := service.InitPostService()
	//postDao := dao.NewPostDao(r.DB)
	//post, _ := postDao.GetPost(ctx, uint32(postID))
	//return &post, nil
	return r.PostService.AddPost(ctx, title, content, userID, commentsAllowed)
}

// ChangeCommentsAllowed is the resolver for the changeCommentsAllowed field.
func (r *mutationResolver) ChangeCommentsAllowed(ctx context.Context, postID int, commentsAllowed bool) (int, error) {
	return r.PostService.ChangeCommentsAllowed(ctx, postID, commentsAllowed)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, postID int) (int, error) {
	return r.PostService.DeletePost(ctx, postID)
}

// AddComment is the resolver for the addComment field.
func (r *mutationResolver) AddComment(ctx context.Context, postID int, parentID *int, content string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: AddComment - addComment"))
}

// ToggleCommentsAllowed is the resolver for the toggleCommentsAllowed field.
func (r *mutationResolver) ToggleCommentsAllowed(ctx context.Context, postID int, commentsAllowed bool) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: ToggleCommentsAllowed - toggleCommentsAllowed"))
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link
	dummyLink := model.Link{
		Title:   "our dummy link",
		Address: "https://address.org",
		User:    &model.User{Username: "admin"},
	}
	links = append(links, &dummyLink)
	return links, nil
}

// GetPost is the resolver for the getPost field.
func (r *queryResolver) GetPost(ctx context.Context, postID int) (*model.Post, error) {
	return r.PostService.GetPost(ctx, postID)
	//postDao := dao.NewPostDao(r.DB)
	//post, _ := postDao.GetPost(ctx, uint32(postID))
	//return &post, nil
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, page int, itemsByPage int) ([]*model.Post, error) {
	return r.PostService.GetPosts(ctx, page, itemsByPage)
}

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID int) (<-chan *model.Comment, error) {
	commentChan := make(chan *model.Comment)

	r.mu.Lock()
	if _, ok := r.commentObservers[postID]; !ok {
		r.commentObservers[postID] = make(map[chan *model.Comment]struct{})
	}
	r.commentObservers[postID][commentChan] = struct{}{}
	r.mu.Unlock()

	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(r.commentObservers[postID], commentChan)
		if len(r.commentObservers[postID]) == 0 {
			delete(r.commentObservers, postID)
		}
		r.mu.Unlock()
		close(commentChan)
	}()

	return commentChan, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
