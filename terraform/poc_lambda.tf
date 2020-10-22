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

//Retrieve lambda
resource "aws_lambda_function" "poc_redis_retrieve_lambda" {
  filename      =  data.archive_file.OnRetrieve.output_path
  function_name = "poc_redis_retrieve_lambda"
  role          = aws_iam_role.redis_lambda_role.arn
  handler       = "main"
  source_code_hash = data.archive_file.OnRetrieve.output_base64sha256
  runtime = "go1.x"
  vpc_config {
    subnet_ids          = flatten([aws_subnet.default.*.id])
    security_group_ids  = [aws_security_group.default.id]
  }
}

//All incoming requests to API Gateway must match with a configured resource and method in order to be handled.
//Define a single proxy resource:
resource "aws_api_gateway_resource" "proxy_create" {
  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  parent_id   = aws_api_gateway_rest_api.poc_gateway_redis.root_resource_id
  path_part   = "create"
}

resource "aws_api_gateway_resource" "proxy_retrieve" {
  rest_api_id = aws_api_gateway_rest_api.poc_gateway_redis.id
  parent_id   = aws_api_gateway_rest_api.poc_gateway_redis.root_resource_id
  path_part   = "retrieve"
}

resource "aws_api_gateway_method" "post_create" {
  rest_api_id   = aws_api_gateway_rest_api.poc_gateway_redis.id
  resource_id   = aws_api_gateway_resource.proxy_create.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "post_retrieve" {
  rest_api_id   = aws_api_gateway_rest_api.poc_gateway_redis.id
  resource_id   = aws_api_gateway_resource.proxy_retrieve.id
  http_method   = "POST"
  authorization = "NONE"
}

//Allow api gateway invoke lambda create
resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_redis_create_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  # The "/*/*" portion grants access from any method on any resource
  # within the API Gateway REST API.
  source_arn = "${aws_api_gateway_rest_api.poc_gateway_redis.execution_arn}/*/*"
}

//Allow api gateway invoke lambda retrieve
resource "aws_lambda_permission" "apigw_retrieve" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_redis_retrieve_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  # The "/*/*" portion grants access from any method on any resource
  # within the API Gateway REST API.
  source_arn = "${aws_api_gateway_rest_api.poc_gateway_redis.execution_arn}/*/*"
}