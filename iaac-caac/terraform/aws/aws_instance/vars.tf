variable "AWS_PROFILE" {
  default = "default"
}

variable "AWS_REGION" {
  default = "eu-central-1"
}

variable "AMIS" {
  type = map(string)
  default = {
    us-east-1    = "ami-13be557e"
    us-west-2    = "ami-06b94666"
    eu-west-1    = "ami-0d729a60"
    eu-central-1 = "ami-0dcc0ebde7b2e00db"
  }
}

variable "instance_name" {
  description = "Value of the Name tag for the EC2 instance"
  type        = string
  default     = "My EC2 Instance"
}