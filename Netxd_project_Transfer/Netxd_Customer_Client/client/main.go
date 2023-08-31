package main

import (
	pb "Netxd_project/netxd_proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect :%v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	// response, err := client.CreateCustomer(context.Background(), &pb.Customer{
	// 	CustomerId: 221,
	// 	FirstName:  "Vasenth",
	// 	LastName:   "D",
	// 	BankId:     "10A",
	// 	Balance:    120000,
	// 	CreatedAt:  "2023.08.30",
	// 	UpdatedAt:  "2023.08.30",
	// 	IsActive:   true,
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to create customer :%v", err)
	// }
	fmt.Printf("Response")

	response, err := client.Transaction(context.Background(), &pb.Transfer{
			FromId: 221,
			ToId: 222,
			Amount: 30,
	})
	
	
	fmt.Printf("Response from customer : %v\n", response.FromId)
	fmt.Printf("Response from customer : %v\n", response.FromId)

	}
	




