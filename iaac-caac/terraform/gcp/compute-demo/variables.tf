variable "project" {}

variable "credentials_file" {}

variable "region" {
  default = "us-central1"
}

variable "zone" {
  default = "us-central1-c"
}

variable "environment" {
  type    = string
  default = "loadtest"
}

variable "machine_types" {
  type = map(any)
  default = {
    loadtest = "f1-micro"
  }
}