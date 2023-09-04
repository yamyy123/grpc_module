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
	CustomerCollection := config.GetCollection(client, "bankdb", "customer")
	TransactionCollection:=config.GetCollection(client,"bankdb","transactions")
controllers.CustomerService =service.InitCustomerService(client,CustomerCollection, TransactionCollection,context.Background())
//this line setsup  the services needed for the operations
/*t initializes the controllers.CustomerService variable by calling the service.InitCustomerService function, passing the MongoDB client and collections as parameters.
 This is probably where your application's business logic for customer services is set up.*/

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
	//It creates a new gRPC server instance s using grpc.NewServer().

	pro.RegisterCustomerServiceServer(s, &controllers.RPCServer{})
     //this line connect the  comntroller and server 
	 //It registers the controllers.RPCServer{} as the gRPC server implementation for the pro package's CustomerServiceServer.
	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
	//Finally, it starts serving gRPC requests using s.Serve(lis) and logs any errors if they occur during the serving process.
}