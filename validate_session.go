package authserver

import (
	"context"
	"errors"
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver/protocol"
)

func (a *authServer) ValidateSession(_ context.Context, req *protocol.ValidateSessionReq) (*protocol.ValidateSessionRes, error) {
	err := a.auth.ValidateSession(req.User.Name, req.User.Code)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrWrongCredentials):
			return nil, ErrUnauthenticated
		case errors.Is(err, auth.ErrSessionNotFound):
			return nil, ErrSessionNotFound
		case errors.Is(err, auth.ErrExpiredSession):
			return nil, ErrExpiredSession
		case errors.Is(err, auth.ErrInternal):
			return nil, ErrInternal
		}
		return nil, ErrUnknown
	}

	return &protocol.ValidateSessionRes{}, nil
}
