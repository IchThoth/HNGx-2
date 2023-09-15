package main

type User struct {
	Id    string `json:"user_id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
