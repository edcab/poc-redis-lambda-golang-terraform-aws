//Create lambda
resource "aws_lambda_function" "poc_redis_create_lambda" {
  filename      =  data.archive_file.OnCreate.output_path
  function_name = "poc_redis_create_lambda"
  role          = aws_iam_role.redis_lambda_role.arn
  handler       = "main"
  source_code_hash = data.archive_file.OnCreate.output_base64sha256
  runtime = "go1.x"
  vpc_config {
    subnet_ids          = flatten([aws_subnet.default.*.id])
    security_group_ids  = [aws_security_group.default.id]
  }
}

//All incoming requests to API Gateway must match with a configured resource and method in order to be handled.
//Define a single proxy resource:
resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  parent_id   = aws_api_gateway_rest_api.poc_gateway_redis.root_resource_id
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "proxy" {
  rest_api_id   = aws_api_gateway_rest_api.poc_gateway_redis.id
  resource_id   = aws_api_gateway_resource.proxy.id
  http_method   = "ANY"
  authorization = "NONE"
}

//Allow api gateway invoke lambda
resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_redis_create_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  # The "/*/*" portion grants access from any method on any resource
  # within the API Gateway REST API.
  source_arn = "${aws_api_gateway_rest_api.poc_gateway_redis.execution_arn}/*/*"
}