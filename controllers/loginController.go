package controllers

import (
	"context"
	"github.com/Metegultekin98/Go-Dynamo-Lambda-Auth-API/initializers"
	"github.com/Metegultekin98/Go-Dynamo-Lambda-Auth-API/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

const TableName = "Users"

var db dynamodb.Client

func init() {
	initializers.LoadEnv()
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	db = *dynamodb.NewFromConfig(sdkConfig)
}

func getUser(ctx context.Context, username string) (*models.User, error) {
	key, err := attributevalue.Marshal(username)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"Username": key,
		},
	}

	result, err := db.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	user := new(models.User)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(ctx context.Context, loginForm models.UserAuth) (*models.AuthResponse, error) {
	user, err := getUser(ctx, loginForm.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginForm.Password))
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString(os.Getenv("HMAC_KEY"))
	if err != nil {
		return nil, err
	}

	res := models.AuthResponse{
		User:        user,
		AccessToken: tokenStr,
	}
	return &res, nil
}
