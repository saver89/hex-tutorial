package main

import (
	"hex-tutorial/internal/adapters/app/api"
	"hex-tutorial/internal/adapters/core/arithmetics"
	rpc "hex-tutorial/internal/adapters/framework/left/grpc"
	"hex-tutorial/internal/adapters/framework/right/db"
	"hex-tutorial/internal/ports"
	"os"
)

func main() {
	// port
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter = db.NewAdapter(dbaseDriver, dsourceName)
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetics.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)
	gRPCAdapter = rpc.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
