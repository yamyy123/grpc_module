package interfaces

import "module/models"

type Icustomer interface{
	CreateCustomer(customer *models.CustomerRequest)(*models.CustomerResponse,error)
}