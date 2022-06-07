package main

import (
	"downloadSong/handler"
	pb "downloadSong/proto"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	service = "downloadsong"
	version = "latest"
)

func main() {
	// 将服务注册到consul的配置信息
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})

	// 创建服务
	srv := micro.NewService(
		micro.Address("127.0.0.1:8002"), // 提供该服务的服务端地址
		micro.Name(service),             // 服务名
		micro.Version(version),          // go-micro版本
		micro.Registry(consulReg),       // 将服务注册到consul上
	)

	// 给该服务的服务端方注册回调函数
	pb.RegisterDownloadSongHandler(srv.Server(), new(handler.DownloadSong))

	// 运行服务，开始监听
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
