package authserver

import (
	"context"
	"errors"
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver/protocol"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *authServer) SignOut(_ context.Context, req *protocol.SignOutReq) (*protocol.SignOutRes, error) {
	err := a.auth.SignOut(req.User.Name, req.User.Code)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrWrongCredentials):
			return nil, status.Error(codes.Unauthenticated, "wrong credentials.")
		case errors.Is(err, auth.ErrInternal):
			return nil, status.Error(codes.Internal, "internal error.")
		}
		return nil, status.Error(codes.Unknown, "unknown error.")
	}

	return &protocol.SignOutRes{}, nil
}
