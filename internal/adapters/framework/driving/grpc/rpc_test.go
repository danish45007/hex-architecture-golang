package rpc

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/danish45007/hex-architecture-golang/internal/adapters/app/api"
	"github.com/danish45007/hex-architecture-golang/internal/adapters/core/arithmetic"
	"github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driven/db"
	"github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driving/grpc/pb"
	"github.com/danish45007/hex-architecture-golang/internal/ports"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const buffSize = 1024 * 1024

var ls *bufconn.Listener

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func init() {
	var err error
	ls = bufconn.Listen(buffSize)
	grpcServer := grpc.NewServer()
	//core(business-logic)
	var core ports.ArithmeticPort
	//application
	var appAdapter ports.APIPort
	//framework(database)
	var dbAdapter ports.DBPort

	var grpcAdapter ports.GrpcPort

	dbDriver := goDotEnvVariable("DBDRIVER")
	dbSource := goDotEnvVariable("DBSOURCENAME")

	dbAdapter, err = db.NewAdapter(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Failed to initiate database %v", err)
	}

	defer dbAdapter.CloseDbConn()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(core, dbAdapter)
	grpcAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, grpcAdapter)

	go func() {
		if err := grpcServer.Serve(ls); err != nil {
			log.Fatalf("Failed to start the test server %v", err)
		}
	}()
}

func bufDailer(context.Context, string) (net.Conn, error) {
	return ls.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDailer), grpc.WithInsecure)
	if err != nil {
		log.Fatalf("Failed to dial to bufnet :%v", err)
	}

	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()
	clinet := pb.NewArithmeticServiceClient(conn)
	params := pb.OperationParameters{
		A: 1,
		B: 1,
	}
	ans, err := clinet.GetAddition(ctx, &params)
	if err != nil {
		t.Fatalf("Expected %v but got %v", nil, err)
	}
	require.Equal(t, ans, int32(2))
}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()
	clinet := pb.NewArithmeticServiceClient(conn)
	params := pb.OperationParameters{
		A: 2,
		B: 1,
	}
	ans, err := clinet.GetSubtraction(ctx, &params)
	if err != nil {
		t.Fatalf("Expected %v but got %v", nil, err)
	}
	require.Equal(t, ans, int32(1))
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()
	clinet := pb.NewArithmeticServiceClient(conn)
	params := pb.OperationParameters{
		A: 2,
		B: 3,
	}
	ans, err := clinet.GetMultiplication(ctx, &params)
	if err != nil {
		t.Fatalf("Expected %v but got %v", nil, err)
	}
	require.Equal(t, ans, int32(6))
}

func TestGetDivison(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()
	clinet := pb.NewArithmeticServiceClient(conn)
	params := pb.OperationParameters{
		A: 10,
		B: 5,
	}
	ans, err := clinet.GetDivison(ctx, &params)
	if err != nil {
		t.Fatalf("Expected %v but got %v", nil, err)
	}
	require.Equal(t, ans, int32(2))
}
