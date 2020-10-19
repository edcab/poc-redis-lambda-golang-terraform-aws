resource "aws_iam_role" "redis_lambda_role" {
  name               = "redis_messages_lambda_role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

#Creo la politica
resource "aws_iam_policy" "default_policy_lambda_create_redis" {
  name   = "lambda_create_redis_default_policy_all"
  policy = data.aws_iam_policy_document.default_lambda_policy.json
}

#Se la atacho
resource "aws_iam_role_policy_attachment" "lambda_read_to_redis" {
  role       = aws_iam_role.redis_lambda_role.name
  policy_arn = aws_iam_policy.default_policy_lambda_create_redis.arn
}