package graph

import (
	"sync"
	"taskOzon/graph/model"
	"taskOzon/internal/service"
)

type Resolver struct {
	UserService      service.UserService
	PostService      service.PostService
	CommentService   service.CommentService
	mu               sync.Mutex
	commentObservers map[int]map[chan *model.Comment]struct{}
}

func NewResolver(postService service.PostService, userService service.UserService, commentService service.CommentService) *Resolver {
	return &Resolver{
		UserService:      userService,
		PostService:      postService,
		CommentService:   commentService,
		commentObservers: make(map[int]map[chan *model.Comment]struct{}),
	}
}
