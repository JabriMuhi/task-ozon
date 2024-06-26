package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"
	"log"
	"taskOzon/graph/model"
)

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, username string, password string, email string) (int, error) {
	return r.UserService.AddUser(ctx, username, password, email)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID int) (int, error) {
	return r.UserService.DeleteUser(ctx, userID)
}

// AddPost is the resolver for the addPost field.
func (r *mutationResolver) AddPost(ctx context.Context, title string, content string, commentsAllowed bool, userID int) (int, error) {
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

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, commentID int) (int, error) {
	return r.CommentService.DeleteComment(ctx, commentID)
}

// AddComment is the resolver for the addComment field.
func (r *mutationResolver) AddComment(ctx context.Context, postID int, content string, userID int) (int, error) {
	if len(content) > 2000 {
		return 0, errors.New("invalid content length! max length 2000 chars")
	}
	return r.CommentService.AddComment(ctx, content, userID, postID)
}

// AddReply is the resolver for the addReply field.
func (r *mutationResolver) AddReply(ctx context.Context, postID int, parentCommentID *int, userID int, content string) (int, error) {
	if len(content) > 2000 {
		return 0, errors.New("invalid content length! max length 2000 chars")
	}
	return r.CommentService.AddReply(ctx, content, userID, *parentCommentID)
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
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, page int, itemsByPage int) ([]*model.Post, error) {
	return r.PostService.GetPosts(ctx, page, itemsByPage)
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, userID int) (string, error) {
	return r.UserService.GetUser(ctx, userID)
}

// GetPostComments is the resolver for the getPostComments field.
func (r *queryResolver) GetPostComments(ctx context.Context, postID int, startLevel int, lastLevel int, limit int) ([]*model.Comment, error) {
	c, err := r.CommentService.GetPostComments(ctx, postID, startLevel, lastLevel, limit)
	if err != nil {
		return nil, err
	}

	resp := make([]*model.Comment, 0)

	for key, _ := range c {
		resp = append(resp, &c[key])
	}

	return resp, nil
}

// GetChildrenComments is the resolver for the getChildrenComments field.
func (r *queryResolver) GetChildrenComments(ctx context.Context, parentCommentID int, startLevel int, lastLevel int, limit int) ([]*model.Comment, error) {
	log.Printf("Fetching children comments_dao for parentCommentID=%d, startLevel=%d, lastLevel=%d, limit=%d", parentCommentID, startLevel, lastLevel, limit)

	c, err := r.CommentService.GetChildrenComments(ctx, parentCommentID, startLevel, lastLevel, limit)
	if err != nil {
		log.Printf("Error fetching children comments_dao: %v", err)
		return nil, err
	}

	log.Printf("Comments received from service: %+v", c)

	resp := make([]*model.Comment, 0)

	for key := range c {
		log.Printf("Appending comment: %+v", c[key])
		resp = append(resp, &c[key])
	}

	log.Printf("Fetched comments_dao response: %+v", resp)

	return resp, nil
}

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID int) (<-chan *model.Comment, error) {
	//channel, _ := r.CommentService.NewSubscriber(ctx, postID)
	var tradeChannel = make(chan *model.Comment, 1)

	// context done check
	go func() {
		<-ctx.Done()
	}()

	// run a concurrent routine to send the data to subscribed client
	//go func(tradeChannel chan *model.Comment) {
	//	ticker := time.NewTicker(1 * time.Second)
	//	for {
	//		select {
	//		case <-ticker.C:
	//			tradeChannel <- r.GenerateTradeData()
	//		}
	//	}
	//}(tradeChannel)

	return tradeChannel, nil
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
