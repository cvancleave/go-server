terraform {
  backend "s3" {
    bucket = "go-server-terraform-state"
    key    = "go-server/terraform.tfstate"
    region = "us-east-1"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.2"
}

provider "aws" {
  region = "us-east-1"
}
