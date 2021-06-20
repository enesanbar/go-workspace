resource "aws_security_group" "app_security_group" {
  name        = "${var.ecs_service_name}-sg"
  description = "Security group for app-info to communicate"
  vpc_id      = data.terraform_remote_state.platform.outputs.vpc_id

  ingress {
    from_port   = 3000
    protocol    = "TCP"
    to_port     = 3000
    cidr_blocks = [data.terraform_remote_state.platform.outputs.vpc_cidr_blocks]
  }

  egress {
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.ecs_service_name}-sg"
  }
}