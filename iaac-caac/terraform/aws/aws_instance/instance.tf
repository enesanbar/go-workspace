resource "aws_instance" "app_server" {
  ami           = lookup(var.AMIS, var.AWS_REGION, "") # last parameter is the default value
  instance_type = "t2.micro"

  # attach ec2 instance a role:
  iam_instance_profile = aws_iam_instance_profile.iam-read-only-role-instanceprofile.name

  tags = {
    Name = var.instance_name
  }
}
