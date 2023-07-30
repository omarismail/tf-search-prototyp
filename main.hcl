provider "aws" {
}

search "aws_vpc_foo" {
  provider = "aws"
  query "explorer" {
    category = "network"
    region   = "us-east-2"
    tags     = ["env:prod"]
  }
}

search "aws_instance_bar" {
  provider = "aws"
  query "raw" {
    category = "compute"
    region   = "us-east-2"
    tags     = ["env:dev"]
  }
}
