package Cmusic

import (
	pb "Cmusic/uploadSongProto"
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	client2 "go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"io/ioutil"
	"strings"
	"time"
)

// UploadSong 向C-Music上传歌曲
func UploadSong(fileName string) {
	// 指定服务所注册到的consul的地址，之后客户端会从该consul中取得提供其所需服务的服务端的地址
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{consulAddress}
	})

	// 创建服务
	service := micro.NewService(
		micro.Name("uploadsong"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)

	// 创建该服务的客户端方
	client := pb.NewUploadSongService("uploadsong", service.Client())
	
	// 设置连接超时时间
	opts := func(o *client2.CallOptions) {
		o.RequestTimeout = 5 * time.Minute
		o.DialTimeout = 5 * time.Minute
	}

	// 读指定文件，将文件名和读到的数据传给服务端
	songData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fileNameSplit := strings.Split(fileName, "/")
	songName := fileNameSplit[len(fileNameSplit)-1]

	req := pb.CallRequest{
		Name:     songName,
		SongFile: songData,
	}
	_, err = client.Call(context.Background(), &req, opts)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("{C-Music}\n歌曲：" + fileName + " 上传成功！")
}
