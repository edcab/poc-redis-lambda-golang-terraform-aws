package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")

	//ctx = context.Background()
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("im starting")
	log.Println("im started")
	log.Println("im started")
	fmt.Println(ctx)

	env, _ := os.LookupEnv("elasticCache")
	log.Println("Environment elasticCache: ", env)

	//redisdb.NewRedisDBStorage(ctx)

	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIP
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string(ip)),
		StatusCode: 200,
	}, nil
}

func main() {
	fmt.Println("im starting")
	log.Println("im starting")
	lambda.Start(handler)

	//request := events.APIGatewayProxyRequest{
	//	Resource:                        "",
	//	Path:                            "",
	//	HTTPMethod:                      "",
	//	Headers:                         nil,
	//	MultiValueHeaders:               nil,
	//	QueryStringParameters:           nil,
	//	MultiValueQueryStringParameters: nil,
	//	PathParameters:                  nil,
	//	StageVariables:                  nil,
	//	RequestContext:                  events.APIGatewayProxyRequestContext{},
	//	Body:                            "",
	//	IsBase64Encoded:                 false,
	//}

	//handler()
}
