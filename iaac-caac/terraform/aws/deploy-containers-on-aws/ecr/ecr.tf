resource "aws_ecr_repository" "kuard" {
  name                 = "kuard"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  tags = merge(
    {
      "Name" = "${local.prefix}-kuard"
    },
    local.common_tags,
  )
}
