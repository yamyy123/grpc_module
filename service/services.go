package service

import (
	"context"
	"fmt"
	"log"

	//"log"

	//"time"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"module/models"

	"module/interface"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	client                *mongo.Client
	CustomerCollection    *mongo.Collection
	TransactionCollection *mongo.Collection
	ctx                   context.Context
}

// UpdateCustomer implements interfaces.Icustomer.
// func (*CustomerService) UpdateCustomer(fromid int32, toid int32, amount int32) (*models.CustomerRequest, error) {
// 	panic("unimplemented")
// }

func InitCustomerService(client *mongo.Client, collection1 *mongo.Collection, collection2 *mongo.Collection, ctx context.Context) interfaces.Icustomer {

	return &CustomerService{client, collection1, collection2, ctx}
}

func (p *CustomerService) CreateCustomer(user *models.CustomerRequest) (*models.CustomerRequest, error) {
	res, err := p.CustomerCollection.InsertOne(p.ctx, &user)

	if err != nil {

		return nil, err
	}

	var insertedCustomer models.CustomerRequest

	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {

		return nil, fmt.Errorf("InsertOne did not return an ObjectID")
	}

	// Use the InsertedID to fetch the inserted document.
	filter := bson.M{"Customer_ID": insertedID}
	err = p.CustomerCollection.FindOne(p.ctx, filter).Decode(&insertedCustomer)
	if err != nil {
		// Handle the error
		return nil, err
	}

	// Return the insertedCustomer as a pointer.
	return &insertedCustomer, nil
}

func (p*CustomerService) UpdateCustomer(fromid int32, toid int32, amount int32)(*models.TransactionRequest,error){
	var transactionToInsert	models.TransactionRequest
	//Create a session
		  session, err := p.client.StartSession()
		  if err != nil {
			  log.Fatal(err)
		  }
		  defer session.EndSession(context.Background())
	  
		  //Start a transaction
		  _, err = session.WithTransaction(context.Background(), func(ctx mongo.SessionContext) (interface{}, error) {
	  
			  //two update queries(dec, inc)
			  //deducting from
			  filter1 := bson.M{"customer_id": fromid}
			  update1 := bson.M{"$inc": bson.M{"balance": -(amount)}}
			  _, err1 := p.CustomerCollection.UpdateOne(context.Background(), filter1, update1)
			  if err1 != nil {
				  fmt.Println("Transaction Failed")
				  return nil, err1
			  }
	  
			  //incrementing to
			  filter2 := bson.M{"customer_id": toid}
			  update2:= bson.M{"$inc": bson.M{"balance": amount}}
			  _, err2 := p.CustomerCollection.UpdateOne(context.Background(), filter2,update2)
			  if err2 != nil {
				  fmt.Println("Transaction Failed")
				  return nil, err2
			  }
	  
			  //inserting the transaction
			  transactionToInsert = models.TransactionRequest{
				  Transaction_ID: 1,
				  From_ID:   fromid,
				  To_ID:     toid,
				  Amount:        amount,
			  }
			  _, err := p.TransactionCollection.InsertOne(context.Background(), transactionToInsert)
			  if err != nil {
				  fmt.Println("Transaction not inserted")
				  return "Transaction not inserted", err
			  }
			  return "Transaction inserted", nil
		  })
		  return &transactionToInsert,nil
		 // panic("unimple")
}	
