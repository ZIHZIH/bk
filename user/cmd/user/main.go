package main

import (
	"google.golang.org/grpc"
	"net"
	"strconv"
	"wzh/pkg/pb"
	"wzh/user/config"
	"wzh/user/infra"
	"wzh/user/internal/dao"
	"wzh/user/internal/service"
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
