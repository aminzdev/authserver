package authserver

import (
	"context"
	"errors"
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver/protocol"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *authServer) ValidateSession(_ context.Context, req *protocol.ValidateSessionReq) (*protocol.ValidateSessionRes, error) {
	err := a.auth.ValidateSession(req.User.Name, req.User.Code)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrWrongCredentials):
			return nil, status.Error(codes.Unauthenticated, "wrong credentials.")
		case errors.Is(err, auth.ErrSessionNotFound):
			return nil, status.Error(codes.Unauthenticated, "session not found.")
		case errors.Is(err, auth.ErrExpiredSession):
			return nil, status.Error(codes.Unauthenticated, "session is expired.")
		case errors.Is(err, auth.ErrInternal):
			return nil, status.Error(codes.Internal, "internal error.")
		}
		return nil, status.Error(codes.Unknown, "unknown error.")
	}

	return &protocol.ValidateSessionRes{}, nil
}
