package main

import (
	controllers "Netxd_project/Netxd_Customer_Server/Netxd-Customer_Controller"
	config "Netxd_project/Netxd_Customer_config/Netxd_Customer_Config"
	constants "Netxd_project/Netxd_Customer_config/Netxd_Customer_Constants"
	Services "Netxd_project/Nexd_Customer_Dal/Netxd_Customer_Dal_Service"

	pro "Netxd_project/netxd_proto"
	"context"
	"fmt"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabse(client *mongo.Client) {
	customerCollection := config.GetCollection(client, constants.DatabaseName, "customer")
	controllers.CustomerService = Services.InitCustomer(customerCollection, context.Background())
}
func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabse(mongoclient)
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
