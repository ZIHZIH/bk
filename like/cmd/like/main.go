package main

import (
	"google.golang.org/grpc"
	"net"
	"strconv"
	"wzh/like/config"
	"wzh/like/infra"
	"wzh/like/internal/dao"
	"wzh/like/internal/service"
	"wzh/pkg/pb"
)

func main() {
	infra.Init()

	likeDao := dao.NewLikeDao(infra.Logger)
	s := grpc.NewServer()
	pb.RegisterLikeServiceServer(s, service.NewLikeService(likeDao))
	lis, err := net.Listen("tcp", config.Config.AppConfig.Host+":"+strconv.Itoa(config.Config.AppConfig.Port))
	if err != nil {
		panic(err)
	}
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
