package handler

import (
	"context"
	"io"
	"io/ioutil"
	"time"

	log "go-micro.dev/v4/logger"

	pb "songList/proto"
)

type SongList struct{}

func (e *SongList) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received SongList.Call request: %v", req)

	files, err := ioutil.ReadDir("../songs")
	if err != nil {
		log.Error(err)
		return err
	}

	for _, i := range files {
		rsp.Songs = append(rsp.Songs, i.Name())
	}
	return nil
}

func (e *SongList) ClientStream(ctx context.Context, stream pb.SongList_ClientStreamStream) error {
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

func (e *SongList) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.SongList_ServerStreamStream) error {
	log.Infof("Received SongList.ServerStream request: %v", req)
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

func (e *SongList) BidiStream(ctx context.Context, stream pb.SongList_BidiStreamStream) error {
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
