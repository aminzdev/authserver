package authserver

import (
	"context"
	"errors"
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver/protocol"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *authServer) SignUp(_ context.Context, req *protocol.SignUpReq) (*protocol.SignUpRes, error) {
	err := a.auth.SignUp(req.User.Name, req.User.Code)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrInvalidUserName):
			return nil, status.Error(codes.InvalidArgument, "username is invalid.")
		case errors.Is(err, auth.ErrUserAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, "user already exists.")
		case errors.Is(err, auth.ErrInternal):
			return nil, status.Error(codes.Internal, "internal error.")
		}
		return nil, status.Error(codes.Unknown, "unknown error.")
	}

	return &protocol.SignUpRes{}, nil
}
