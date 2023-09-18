variable "name" {
  type    = string
  default = "go-server"
}

variable "port" {
  type    = number
  default = 4000
}

variable "region" {
  type    = string
  default = "us-east-1"
}

variable "image" {
  type    = string
  default = "go-server"
}

variable "account" {
  type    = string
  default = "default"
}