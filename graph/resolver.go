package graph

import (
	"sync"
	"taskOzon/graph/model"
	"taskOzon/internal/service"
)

type Resolver struct {
	PostService      service.PostService
	mu               sync.Mutex
	commentObservers map[int]map[chan *model.Comment]struct{}
}

func NewResolver(postService service.PostService) *Resolver {
	return &Resolver{
		PostService:      postService,
		commentObservers: make(map[int]map[chan *model.Comment]struct{}),
	}
}
