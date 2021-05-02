package main

import (
	"log"

	"github.com/app/routes"
	clientHanlder "github.com/client/http"
	"github.com/client/service"
	"github.com/config"
	"github.com/internal/connection"
	"github.com/model"

	grpcb "github.com/protopb"

	"google.golang.org/grpc"
)

func grpcClient(config config.Config) *grpc.ClientConn { //create grpc client
	conn, err := grpc.Dial("localhost:"+config.GrpcPort, grpc.WithInsecure())
	if err != nil {
		log.Panic("failed to connect: %v", err)
	}

	return conn
}

func main() {

	config := config.Load() //load config file

	// Database
	connection.Connect(config.DatabaseDialect, config.DatabaseUrl, true) //create database connection

	model.MigrateDB(connection.Get()) //migrate db

	client := grpcClient(config) // create grpc client
	defer client.Close()

	c := grpcb.NewSumClient(client)

	//create clinet servic
	svc := service.NewClientService(c, &model.User{})
	clientService := clientHanlder.NewClientHandler(svc)

	//attach all routes
	h := routes.AttachRoutes(clientService)

	//server routes
	routes.ServeHTTP(h, config.Port)

}
