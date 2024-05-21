package authserver

import (
	"context"
	"errors"
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver/protocol"
	"time"
)

func (a *authServer) SignIn(_ context.Context, req *protocol.SignInReq) (*protocol.SignInRes, error) {
	err := a.auth.SignIn(req.User.Name, req.User.Code, time.Hour)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrWrongCredentials):
			return nil, ErrUnauthenticated
		case errors.Is(err, auth.ErrInternal):
			return nil, ErrInternal
		}
		return nil, ErrUnknown
	}

	return &protocol.SignInRes{}, nil
}
