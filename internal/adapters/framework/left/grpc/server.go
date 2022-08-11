package rpc

import (
	"hex-tutorial/internal/adapters/framework/left/grpc/pb"
	"hex-tutorial/internal/ports"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (grpca Adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticsServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticsServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc server over port 9000: %v", err)
	}
}
