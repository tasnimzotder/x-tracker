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

#locals {
#  vpc = {
#    name = "x-tracker"
#    cidr = "10.0.0.0/16"
#    azs  = ["us-east-1a", "us-east-1b", "us-east-1c"]
#
#    public_subnets = [
#      {
#        cidr = "10.0.1.0/24"
#      }
#    ]
#
#    private_subnets = [
#      {
#        cidr = "10.0.2.0/24"
#      }
#    ]
#  }
#}

locals {
  iot_rule = {
    name    = "xt-rule"
    sql     = "SELECT * FROM 'xt/data'"
    actions = [
      {
        timestream = {
          database_name = "xt"
          table_name    = "data"
          dimensions    = ["device_id"]
          #          role_arn      = aws_iam_role.iot_timestream_role.arn
        }
      }
    ]
  }
}

variable "xt-public-subnet-ciders" {
  type        = list(string)
  description = "List of public subnet CIDRs"
  default     = [
    "10.0.1.0/24",
  ]
}

variable "xt-private-subnet-ciders" {
  type        = list(string)
  description = "List of private subnet CIDRs"
  default     = [
    "10.0.4.0/24",
  ]
}