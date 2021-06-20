variable "AWS_REGION" {
  default     = "eu-west-1"
  description = "AWS Region"
}

variable "AWS_PROFILE" {
  default = "default"
}

variable "platform_remote_state_region" {}
variable "platform_remote_state_bucket" {}
variable "platform_remote_state_key" {}

variable "ecs_service_name" {}
variable "docker_image_repo" {}
variable "docker_memory" {}
variable "docker_cpu" {}
variable "docker_container_port" {}
variable "desired_task_number" {}
