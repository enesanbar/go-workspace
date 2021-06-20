data "template_file" "ecs_task_definition_template" {
  template = file("task_definition.json")

  vars = {
    app_name              = var.ecs_service_name
    task_definition_name  = var.ecs_service_name
    ecs_service_name      = var.ecs_service_name
    docker_image_repo     = var.docker_image_repo
    docker_memory         = var.docker_memory
    docker_container_port = var.docker_container_port
    region                = var.AWS_REGION
  }
}

resource "aws_ecs_task_definition" "app-info-task-definition" {
  container_definitions    = data.template_file.ecs_task_definition_template.rendered
  family                   = var.ecs_service_name
  cpu                      = var.docker_cpu
  memory                   = var.docker_memory
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.fargate_execution_role.arn
  task_role_arn            = aws_iam_role.fargate_execution_role.arn
}
