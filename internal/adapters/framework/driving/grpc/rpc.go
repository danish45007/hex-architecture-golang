package rpc

import (
	"context"
	"log"

	"github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driving/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")

	}

	answer, err := grpca.api.GetAddition(req.A, req.B)
	if err != nil {
		log.Fatalf("Failed to perform operation %v", err)
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")

	}

	answer, err := grpca.api.GetSubtraction(req.A, req.B)
	if err != nil {
		log.Fatalf("Failed to perform operation %v", err)
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")

	}

	answer, err := grpca.api.GetMultiplication(req.A, req.B)
	if err != nil {
		log.Fatalf("Failed to perform operation %v", err)
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca Adapter) GetDivison(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}
	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")

	}

	answer, err := grpca.api.GetDivision(req.A, req.B)
	if err != nil {
		log.Fatalf("Failed to perform operation %v", err)
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}
