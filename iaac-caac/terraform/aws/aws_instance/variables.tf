variable "AWS_PROFILE" {
  default = "default"
}

variable "AWS_REGION" {
  default = "eu-west-1"
}

variable "instance_name" {
  description = "Value of the Name tag for the EC2 instance"
  type        = string
  default     = "My EC2 Instance"
}

variable "project" {
  default = "ec2-instance"
}

variable "contact" {
  default = "enesanbar@gmail.com"
}
