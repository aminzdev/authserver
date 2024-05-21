package authserver

import (
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver/protocol"
	"github.com/aminzdev/db"
	"github.com/aminzdev/server"
	"google.golang.org/grpc"
)

type authServer struct {
	protocol.UnimplementedAuthServer
	auth *auth.Auth
}

func RunAuthServer(grpcHost, grpcWebHost string, db *db.DB) {
	server.RunServer(
		&authServer{auth: auth.NewAuth(db)},
		"auth",
		grpcHost,
		grpcWebHost,
		func(registrar grpc.ServiceRegistrar, i interface{}) {
			protocol.RegisterAuthServer(registrar, i.(protocol.AuthServer))
		},
	)
}
