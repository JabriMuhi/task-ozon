package graph

import (
	"sync"
	"taskOzon/graph/model"
	"taskOzon/internal/service"
)

type Resolver struct {
	UserService      service.UserService
	PostService      service.PostService
	mu               sync.Mutex
	commentObservers map[int]map[chan *model.Comment]struct{}
}

func NewResolver(postService service.PostService, userService service.UserService) *Resolver {
	return &Resolver{
		UserService:      userService,
		PostService:      postService,
		commentObservers: make(map[int]map[chan *model.Comment]struct{}),
	}
}
