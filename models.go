package models

import "time"

// User is struct data model for user
// Using In Api Service And Sync Service
type User struct {
	Name      string    `json:"name" bson:"name"`
	Age       uint8     `json:"age" bson:"age"`
	Address   string    `json:"adr,omitempty" bson:"adr,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
