package models

import (
	"time"
)

type Customer struct {
	Customer_Id int32     `json:"customer_id" bson:"customer_id"`
	FirstName   string    `json:"first_name" bson:"first_name"`
	LastName    string    `json:"last_name" bson:"last_name"`
	BankId      string    `json:"bank_id" bson:"bank_id"`
	Balance     float32   `json:"balance" bson:"balance"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	IsActive    bool      `json:"is_active" bson:"is_active"`
}

type CustomerResponse struct {
	CustomerId int32     `json:"customer_id" bson:"customer_id"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}
type Transfer struct {
    FromID   int32   `json:"from_id" bson:"from_id"`
    ToID     int32   `json:"to_id" bson:"to_id"`
    Amount   float32 `json:"amount" bson:"amount"`
}

type TransferResponse struct {
    FromID          int32   `json:"from_id" bson:"from_id"`
    FromIDBalance   float32 `json:"from_id_balance" bson:"from_id_balance"`
    ToID            int32   `json:"to_id" bson:"to_id"`
    ToIDBalance     float32 `json:"to_id_balance" bson:"to_id_balance"`
}