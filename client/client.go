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

	// response, err := client.CreateCustomer(context.Background(), &pb.CustomerRequest{
	// 	BankID:        566,
	// 	Customer_Name: "yameen",
	// 	Customer_ID:   50,
	// 	Balance:       500,
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to call SayHello: %v", err)
	// }

	// fmt.Printf("Response: %s\n", response.Balance)
	response1, err1 := client.UpdateCustomer(context.Background(), &pb.TransactionRequest{
		From_ID: 50,
		TO_ID:   60,
		Amount:  500,
	})
	if err1 != nil {
		log.Fatalf("Failed to call SayHello: %v", err1)
	}

	fmt.Printf("Response: %s\n", response1.From_ID)
 }
