variable "AWS_REGION" {
  default     = "eu-west-1"
  description = "AWS Region"
}

variable "AWS_PROFILE" {
  default = "default"
}

variable "infrastructure_remote_state_region" {}
variable "infrastructure_remote_state_bucket" {}
variable "infrastructure_remote_state_key" {}
variable "ecs_cluster_name" {}
variable "internet_cidr_blocks" {}
#variable "ecs_domain_name" {}
