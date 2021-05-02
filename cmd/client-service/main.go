package main

import (
	"log"

	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/app/routes"
	clientHanlder "github.com/PersonalGithubAccount/http-service-with-grpc-POC/client/http"
	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/client/service"
	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/config"
	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/internal/connection"
	"github.com/PersonalGithubAccount/http-service-with-grpc-POC/model"

	grpcb "github.com/PersonalGithubAccount/http-service-with-grpc-POC/protopb"

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
