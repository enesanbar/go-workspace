terraform {
  required_providers {
    aws = {
      version = "~> 4.8.0"
    }
    random = {
      version = ">= 3.1.2"
    }
  }

  required_version = ">= 1.1.7"
}
