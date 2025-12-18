package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type Profile struct {
	Email   string `json:"email" bson:"email"`
	Age     int    `json:"age" bson:"age"`
	DOB     string `json:"dob" bson:"dob"`
	Contact string `json:"contact" bson:"contact"`
	Address string `json:"address" bson:"address"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
