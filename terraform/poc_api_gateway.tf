resource "aws_api_gateway_rest_api" "poc_gateway_redis" {
  name        = "ServerlessExample"
  description = "Terraform Serverless Application POC Golang Redis"
}

resource "aws_api_gateway_integration" "lambda" {
  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  resource_id = aws_api_gateway_method.proxy.resource_id
  http_method = aws_api_gateway_method.proxy.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_redis_create_lambda.invoke_arn
}

resource "aws_api_gateway_method" "proxy_root" {
  rest_api_id   = aws_api_gateway_rest_api.poc_gateway_redis.id
  resource_id   = aws_api_gateway_rest_api.poc_gateway_redis.root_resource_id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda_root" {
  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  resource_id = aws_api_gateway_method.proxy_root.resource_id
  http_method = aws_api_gateway_method.proxy_root.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_redis_create_lambda.invoke_arn
}

resource "aws_api_gateway_deployment" "deploy_gateway_lambda" {
  depends_on = [
    aws_api_gateway_integration.lambda,
    aws_api_gateway_integration.lambda_root,
  ]

  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  stage_name  = "poc"
}

