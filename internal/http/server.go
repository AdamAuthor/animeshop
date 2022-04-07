package http

import (
	"animeProject/internal/store"
	"context"
)

type Server struct {
	ctx        context.Context
	idleConsCh chan struct{}
	store      store.Store
	Address    string
}

func NewServer(ctx context.Context, address string, store store.Store) *Server {
	return &Server{
		ctx:        ctx,
		Address:    address,
		store:      store,
		idleConsCh: make(chan struct{}),
	}
}
