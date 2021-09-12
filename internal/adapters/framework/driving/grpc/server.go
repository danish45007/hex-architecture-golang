package rpc

import (
	"log"
	"net"

	"github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driving/grpc/pb"
	"github.com/danish45007/hex-architecture-golang/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

//starts the gRPc server
func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000 %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to server gRPC over port 9000 %v", err)
	}
}
