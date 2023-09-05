terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.72.0"
      }
  }
}
provider "aws" {
  region = "us-east-2"
  access_key = var.accesskey
  secret_key = var.secretkey
}

module "front-end" {
  source        = "../modules/ec2"
  ec2_count     = 1
  ami_id        = "ami-0fb653ca2d3203ac1"
  instance_type = "t2.micro"
}

module "back-end" {
  source        = "../modules/ec2"
  ec2_count     = 1
  ami_id        = "ami-0fb653ca2d3203ac1"
  instance_type = "t2.micro"
}