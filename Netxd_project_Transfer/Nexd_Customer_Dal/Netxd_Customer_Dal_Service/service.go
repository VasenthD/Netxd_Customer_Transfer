package Services

import (
	interfaces "Netxd_project/Nexd_Customer_Dal/Netxd_Customer_Interfaces"
	models "Netxd_project/Nexd_Customer_Dal/Netxd_Customer_models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomer(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (c *CustomerService) CreateCustomer(user *models.Customer) (*models.CustomerResponse, error) {

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.IsActive = true

	_, err := c.CustomerCollection.InsertOne(c.ctx, &user)
	if err != nil {
		return nil, err
	}
	response := &models.CustomerResponse{
		CustomerId: user.Customer_Id,
		CreatedAt:  user.CreatedAt.UTC(),
	}

	return response, nil
}
func (c *CustomerService) Transaction(user *models.Transfer) (*models.TransferResponse, error) {

	fmt.Println("im running TRANSACTION **********************")

	_, err := c.CustomerCollection.UpdateOne(context.Background(),
		bson.M{"customer_id": user.FromID},
		bson.M{"$inc": bson.M{"balance": (5000)}})
	if err != nil {
		fmt.Println("Transaction Failed1", err)
		return nil, err
	}
	_, err2 := c.CustomerCollection.UpdateOne(context.Background(), bson.M{"customer_id": user.ToID}, bson.M{"$inc": bson.M{"balance": 4000}})

	if err2 != nil {
		fmt.Println("Transaction Failed")
		return nil, err2
	}
	fmt.Println("im running done TRANSACTION **********************")
	return nil, nil
}
