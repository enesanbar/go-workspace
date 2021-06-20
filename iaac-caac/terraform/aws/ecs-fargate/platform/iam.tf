data "aws_iam_policy_document" "ecs_cluster_assume_role_policy_document" {
  version = "2008-10-17"

  statement {
    effect = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs.amazonaws.com", "ec2.amazonaws.com", "application-autoscaling.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ecs_cluster_role" {
  name = "${var.ecs_cluster_name}-ECS-Role"
  assume_role_policy = data.aws_iam_policy_document.ecs_cluster_assume_role_policy_document.json
}

data "aws_iam_policy" "AmazonEC2ContainerServiceforEC2RolePolicy" {
  name = "AmazonEC2ContainerServiceforEC2Role"
}

resource "aws_iam_role_policy_attachment" "ecs_cluster_role_policy" {
  role   = aws_iam_role.ecs_cluster_role.id
  policy_arn = data.aws_iam_policy.AmazonEC2ContainerServiceforEC2RolePolicy.arn
}
