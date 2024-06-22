package graph

import (
	"sync"
	"taskOzon/graph/model"
)

type Resolver struct {
	mu               sync.Mutex
	commentObservers map[uint32]map[chan *model.Comment]struct{}
}

func NewResolver() *Resolver {
	return &Resolver{
		commentObservers: make(map[uint32]map[chan *model.Comment]struct{}),
	}
}
