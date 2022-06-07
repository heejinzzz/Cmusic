package handler

import (
	"context"
	"io"
	"io/ioutil"
	"time"

	log "go-micro.dev/v4/logger"

	pb "downloadSong/proto"
)

type DownloadSong struct{}

func (e *DownloadSong) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received DownloadSong.Call request: %v", req)

	fileName := "../songs/" + req.Name
	songData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Error(err)
		return err
	}

	rsp.Name = req.Name
	rsp.SongData = songData
	return nil
}

func (e *DownloadSong) ClientStream(ctx context.Context, stream pb.DownloadSong_ClientStreamStream) error {
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

func (e *DownloadSong) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.DownloadSong_ServerStreamStream) error {
	log.Infof("Received DownloadSong.ServerStream request: %v", req)
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

func (e *DownloadSong) BidiStream(ctx context.Context, stream pb.DownloadSong_BidiStreamStream) error {
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
