package authserver

import (
	"context"
	"errors"
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver/protocol"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (a *authServer) SignIn(_ context.Context, req *protocol.SignInReq) (*protocol.SignInRes, error) {
	err := a.auth.SignIn(req.User.Name, req.User.Code, time.Hour)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrWrongCredentials):
			return nil, status.Error(codes.Unauthenticated, "wrong credentials.")
		case errors.Is(err, auth.ErrInternal):
			return nil, status.Error(codes.Internal, "internal error.")
		}
		return nil, status.Error(codes.Unknown, "unknown error.")
	}

	return &protocol.SignInRes{}, nil
}
