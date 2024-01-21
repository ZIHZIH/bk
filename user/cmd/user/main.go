package main

import (
	"bk/user/api/pb"
	"bk/user/config"
	"bk/user/infra"
	"bk/user/internal/dao"
	"bk/user/internal/service"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func main() {
	infra.Init()

	userDao := dao.NewUserDao(infra.Logger)
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, service.NewUserService(userDao))
	lis, err := net.Listen("tcp", config.Config.AppConfig.Host+":"+strconv.Itoa(config.Config.AppConfig.Port))
	if err != nil {
		panic(err)
	}
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
