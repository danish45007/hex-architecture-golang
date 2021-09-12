package ports

import (
	"context"

	"github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driving/grpc/pb"
)

type GrpcPort interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetDivison(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
}
