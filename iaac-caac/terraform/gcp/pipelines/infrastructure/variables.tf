variable "project_id" {
  description = "The GCP project id"
  type        = string
}

variable "region" {
  default     = "europe-west3"
  description = "GCP region"
  type        = string
}

variable "namespace" {
  description = "The project namespace to use for unique resource naming"
  type        = string
}
