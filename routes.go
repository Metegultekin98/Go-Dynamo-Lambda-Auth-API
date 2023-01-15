package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator/v10"
	"log"
)

var validate *validator.Validate = validator.New()

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request is %#v", req)

	switch req.HTTPMethod {
	case "GET":
		return nil, nil
	}
}

func processGet(ctx context.Context, id string) (events.APIGatewayProxyResponse, error) {

}
