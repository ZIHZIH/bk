package main

import (
	"bk/article/api/pb"
	"bk/article/config"
	"bk/article/infra"
	"bk/article/internal/dao"
	"bk/article/internal/service"
	"bk/article/utils/kafka"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func main() {
	infra.Init()

	articleMongodbDao := dao.NewArticleMongodbDao(infra.Mongodb.Database("bk").Collection("article"), infra.Logger)
	s := grpc.NewServer()
	pb.RegisterArticleServiceServer(s, service.NewArticleService(articleMongodbDao, kafka.NewProducer()))

	consumer := kafka.NewConsumer(infra.Logger)
	go func() {
		consumer.Run()
	}()

	lis, err := net.Listen("tcp", config.Config.AppConfig.Host+":"+strconv.Itoa(config.Config.AppConfig.Port))
	if err != nil {
		panic(err)
	}
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
