resource "aws_key_pair" "ssh_key" {
  key_name   = "ssh_key"
  public_key = file("~/.ssh/id_rsa.pub")
}


# EBS
resource "aws_ebs_volume" "xt-ebs" {
  availability_zone = "us-east-1a"
  size              = 20
  type              = "gp3"
  encrypted         = false

  tags = {
    Name = "xt-ebs"
  }
}

resource "aws_volume_attachment" "xt-ebs-attach" {
  device_name = "/dev/sdf"
  volume_id   = aws_ebs_volume.xt-ebs.id
  instance_id = aws_instance.xt-be.id
}

# EC2 instances
resource "aws_instance" "xt-be" {
  ami               = "ami-058bd2d568351da34"
  instance_type     = "t2.micro"
  key_name          = aws_key_pair.ssh_key.key_name
  #  count         = length(local.instances)
  availability_zone = "us-east-1a"

  #  associate_public_ip_address = true
  subnet_id              = aws_subnet.xt-public-subnets[0].id
  vpc_security_group_ids = [
    aws_security_group.xt-web-sg.id
  ]

  credit_specification {
    cpu_credits = "unlimited"
  }

  iam_instance_profile = aws_iam_instance_profile.ec2_ecr_timestream_access_profile.name

  tags = {
    Name = "xt-be"
  }
}


output "ec2_instance_be" {
#  value = {
#    dns_addresses = aws_instance.xt-be.*.public_dns
#  }

  value = aws_instance.xt-be.public_dns
}

