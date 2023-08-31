package main

import (
	"context"
	"fmt"
	"module/config"
	"module/constants"
	pro "module/netxd_customer"
	controllers "module/controller/controllers"
	//"module/netxd_dal/netxd_dal_service"
	"net"
     "module/service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	CustomerCollection := config.GetCollection(client, "netxdb", "customer")
	controllers.CustomerService = service.InitCustomerService(CustomerCollection, context.Background())

}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &controllers.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}