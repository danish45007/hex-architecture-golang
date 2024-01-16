package main

import (
	"log"
	"os"

	// application
	"github.com/danish45007/hex-architecture-golang/internal/application/api"
	"github.com/danish45007/hex-architecture-golang/internal/application/core/arithmetic"

	// right adapter
	"github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driven/db"

	// left adapters
	gRPC "github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driving/grpc"
	"github.com/danish45007/hex-architecture-golang/internal/adapters/framework/driving/rest"
)

func main() {
	var err error
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DS_NAME")

	dbAdapter, err := db.NewAdapter(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Failed to initiate database %v", err)
	}

	defer dbAdapter.CloseDbConn()

	// core
	core := arithmetic.New()

	// NOTE: The application's right side port for driven
	// adapters, in this case, a db adapter.
	// Therefore the type for the dbAdapter parameter
	// that is to be injected into the NewApplication will
	// be of type DbPort
	applicationAPI := api.NewApplication(dbAdapter, core)

	// NOTE: We use dependency injection to give the grpc
	// adapter access to the application, therefore
	// the location of the port is inverted. That is
	// the grpc adapter accesses the hexagon's driving port at the
	// application boundary via dependency injection,
	// therefore the type for the applicationAPI parameter
	// that is to be injected into the gRPC adapter will
	// be of type APIPort which is our hexagons left side
	// port for driving adapters

	// rest based implementation of the APIPort for the application
	restAdapter := rest.NewAdapter(applicationAPI)
	restAdapter.Run()

	// gRPC based implementation of the APIPort for the application
	gRPCAdapter := gRPC.NewAdapter(applicationAPI)
	gRPCAdapter.Run()

}
