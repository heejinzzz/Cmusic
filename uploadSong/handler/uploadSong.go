package handler

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"time"

	log "go-micro.dev/v4/logger"

	pb "uploadSong/proto"
)

type UploadSong struct{}

func (e *UploadSong) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received UploadSong.Call request: %v", req.Name)

	fileName := "../songs/" + req.Name
	_, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	err = ioutil.WriteFile(fileName, req.SongFile, 0777)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	return nil
}

func (e *UploadSong) ClientStream(ctx context.Context, stream pb.UploadSong_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *UploadSong) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.UploadSong_ServerStreamStream) error {
	log.Infof("Received UploadSong.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *UploadSong) BidiStream(ctx context.Context, stream pb.UploadSong_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
