package main

import (
	"context"
	"fmt"
	
	"log"

	pb "module/netxd_customer" 
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	

	response, err := client.CreateCustomer(context.Background(), &pb.CustomerRequest{Customer_Id: "2",CreatedAt:"28-06-2024"})
if err != nil {
log.Fatalf("Failed to call CreateCustomer: %v", err)
}

fmt.Printf("Response:\nCustomer_id:%v\nCreatedAt:%v\n", response.Customer_Id,response.CreatedAt)


}