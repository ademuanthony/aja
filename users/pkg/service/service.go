package service

import (
	"context"
	"github.com/ademuanthony/aja/notificator/pkg/grpc/pb"
)

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	Create(ctx context.Context, email string) error
}

type basicUsersService struct{
	notificatorService pb.NotificatorService
}

func (b *basicUsersService) Create(ctx context.Context, email string) (e0 error) {
	// TODO implement the business logic of Create
	return e0
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {
	return &basicUsersService{}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
