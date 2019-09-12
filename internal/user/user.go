package user

import "context"

// Client interface hold the service
type Client interface {
	Service() Service
}

// Service interface definition
type Service interface {
	Login(ctx context.Context, email string) (LoginOutput, error)
	CreateUser(ctx context.Context, email, username string) error
}

// Validator just validate anything
type Validator interface {
	OK() error
}
