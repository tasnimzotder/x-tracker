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

variable "xt-public-subnet-ciders" {
  type        = list(string)
  description = "List of public subnet CIDRs"
  default = [
    "10.0.1.0/24",
  ]
}

variable "xt-private-subnet-ciders" {
  type        = list(string)
  description = "List of private subnet CIDRs"
  default = [
    "10.0.4.0/24",
  ]
}
