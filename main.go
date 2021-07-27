package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tao-yi/go-grpc-demo/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUserInfo(context.Context, *pb.UserID) (*pb.UserInfo, error) {
	return &pb.UserInfo{
		Id:     1,
		Name:   "Hell",
		Age:    25,
		Gender: pb.UserInfo_MALE}, nil
}

func (s *server) GetArticles(context.Context, *pb.UserID) (*pb.Articles, error) {
	return &pb.Articles{
		Articles: []*pb.Articles_Article{
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
	userServiceServer := &server{}
	pb.RegisterUserServiceServer(s, userServiceServer)
	reflection.Register(s)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	gwmux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterUserServiceHandlerServer(ctx, gwmux, userServiceServer)
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
