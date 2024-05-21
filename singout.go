package authserver

import (
	"context"
	"errors"
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver/protocol"
)

func (a *authServer) SignOut(_ context.Context, req *protocol.SignOutReq) (*protocol.SignOutRes, error) {
	err := a.auth.SignOut(req.User.Name, req.User.Code)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrWrongCredentials):
			return nil, ErrUnauthenticated
		case errors.Is(err, auth.ErrInternal):
			return nil, ErrInternal
		}
		return nil, ErrUnknown
	}

	return &protocol.SignOutRes{}, nil
}
