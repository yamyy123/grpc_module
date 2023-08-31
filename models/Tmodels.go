package models


type TransactionRequest struct{
	Transaction_ID int64    `json:"transaction_id" bson:"transaction_id"`
	From_ID int64            `json:"fromid" bson:"fromid"`
	To_ID int64              `json:"toid" bson:"toid"`
	Amount int64             `json:"amount" bson:"amount"`
	Timestamp string         `json:"timestamp" bson:"timestamp"`
}

type TransactionResponse struct{
   Message string            `json:"customer_id" bson:"customer_id"`
}