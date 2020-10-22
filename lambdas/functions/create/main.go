package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/poc-redis-lambda-golang-terraform-aws/lambdas/config/redis"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("im started")

	//im creating a session
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("us-east-1")},
		SharedConfigState: session.SharedConfigEnable,
	})
	log.Println("im sess", sess)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	//im creating a client
	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	log.Println("im ssmsvc", ssmsvc)

	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("redis_connection_endpoint"),
		WithDecryption: aws.Bool(false),
	})

	log.Println("im param", param)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	value := *param.Parameter.Value
	fmt.Println("value from ssm is: ", value)

	//value := "redis-cluster.sfofiy.clustercfg.use1.cache.amazonaws.com"

	storage, err := redis.NewRedisDBStorage(ctx)

	if err != nil {
		log.Println("was error connecting to redis: ", err)
		return events.APIGatewayProxyResponse{}, err
	}
	log.Println("intentando traer una llave")
	stringValue := storage.GetConnection().Get(ctx, "mykey")

	if stringValue != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Value was, %v", stringValue),
			StatusCode: 200,
		}, nil
	} else {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("For key %v not exists the Value", "mykey"),
			StatusCode: 404,
		}, nil
	}
}

func main() {
	log.Println("im starting from main")
	lambda.Start(handler)
}
