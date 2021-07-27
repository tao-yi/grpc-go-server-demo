package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tao-yi/go-grpc-demo/proto/article"
	"github.com/tao-yi/go-grpc-demo/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	user.UnimplementedUserServiceServer
}

func (s *server) GetUserInfo(context.Context, *user.UserID) (*user.UserInfo, error) {
	return &user.UserInfo{
		Id:     1,
		Name:   "Hell",
		Age:    25,
		Gender: user.UserInfo_MALE}, nil
}

func (s *server) GetArticles(context.Context, *user.UserID) (*article.Articles, error) {
	return &article.Articles{
		Articles: []*article.Articles_Article{
			{
				Id:    2,
				Title: "title",
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = user.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":50052",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:50052")
	log.Fatalln(gwServer.ListenAndServe())
}
