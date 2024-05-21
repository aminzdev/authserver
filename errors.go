package authserver

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknown           = status.Error(codes.Unknown, "unknown error.")
	ErrInternal          = status.Error(codes.Internal, "internal error.")
	ErrUnauthenticated   = status.Error(codes.Unauthenticated, "wrong credentials.")
	ErrSessionNotFound   = status.Error(codes.Unauthenticated, "session not found.")
	ErrExpiredSession    = status.Error(codes.Unauthenticated, "session is expired.")
	ErrUserAlreadyExists = status.Error(codes.AlreadyExists, "user already exists.")
	ErrInvalidUserName   = status.Error(codes.InvalidArgument, "username is invalid.")
)
