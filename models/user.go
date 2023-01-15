package models

type User struct {
	Id       string `json:"id" dynamodbav:"id"`
	Username string `json:"username" dynamodbav:"username"`
	Email    string `json:"email" dynamodbav:"email"`
	Password string `json:"password" dynamodbav:"password"`
	Status   bool   `json:"status" dynamodbav:"status"`
}

type CreateUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
