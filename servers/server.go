package servers

import (
	"context"
	"protobuf/ganerete-file/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
)

// Implement the books service
type BookServer struct {
	tutorialpb.UnimplementedBookServer
}

// Implement the tutorialpb.BookServer interface
func (s *BookServer) AddNewUser(ctx context.Context, user *tutorialpb.UserRequest) (*tutorialpb.UserResponse, error) {
	return &tutorialpb.UserResponse{
		Success: true,
	}, nil
}
