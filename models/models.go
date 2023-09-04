package models

import (
	//"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerRequest struct {
	BankID int32  `json:"bankid" bson:"bankid"`
	Customer_Name string `json:"customer_name" bson:"customer_name"`
	Customer_ID int32  `json:"customer_id" bson:"customer_id"`
	Balance int32  `json:"balance" bson:"balance"`
}

type TransactionRequest struct{
	Transaction_ID int32 `json:"transaction_id" bson:"transaction_id"`
	From_ID int32 `json:"from_id" bson:"from_id"`
	To_ID int32 `json:"to_id" bson:"to_id"`
	Amount int32 `json:"amount" bson:"amount"`
	
}

