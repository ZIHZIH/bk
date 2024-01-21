package main

import (
	"bk/comment/api/pb"
	"bk/comment/config"
	"bk/comment/infra"
	"bk/comment/internal/dao"
	"bk/comment/internal/service"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func main() {
	infra.Init()

	commentDao := dao.NewCommentDao(infra.Logger)
	s := grpc.NewServer()
	pb.RegisterCommentServiceServer(s, service.NewCommentService(commentDao))
	lis, err := net.Listen("tcp", config.Config.AppConfig.Host+":"+strconv.Itoa(config.Config.AppConfig.Port))
	if err != nil {
		panic(err)
	}
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
