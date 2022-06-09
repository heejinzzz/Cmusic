package Cmusic

import (
	pb "Cmusic/songListProto"
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	client2 "go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"strconv"
	"time"
)

var consulAddress = "127.0.0.1:8500"

// GetSongList 获取C-Music中的歌曲列表并打印出来
func GetSongList() {
	// 指定服务所注册到的consul的地址，之后客户端会从该consul中取得提供其所需服务的服务端的地址
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{consulAddress}
	})

	// 创建服务
	srv := micro.NewService(
		micro.Name("songlist"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)

	// 创建该服务的客户端方
	client := pb.NewSongListService("songlist", srv.Client())
	
	// 设置连接超时时间
	opts := func(o *client2.CallOptions) {
		o.RequestTimeout = 20 * time.Second
		o.DialTimeout = 20 * time.Second
	}

	// 客户端调用服务，获取歌曲列表，打印出来
	songs, err := client.Call(context.Background(), &pb.CallRequest{}, opts)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("{ C-Music }\nSongList:")
	for i, song := range songs.Songs {
		fmt.Println("【" + strconv.Itoa(i+1) + "】" + ": " + song)
	}
}
