package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	github.com/inzmamulhassan10p/gohello "github.com/inzmamulhassan-10p/gohello/proto"
)

type Github.Com/Inzmamulhassan10p/Gohello struct{}

// Return a new handler
func New() *Github.Com/Inzmamulhassan10p/Gohello {
	return &Github.Com/Inzmamulhassan10p/Gohello{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Github.Com/Inzmamulhassan10p/Gohello) Call(ctx context.Context, req *github.com/inzmamulhassan10p/gohello.Request, rsp *github.com/inzmamulhassan10p/gohello.Response) error {
	log.Info("Received Github.Com/Inzmamulhassan10p/Gohello.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Github.Com/Inzmamulhassan10p/Gohello) Stream(ctx context.Context, req *github.com/inzmamulhassan10p/gohello.StreamingRequest, stream github.com/inzmamulhassan10p/gohello.Github.Com/Inzmamulhassan10p/Gohello_StreamStream) error {
	log.Infof("Received Github.Com/Inzmamulhassan10p/Gohello.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&github.com/inzmamulhassan10p/gohello.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Github.Com/Inzmamulhassan10p/Gohello) PingPong(ctx context.Context, stream github.com/inzmamulhassan10p/gohello.Github.Com/Inzmamulhassan10p/Gohello_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&github.com/inzmamulhassan10p/gohello.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
