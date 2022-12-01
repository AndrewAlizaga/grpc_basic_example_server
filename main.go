package main

import (
	"log"
	"net"

	accountapiv1 "github.com/AndrewAlizaga/grpc_basic_example_proto/pkg/proto/v1/services/account"
	accountsvc1 "github.com/AndrewAlizaga/grpc_basic_example_server/internal/account/v1"
	"google.golang.org/grpc"
)

func main() {

	//create listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error using tpc port 8080: ", err.Error())
	}

	// create gRPC server object
	grpcServer := grpc.NewServer()

	//register out grpc methods into the server
	accountapiv1.RegisterAccountServiceServer(grpcServer, &accountsvc1.AccountServer{})

	log.Println("Server running on port 8080")

	// start the server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
