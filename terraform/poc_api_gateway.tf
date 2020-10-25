resource "aws_api_gateway_rest_api" "poc_gateway_redis" {
  name        = "ServerlessExample"
  description = "Terraform Serverless Application POC Golang Redis"
}

resource "aws_api_gateway_method_settings" "general_settings" {
  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  stage_name  = aws_api_gateway_deployment.deploy_gateway_lambda.stage_name
  method_path = "*/*"

  settings {
    # Enable CloudWatch logging and metrics
    metrics_enabled        = true
    data_trace_enabled     = true
    logging_level          = "INFO"

    # Limit the rate of calls to prevent abuse and unwanted charges
    throttling_rate_limit  = 100
    throttling_burst_limit = 50
  }
}

resource "aws_api_gateway_integration" "lambda_create" {
  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  resource_id = aws_api_gateway_method.post_create.resource_id
  http_method = aws_api_gateway_method.post_create.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_redis_create_lambda.invoke_arn
}

resource "aws_api_gateway_integration" "lambda_retrieve" {
  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  resource_id = aws_api_gateway_method.post_retrieve.resource_id
  http_method = aws_api_gateway_method.post_retrieve.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_redis_retrieve_lambda.invoke_arn
}

resource "aws_api_gateway_deployment" "deploy_gateway_lambda" {
  depends_on = [
    aws_api_gateway_integration.lambda_create,
    aws_api_gateway_integration.lambda_retrieve,
  ]

  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  stage_name  = "poc"
}



