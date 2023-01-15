package main

import (
	"context"
	"encoding/json"
	"github.com/Metegultekin98/Go-Dynamo-Lambda-Auth-API/controllers"
	"github.com/Metegultekin98/Go-Dynamo-Lambda-Auth-API/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

var validate *validator.Validate = validator.New()

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request is %#v", req)

	switch req.HTTPMethod {
	case "POST":
		return processPost(ctx, req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func processPost(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var loginForm models.UserAuth
	err := json.Unmarshal([]byte(req.Body), &loginForm)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	err = validate.Struct(&loginForm)
	if err != nil {
		return clientError(http.StatusBadRequest)
	}

	res, err := controllers.Login(ctx, loginForm)
	if err != nil {
		return serverError(err)
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsonData),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(status),
		StatusCode: status,
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())

	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(http.StatusInternalServerError),
		StatusCode: http.StatusInternalServerError,
	}, nil
}
