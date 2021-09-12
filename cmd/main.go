package main

import (
	"log"
	"os"

	"github.com/danish45007/hex-architecture-golang/internal/adapters/app/api"
	"github.com/danish45007/hex-architecture-golang/internal/adapters/core/arithmetic"
	"github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driven/db"
	gRPC "github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driving/grpc"
	"github.com/danish45007/hex-architecture-golang/internal/ports"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	var err error
	//ports

	//core(business-logic)
	var core ports.ArithmeticPort
	//application
	var appAdapter ports.APIPort
	//framework(database)
	var dbAdapter ports.DBPort
	//framework(gRPC)
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
	grpcAdapter = gRPC.NewAdapter(appAdapter)
	grpcAdapter.Run()

}
