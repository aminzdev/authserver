package authserver

import (
	"context"
	"errors"
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver/protocol"
)

func (a *authServer) SignUp(_ context.Context, req *protocol.SignUpReq) (*protocol.SignUpRes, error) {
	err := a.auth.SignUp(req.User.Name, req.User.Code)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrInvalidUserName):
			return nil, ErrInvalidUserName
		case errors.Is(err, auth.ErrUserAlreadyExists):
			return nil, ErrUserAlreadyExists
		case errors.Is(err, auth.ErrInternal):
			return nil, ErrInternal
		}
		return nil, ErrUnknown
	}

	return &protocol.SignUpRes{}, nil
}
