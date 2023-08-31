package controllers

import (
	interfaces "Netxd_project/Nexd_Customer_Dal/Netxd_Customer_Interfaces"
	models "Netxd_project/Nexd_Customer_Dal/Netxd_Customer_models"
	pro "Netxd_project/netxd_proto"
	"context"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	dbCustomer := &models.Customer{Customer_Id: req.CustomerId}

	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerId: result.CustomerId,
		}
		return responseCustomer, nil
	}
}

func (s *RPCServer) Transaction(ctx context.Context, req *pro.Transfer) (*pro.TransferResponse, error) {

	response := &models.Transfer{
		FromID: req.FromId,
		ToID:   req.ToId,
		Amount: req.Amount,
	}
	result, err := CustomerService.Transaction(response)
	
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.TransferResponse{
			FromId: result.FromID,
			ToId: result.ToID,
		}
		return responseCustomer, nil



}}
