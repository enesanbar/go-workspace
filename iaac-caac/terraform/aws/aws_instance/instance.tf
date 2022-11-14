data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-gp2"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }
}

resource "aws_instance" "app_server" {
  ami           = data.aws_ami.amazon_linux.id
  instance_type = "t2.micro"

  # attach ec2 instance a role:
  iam_instance_profile = aws_iam_instance_profile.iam-read-only-role-instanceprofile.name

  tags = merge(
    {
      "Name" = "${local.prefix}-${var.instance_name}"
    },
    local.common_tags,
  )
}
