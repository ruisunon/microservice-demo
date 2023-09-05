terraform {
  backend "s3" {
    bucket = "terraform-state-zhajili"
    key    = "petclinic/terraform.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  # No need to add since we have added all configuration including region,access_key and secret_access_key via aws configure list command 
  # it is located at ~/.aws/credentials
}


variable "vpc_cidr_block" {}
variable "subnet_cidr_block" {}
variable "availibility_zone" {}
variable "env_prefix" {}
variable "my_source_ip" {}
variable "instance_type" {}
variable "my_public_key_location" {}
variable "ssh_private_key" {}
variable "user_name" {}

resource "aws_vpc" "petclinic_vpc" {
  cidr_block = var.vpc_cidr_block
  enable_dns_hostnames = true # enables public DNS
  tags = {
    "Name" = "${var.env_prefix}-vpc"
  }
}

resource "aws_subnet" "my_app_subnet-1" {
  vpc_id = aws_vpc.petclinic_vpc.id
  cidr_block = var.subnet_cidr_block 
  availability_zone = var.availibility_zone
  tags = {
    "Name" = "${var.env_prefix}"
  } 
}

resource "aws_internet_gateway" "petclinic_vpc_internet_gateway" {
     vpc_id = aws_vpc.petclinic_vpc.id
     tags = {
     "Name" = "${var.env_prefix}-internet-gateway"}
}


resource "aws_default_route_table" "main-route-table" {
  default_route_table_id = aws_vpc.petclinic_vpc.default_route_table_id

   route {
    cidr_block = "0.0.0.0/0"
    gateway_id =aws_internet_gateway.petclinic_vpc_internet_gateway.id
   }
   tags = {
     "Name" = "${var.env_prefix}-main-route-table"
   }
}


resource "aws_security_group" "petclinic_sg" {
  name        = "PETCLINIC_SG"
  description = "Allow Inbound SSH and TCP port 8080 traffic"
  vpc_id      = aws_vpc.petclinic_vpc.id

  ingress {
    description      = "SSH from VPC"
    from_port        = 22
    to_port          = 22
    protocol         = "tcp"
    cidr_blocks      = [var.my_source_ip]
  }

    ingress {
    description      = "HTTP from VPC"
    from_port        = 8080
    to_port          = 8080
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
  }

    ingress {
    description      = "MYSQL port"
    from_port        = 3306
    to_port          = 3306
    protocol         = "tcp"
    cidr_blocks      = ["0.0.0.0/0"]
  }
  
  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1" # any protocol
    cidr_blocks      = ["0.0.0.0/0"]
  }

  tags = {
     "Name" = "${var.env_prefix}-sg"
   }
}
data "aws_ami" "ubuntu18" {
    most_recent = true
    owners = ["amazon"]

    filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*"]
    }

    filter {
        name = "virtualization-type"
        values = ["hvm"]
    }
}

data "aws_ami" "ubuntu22" {
    most_recent = true
    owners = ["amazon"]
    filter {
        name = "name"
        values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
    }
    filter {
        name = "virtualization-type"
        values = ["hvm"]
    }
}

resource "aws_instance" "petclinic_mysql" {
  ami = data.aws_ami.ubuntu18.id
  instance_type = var.instance_type
  availability_zone = var.availibility_zone
  subnet_id =aws_subnet.my_app_subnet-1.id
  private_ip = "10.20.15.200"
  vpc_security_group_ids = [aws_security_group.petclinic_sg.id]

  associate_public_ip_address = true
  key_name = "AWS_KEY_PAIR"
  tags = {
    "Name" = "${var.env_prefix}-db"
  }

  # connection {
  #   type     = "ssh"
  #   user     = "ubuntu"
  #   host     = self.public_dns
  #   private_key = file("~/.ssh/AWS_KEY_PAIR.pem")
  # }

  # provisioner "file" {
  #   source      = "mysql.sh"
  #   destination = "/tmp/script.sh"
  # }

  # provisioner "remote-exec" {
  #   inline = [
  #     "chmod +x /tmp/script.sh",
  #     "/tmp/script.sh args",
  #     "sudo cloud-init status --wait",
  #   ]
  # }
  user_data = file("mysql.sh") # execute user-data script which will install dependencies for MySQL DB
}


resource "aws_instance" "petclinic_application" {
  ami = data.aws_ami.ubuntu22.id
  instance_type = var.instance_type
  availability_zone = var.availibility_zone

  subnet_id =aws_subnet.my_app_subnet-1.id 
  vpc_security_group_ids = [aws_security_group.petclinic_sg.id]

  associate_public_ip_address = true
  key_name = "AWS_KEY_PAIR"

  tags = {
    "Name" = "${var.env_prefix}-app"
  }

  # connection {
  #   type     = "ssh"
  #   user     = "ubuntu"
  #   host     = self.public_dns
  #   private_key = file("~/.ssh/AWS_KEY_PAIR.pem")
  # }

  # provisioner "file" {
  #   source      = "application.sh"
  #   destination = "/tmp/script.sh"
  # }
  # provisioner "remote-exec" {
  #   inline = [
  #     "chmod +x /tmp/script.sh",
  #     "/tmp/script.sh args",
  #     "sudo cloud-init status --wait",
  #   ]
  # }

  user_data = file("application.sh") # execute user-data script which will install dependencies for APP

}

output "application_public_public_dns" {
    value = aws_instance.petclinic_application.public_dns
}
