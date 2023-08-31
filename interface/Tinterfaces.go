package interfaces

import "module/models"

type ITransaction interface{
	//CreateCustomer(customer *models.CustomerRequest)(*models.CustomerResponse,error)
	CreateTransaction(transaction *models.TransactionRequest)(*models.CustomerResponse,error)
}