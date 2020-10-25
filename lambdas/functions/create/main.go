package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/poc-redis-lambda-golang-terraform-aws/lambdas/aws/redis"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type bodyRequest struct {
	RequestKey   string `json:"key"`
	RequestValue string `json:"value"`
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

	////im creating a session
	//sess, err := session.NewSessionWithOptions(session.Options{
	//	Config:            aws.Config{Region: aws.String("us-east-1")},
	//	SharedConfigState: session.SharedConfigEnable,
	//})
	//log.Println("im sess", sess)
	//if err != nil {
	//	return events.APIGatewayProxyResponse{}, err
	//}
	//
	////im creating a client
	//ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-east-1"))
	//
	//log.Println("im ssmsvc", ssmsvc)
	//
	//param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
	//	Name:           aws.String("redis_connection_endpoint"),
	//	WithDecryption: aws.Bool(false),
	//})

	//log.Println("im param", param)

	//if err != nil {
	//	return events.APIGatewayProxyResponse{}, err
	//}

	//value := *param.Parameter.Value
	//fmt.Println("value from ssm is: ", value)

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
	log.Println("Trying insert a key")
	_, err = storage.GetConnection().Set(
		ctx,
		bodyRequestExtracted.RequestKey,
		bodyRequestExtracted.RequestValue,
		0,
	).Result()

	log.Println(err)

	if err != nil {
		log.Println("Hubo un error")
		return events.APIGatewayProxyResponse{
			StatusCode:        500,
			Headers:           nil,
			MultiValueHeaders: nil,
			Body:              err.Error(),
			IsBase64Encoded:   false,
		}, nil
	} else {
		log.Println("Exitoso")
		return events.APIGatewayProxyResponse{
			Body:       "Successfully Insert",
			StatusCode: 200,
		}, nil
	}
}

func main() {
	log.Println("im starting from main")
	lambda.Start(handler)
}
