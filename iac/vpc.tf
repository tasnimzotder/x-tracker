# VPC
resource "aws_vpc" "xt" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "xt-vpc"
  }
}

resource "aws_subnet" "xt-public-subnets" {
  count             = length(var.xt-public-subnet-ciders)
  vpc_id            = aws_vpc.xt.id
  cidr_block        = element(var.xt-public-subnet-ciders, count.index)
  availability_zone = "us-east-1a"
  map_public_ip_on_launch = true

  tags = {
    Name = "xt-public-subnet-${count.index + 1}"
  }
}

#resource "aws_subnet" "xt-private-subnets" {
#  count             = length(var.xt-private-subnet-ciders)
#  vpc_id            = aws_vpc.xt.id
#  cidr_block        = element(var.xt-private-subnet-ciders, count.index)
#  availability_zone = "us-east-1a"
#
#  tags = {
#    Name = "xt-private-subnet-${count.index + 1}"
#  }
#}

resource "aws_internet_gateway" "xt-igw" {
  vpc_id = aws_vpc.xt.id

  tags = {
    Name = "xt-igw"
  }
}


resource "aws_route_table" "xt-public-rt" {
  vpc_id = aws_vpc.xt.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.xt-igw.id
  }

  tags = {
    Name = "xt-public-rt"
  }
}

resource "aws_route_table_association" "xt-public-rt-association" {
  count          = length(var.xt-public-subnet-ciders)
  subnet_id      = element(aws_subnet.xt-public-subnets.*.id, count.index)
  route_table_id = aws_route_table.xt-public-rt.id
}

# security groups
resource "aws_security_group" "xt-web-sg" {
  vpc_id      = aws_vpc.xt.id
  name        = "xt-sg"
  description = "Allow SSH & ICMP inbound traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = -1
    to_port     = -1
    protocol    = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "xt-sg"
  }
}

resource "aws_security_group" "xt-bastion-sg" {
  vpc_id      = aws_vpc.xt.id
  name        = "xt-bastion-sg"
  description = "Allow SSH inbound traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "xt-bastion-sg"
  }
}