terraform {
  required_version = "~>1.1.8"
}


provider "aws" {
  region  = var.aws_region
  profile = var.aws_profile
}