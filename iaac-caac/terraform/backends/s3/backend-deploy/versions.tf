terraform {
  required_providers {
    aws = {
      version = "~> 4.24.0"
    }
  }

  required_version = ">= 1.2.4"
}

locals {
  prefix = "${var.project}-${terraform.workspace}"
  common_tags = {
    Environment = terraform.workspace
    Project     = var.project
    Owner       = var.contact
    ManagedBy   = "Terraform"
  }
}
