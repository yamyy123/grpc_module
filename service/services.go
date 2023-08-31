package service

import (
	"context"

	//"time"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	 "module/models"

	"go.mongodb.org/mongo-driver/mongo"
	"module/interface"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.Icustomer {
	return &CustomerService{collection, ctx}
}

func (p *CustomerService) CreateCustomer(user *models.CustomerRequest) (*models.CustomerResponse, error) {
	// id :=primitive.NewObjectID()
	// user.Customer_Id=id.Hex()
	// user.FirstName ="mohammed"
	// user.LastName = "yameen"
	// user.BankId =1172
	// user.Balance =0
	// user.CreatedAt ="06-03-2003"
	// user.UpdatedAt =user.CreatedAt
	// user.IsActive =true

	_, err := p.CustomerCollection.InsertOne(p.ctx, user)
	if err != nil {
		return nil, err
	}

	// Construct the response with customer ID and CreatedAt
	response := &models.CustomerResponse{
		Customer_Id: user.Customer_Id,
		CreatedAt:   user.CreatedAt,
	}

	return response, nil
}
