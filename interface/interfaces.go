package interfaces

import "module/models"

type Icustomer interface{
	CreateCustomer(user *models.CustomerRequest) (*models.CustomerRequest, error)
	UpdateCustomer(fromid int32, toid int32, amount int32)(*models.TransactionRequest,error)
}