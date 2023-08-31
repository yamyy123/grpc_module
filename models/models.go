package models

import (
	//"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerRequest struct {
	Customer_Id string `json:"customer_id" bson:"customer_id"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName string `json:"lastname" bson:"lastname"`
	BankId int32 `json:"bankid" bson:"bankid"`
	Balance int64 `json:"balance" bson:"balance"`
	CreatedAt string `json:"createdat" bson:"createdat"`
	UpdatedAt string `json:"updatedat" bson:"updatedat"`
	IsActive bool `json:"isactive" bson:"isactive"`
	}

type CustomerResponse struct{
	Customer_Id string `json:"customer_id" bson:"customer_id"`
	CreatedAt string `json:"createdat" bson:"createdat"`
}