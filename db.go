package main

import (
	"context"
	"github.com/Metegultekin98/Go-Dynamo-Lambda-Auth-API/models"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

const TableName = "Users"

var db dynamodb.Client

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	db = *dynamodb.NewFromConfig(sdkConfig)
}

func getUser(ctx context.Context, id string) *models.User {

}
