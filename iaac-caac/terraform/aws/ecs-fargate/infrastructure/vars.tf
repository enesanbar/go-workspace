variable "AWS_REGION" {
  default     = "eu-west-1"
  description = "AWS Region"
}

variable "AWS_PROFILE" {
  default = "default"
}

variable "VPC_CIDR" {
  default     = "10.0.0.0/16"
  description = "VPC CIDR block"
}

variable "VPC_PUBLIC_SUBNET_1_CIDR" {
  description = "Public subnet 1 CIDR"
}

variable "VPC_PUBLIC_SUBNET_2_CIDR" {
  description = "Public subnet 2 CIDR"
}

variable "VPC_PUBLIC_SUBNET_3_CIDR" {
  description = "Public subnet 3 CIDR"
}

variable "VPC_PRIVATE_SUBNET_1_CIDR" {
  description = "Private subnet 1 CIDR"
}

variable "VPC_PRIVATE_SUBNET_2_CIDR" {
  description = "Private subnet 2 CIDR"
}

variable "VPC_PRIVATE_SUBNET_3_CIDR" {
  description = "Private subnet 3 CIDR"
}
