data "aws_iam_policy_document" "ecs_cluster_assume_role_policy_document" {
  version = "2008-10-17"

  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs.amazonaws.com", "ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "fargate_execution_role" {
  name               = "${var.ecs_service_name}-iam-role"
  assume_role_policy = data.aws_iam_policy_document.ecs_cluster_assume_role_policy_document.json
}

data "aws_iam_policy_document" "ecs_cluster_assume_role_policy_document" {
  version = "2012-10-17"

  statement {
    effect = "Allow"
    actions = [
      "ecs:*",
      "ecr:*",
      "logs:*",
      "cloudwatch:*",
      "elasticloadbalancing:*"
    ]

    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "fargate_execution_role_policy" {
  name   = "${var.ecs_service_name}-iam-role-policy"
  policy = data.aws_iam_policy_document.ecs_cluster_assume_role_policy_document.json
  role   = aws_iam_role.fargate_execution_role.id
}
