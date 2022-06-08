package Cmusic

import (
	pb "Cmusic/downloadSongProto"
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"io/ioutil"
	"os"
)

// DownloadSong 从C-Music下载歌曲保存到本地
func DownloadSong(songName string, savePath string) {
	// 指定服务所注册到的consul的地址，之后客户端会从该consul中取得提供其所需服务的服务端的地址
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{consulAddress}
	})

	// 创建服务
	service := micro.NewService(
		micro.Name("downloadsong"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)

	// 创建该服务的客户端方
	client := pb.NewDownloadSongService("downloadsong", service.Client())

	// 客户端调用服务，获取指定歌曲的数据，将其保存为歌曲文件
	rep, err := client.Call(context.Background(), &pb.CallRequest{Name: songName})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 保存文件
	fileName := savePath + "/" + rep.Name
	_, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = ioutil.WriteFile(fileName, rep.SongData, 0777)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("{C-Music}\n歌曲：" + rep.Name + " 下载成功\n已保存至：" + fileName)
}
