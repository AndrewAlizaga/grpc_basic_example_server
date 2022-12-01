package account

import (
	"context"
	"log"

	accountapiv1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/services/account"
)

// AccountServer ...
type AccountServer struct{}


func (AccountServer) LoginService(ctx context.Context, in *accountapiv1.LoginRequest) (*accountapiv1.LoginResponse, error) {
	log.Println("INVOKE - LoginService")
	return nil, nil
}

func (AccountServer) SignUpService(ctx context.Context, in *accountapiv1.SignUpRequest) (*accountapiv1.SignUpResponse, error) {
	log.Println("INVOKE - SignUpService")
	return nil, nil
}