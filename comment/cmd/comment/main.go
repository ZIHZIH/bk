package main

import (
	"google.golang.org/grpc"
	"net"
	"strconv"
	"wzh/comment/config"
	"wzh/comment/infra"
	"wzh/comment/internal/dao"
	"wzh/comment/internal/service"
	"wzh/pkg/pb"
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
