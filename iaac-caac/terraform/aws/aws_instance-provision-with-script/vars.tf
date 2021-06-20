variable "AWS_REGION" {
  default = "eu-central-1"
}

variable "AWS_PROFILE" {
  default = "default"
}

variable "AMIS" {
  type = map(string)
  default = {
    us-east-1    = "ami-13be557e"
    us-west-2    = "ami-06b94666"
    eu-west-1    = "ami-0d729a60"
    eu-central-1 = "ami-0d527b8c289b4af7f"
  }
}

variable "PATH_TO_PRIVATE_KEY" {
  default = "../../keys/mykey"
}

variable "PATH_TO_PUBLIC_KEY" {
  default = "../../keys/mykey.pub"
}

variable "INSTANCE_USERNAME" {
  default = "ubuntu"
}
