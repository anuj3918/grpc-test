package main

import (
	"fmt"
	"log"
	"net"

	pb "grpc-test/messages/pb"
	services "grpc-test/services"

	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// secure := flag.Bool("secure", false, "Flag to create a server with certificates")
	// flag.Parse()

	secure := false

	const port = ":5001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var server *grpc.Server
	if secure == true {
		// Create the TLS credentials
		creds, err := credentials.NewServerTLSFromFile("./cert/localhost.crt", "./cert/localhost.key")
		if err != nil {
			log.Fatalf("could not load TLS keys: %s", err)
		}
		// Create an array of gRPC options with the credentials
		opts := []grpc.ServerOption{grpc.Creds(creds)}

		server = grpc.NewServer(opts...)
	} else {
		server = grpc.NewServer()
	}

	pb.RegisterUserServiceServer(server, &services.UserService{})

	fmt.Println("gRPC server now listening on PORT " + port)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
