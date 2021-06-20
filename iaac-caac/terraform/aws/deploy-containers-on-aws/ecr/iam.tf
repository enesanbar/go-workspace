# github iam user
resource "aws_iam_user" "github-actions" {
  name = "github-actions"

  tags = merge(
    {
      "Name" = "${local.prefix}-github-actions"
    },
    local.common_tags,
  )
}

data "aws_iam_policy_document" "github-actions-user-policy-document" {
  statement {
    effect = "Allow"
    actions = [
      "ecr:*",
    ]
    resources = [
      aws_ecr_repository.kuard.arn,
    ]
  }
}

resource "aws_iam_user_policy" "github-actions-user-policy" {
  name   = "github-user-policy"
  user   = aws_iam_user.github-actions.name
  policy = data.aws_iam_policy_document.github-actions-user-policy-document.json
}
