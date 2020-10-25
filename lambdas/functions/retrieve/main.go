package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/poc-redis-lambda-golang-terraform-aws/lambdas/aws/redis"
	"log"
)

type bodyRequest struct {
	RequestKey string `json:"key"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("im started")
	log.Println("The request body was: ", request.Body)

	var bodyRequestExtracted bodyRequest

	// Unmarshal the json, return 404 if error
	err := json.Unmarshal([]byte(request.Body), &bodyRequestExtracted)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	storage, err := redis.NewRedisDBStorage(ctx)

	if err != nil {
		log.Println("was error connecting to redis: ", err)
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              "Error",
			IsBase64Encoded:   false,
		}, err
	}
	log.Println("Trying to retrieve value")
	result, err := storage.GetConnection().Get(
		ctx,
		bodyRequestExtracted.RequestKey,
	).Result()

	log.Println(err)

	if err != nil {
		log.Println("We did an error")
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              err.Error(),
			IsBase64Encoded:   false,
		}, nil
	} else {
		log.Println("Successful")
		return events.APIGatewayProxyResponse{
			Body:       result,
			StatusCode: 200,
		}, nil
	}
}

func main() {
	log.Println("im starting from main")
	lambda.Start(handler)
}
