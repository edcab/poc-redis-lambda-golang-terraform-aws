resource "aws_ssm_parameter" "secretsmanager_redis_connection_endpoint" {
  name                    = "redis_connection_endpoint"
  type                    = "String"
  value                   = aws_elasticache_replication_group.default.configuration_endpoint_address
}