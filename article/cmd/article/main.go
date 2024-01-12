package main

import (
	"google.golang.org/grpc"
	"net"
	"strconv"
	"wzh/article/config"
	"wzh/article/infra"
	"wzh/article/internal/dao"
	"wzh/article/internal/service"
	"wzh/article/utils/kafka"
	"wzh/pkg/pb"
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
