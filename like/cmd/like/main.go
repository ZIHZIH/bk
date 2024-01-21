package main

import (
	"bk/like/api/pb"
	"bk/like/config"
	"bk/like/infra"
	"bk/like/internal/dao"
	"bk/like/internal/service"
	"google.golang.org/grpc"
	"net"
	"strconv"
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
