package controllers

import (
	"context"
	pro "module/netxd_customer"
	interfaces "module/interface"
	"module/models"
	//"module/netxd_dal/netxd_dal_interface"
	//"module/netxd_dal/netxd_dal_interface"
	//"module/netxd_dal/netxd_dal_models"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
	/*This defines a struct named RPCServer. It embeds the pro.UnimplementedCustomerServiceServer type, which is likely a generated type from your gRPC service definition.
	 This implies that RPCServer is intended to be used as a server for the gRPC service. */
}

var (
	CustomerService interfaces.Icustomer
	/* This declares a global variable named CustomerService, which is of type interfaces.Icustomer. It appears to be used to interact with some external service or component, possibly related to customer operations.
go
*/
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerRequest) (*pro.CustomerResponse, error) {
	dbCustomer := &models.CustomerRequest{BankID: req.BankID,Customer_Name:req.Customer_Name,Customer_ID:req.Customer_ID,Balance: req.Balance}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			Balance: result.Balance,
			
		}
		
		return responseCustomer, nil
	}
}

func (s *RPCServer) UpdateCustomer(ctx context.Context, req *pro.TransactionRequest) (*pro.TransactionResponse, error) {
	
	
	result,err:=CustomerService.UpdateCustomer(req.From_ID,req.TO_ID,req.Amount)
	if err != nil {
	 return nil, err
 }
 responseCustomer := &pro.TransactionResponse{
	 From_ID: result.From_ID,
	 
 }
 
 return responseCustomer, nil
//panic("unimple")
 }