variable "AWS_REGION" {
  default = "eu-west-1"
}

variable "AWS_PROFILE" {
  default = "default"
}

variable "AMIS" {
  type = map(string)
  default = {
    us-east-1 = "ami-13be557e"
    us-west-2 = "ami-06b94666"
    eu-west-1 = "ami-844e0bf7"
  }
}

variable "project" {
  default = "module-test"
}

variable "contact" {
  default = "enesanbar@gmail.com"
}
