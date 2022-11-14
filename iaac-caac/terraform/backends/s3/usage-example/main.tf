terraform {
  backend "s3" {
    bucket         = "module-test-default-tf2d-state-bucket"
    key            = "state.tfstate"
    region         = "eu-west-1"
    profile        = "s3-backend"
    encrypt        = true
    role_arn       = "arn:aws:iam::876669042460:role/module-test-default-tf2d-tf-assume-role"
    dynamodb_table = "module-test-default-tf2d-state-lock"
  }

  required_version = ">= 1.2.4"

  required_providers {
    null = {
      source  = "hashicorp/null"
      version = "~> 3.0"
    }
  }
}

resource "null_resource" "motto" {
  triggers = {
    always = timestamp()
  }
  provisioner "local-exec" {
    command = "echo gotta catch em all" #A
  }
}
