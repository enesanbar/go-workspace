resource "aws_vpc" "production-vpc" {
  cidr_block           = var.VPC_CIDR
  enable_dns_hostnames = true

  tags = {
    Name = "Production VPC"
  }
}
