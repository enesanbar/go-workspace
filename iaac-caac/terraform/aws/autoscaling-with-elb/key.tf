resource "aws_key_pair" "mykeypair" {
  key_name   = "mykeypair"
  public_key = file(var.PATH_TO_PUBLIC_KEY)

  lifecycle {
    ignore_changes = [public_key]
  }

  tags = merge(
    {
      "Name" = "${local.prefix}-mykeypair"
    },
    local.common_tags,
  )
}
