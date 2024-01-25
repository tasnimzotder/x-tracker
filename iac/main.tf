terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }

    local = {
      source = "hashicorp/local"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_key_pair" "ssh_key" {
  key_name   = "ssh_key"
  public_key = file("~/.ssh/id_rsa.pub")
}

locals {
  instances = [
    {
      name          = "fe",
      instance_type = "t2.micro"
    },
    {
      name          = "be",
      instance_type = "t2.micro"
    }
  ]
}


resource "aws_instance" "x-tracker-instances" {
  ami           = "ami-058bd2d568351da34"
  instance_type = local.instances[count.index].instance_type
  key_name      = aws_key_pair.ssh_key.key_name
  count         = length(local.instances)

  credit_specification {
    cpu_credits = "unlimited"
  }

  iam_instance_profile = aws_iam_instance_profile.ec2_ecr_timestream_access_profile.name

  tags = {
    Name = "xtracker-${local.instances[count.index].name}"
  }
}


output "instances_pub_dns_addresses" {
  value = {
    for instance in aws_instance.x-tracker-instances :
    instance.tags.Name => instance.public_dns
  }
}

