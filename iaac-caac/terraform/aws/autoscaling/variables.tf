variable "AWS_REGION" {
  default = "eu-west-1"
}

variable "AWS_PROFILE" {
  default = "default"
}

variable "PATH_TO_PRIVATE_KEY" {
  default = "../../keys/mykey"
}

variable "PATH_TO_PUBLIC_KEY" {
  default = "../../keys/mykey.pub"
}

variable "project" {
  default = "autoscaling-test"
}

variable "contact" {
  default = "enesanbar@gmail.com"
}
