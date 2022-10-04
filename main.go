package main

import (
	"github.com/inzmamulhassan-10p/gohello/handler"
	pb "github.com/inzmamulhassan-10p/gohello/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("github.com/inzmamulhassan-10p/gohello"),
	)

	// Register handler
	pb.RegisterGithub.Com/Inzmamulhassan10p/GohelloHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
