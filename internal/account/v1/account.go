package account

import (
	"context"
	"errors"
	"log"

	accountasvc1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/account"
	accountapiv1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/services/account"
	uuid "github.com/satori/go.uuid"
)

// AccountsServer ...
type AccountsServer struct{}

var Accounts []*accountasvc1.Account = []*accountasvc1.Account{
	{
		Id:       uuid.NewV4().String(),
		Name:     "Jos",
		Email:    "jos1@gmail.com",
		Password: "qwerty",
	},
	{
		Id:       uuid.NewV4().String(),
		Name:     "Max",
		Email:    "max1@gmail.com",
		Password: "qwerty",
	},
	{
		Id:       uuid.NewV4().String(),
		Name:     "Day",
		Email:    "day1@gmail.com",
		Password: "qwerty",
	},
}

func (AccountsServer) LoginService(ctx context.Context, in *accountapiv1.LoginRequest) (*accountapiv1.LoginResponse, error) {
	log.Println("INVOKE - LoginService")

	var response accountapiv1.LoginResponse
	var err error

	for _, account := range Accounts {
		if account.GetEmail() == in.GetEmail() {
			if account.GetPassword() == in.GetPassword() {
				response.Account = account
				response.Jwt = uuid.NewV4().String()
			}
		}
	}

	if response.GetAccount() == nil {
		response.Error = "bad email / password"
		err = errors.New(response.Error)
	}

	return &response, err
}

func (AccountsServer) SignUpService(ctx context.Context, in *accountapiv1.SignUpRequest) (*accountapiv1.SignUpResponse, error) {
	log.Println("INVOKE - SignUpService")

	var response accountapiv1.SignUpResponse
	var err error

	if in.GetAccount() != nil && in.GetAccount().GetEmail() != "" && in.GetAccount().GetName() != "" && in.GetAccount().GetPassword() != "" {
		Accounts = append(Accounts, in.GetAccount())
		response.Jwt = uuid.NewV4().String()
	} else {
		response.Error = "bad formed request"
		err = errors.New(response.Error)
	}

	return &response, err
}
