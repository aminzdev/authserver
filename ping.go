package authserver

import (
	"context"
	"github.com/aminzdev/authserver/protocol"
)

func (a *authServer) Ping(_ context.Context, _ *protocol.PingReq) (*protocol.PingRes, error) {
	return &protocol.PingRes{}, nil
}
