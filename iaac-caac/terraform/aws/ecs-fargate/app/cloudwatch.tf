resource "aws_cloudwatch_log_group" "appinfo_log_group" {
  name = "${var.ecs_service_name}-LogGroup"
}
