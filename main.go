package main

import (
	"context"
	"log"
	"net"

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

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
