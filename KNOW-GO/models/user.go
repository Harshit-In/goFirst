package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type USER struct {
	ID                       primitive.ObjectID `json:"_id" bson:"_id"`
	Member_id                string             `json:"member_id" bson:"member_id"`
	Member_name              string             `json:"member_name" bson:"member_name"`
	Sponsor_id               float64            `json:"sponsor_id" bson:"sponsor_id"`
	Sponsor_name             string             `json:"sponsor_name" bson:"sponsor_name"`
	Email                    string             `json:"email" bson:"email"`
	Hash_password            string             `json:"hash_password" bson:"hash_password"`
	ActiveStatus             int64              `json:"activeStatus" bson:"activeStatus"`
	State                    string             `json:"state" bson:"state"`
	Wallet_amount            int64              `json:"wallet_amount" bson:"wallet_amount"`
	New_wallet_amount        int64              `json:"new_wallet_amount" bson:"new_wallet_amount"`
	Test_wallet_amount       int32              `json:"test_wallet_amount" bson:"test_wallet_amount"`
	Direct                   int64              `json:"direct" bson:"direct"`
	Total_child              int64              `json:"total_child" bson:"total_child"`
	Role                     string             `json:"role,user" bson:"role,user"`
	Level                    int64              `json:"level" bson:"level"`
	Contact                  string             `json:"contact" bson:"contact"`
	Addr                     int64              `json:"addr" bson:"addr"`
	City                     int32              `json:"city" bson:"city"`
	Pending_withdrawl_amount int64              `json:"pending_withdrawl_amount" bson:"pending_withdrawl_amount"`
	Pincod                   int64              `json:"pincod" bson:"pincod"`
	Activation_date          time.Time          `json:"activation_date" bson:"activation_date"`
	Withdrawl_amount         int64              `json:"withdrawl_amount" bson:"withdrawl_amount"`
}
