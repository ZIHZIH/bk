package utils

import (
	"google.golang.org/grpc"
	"wzh/pkg/pb"
)

var UserServiceClient pb.UserServiceClient
var CommentServiceClient pb.CommentServiceClient
var ArticleServiceClient pb.ArticleServiceClient
var LikeServiceClient pb.LikeServiceClient

func NewUserServiceClient() {
	conn, err := grpc.Dial("127.0.0.1:15034", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	UserServiceClient = pb.NewUserServiceClient(conn)
}

func NewCommentServiceClient() {
	conn, err := grpc.Dial("127.0.0.1:15032", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	CommentServiceClient = pb.NewCommentServiceClient(conn)
}

func NewArticleServiceClient() {
	conn, err := grpc.Dial("127.0.0.1:15031", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	ArticleServiceClient = pb.NewArticleServiceClient(conn)
}

func NewLikeServiceClient() {
	conn, err := grpc.Dial("127.0.0.1:15033", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	LikeServiceClient = pb.NewLikeServiceClient(conn)
}
