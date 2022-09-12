package gapi

import (
	"fmt"

	db "github.com/copter01252/simplebank/db/sqlc"
	"github.com/copter01252/simplebank/pb"
	"github.com/copter01252/simplebank/token"
	"github.com/copter01252/simplebank/util"
)

// Server serves gRPC request for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServe create a new gRPC server
func NewServe(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot creat token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
