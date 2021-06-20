terraform {
  required_providers {
    aws = {
      version = "~> 4.8.0"
    }
  }

  required_version = ">= 1.1.7"

  backend "s3" {}
}

data "terraform_remote_state" "infrastructure" {
  backend = "s3"

  config = {
    region = var.infrastructure_remote_state_region
    bucket = var.infrastructure_remote_state_bucket
    key    = var.infrastructure_remote_state_key
  }
}