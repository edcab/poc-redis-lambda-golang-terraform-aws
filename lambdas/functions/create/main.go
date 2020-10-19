package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("im started")

	//im creating a session
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("us-east-1")},
		//SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	//im creating a client
	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-east-1"))
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("redis_connection_endpoint"),
		WithDecryption: aws.Bool(false),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	value := *param.Parameter.Value
	fmt.Println(value)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Value was, %v", value),
		StatusCode: 200,
	}, nil
}

func main() {
	fmt.Println("im starting")
	lambda.Start(handler)
}
