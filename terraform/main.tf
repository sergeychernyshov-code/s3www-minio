provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

locals {
  env_vars = [
    {
      name  = "S3_BUCKET"
      value = var.s3_bucket
    },
    {
      name  = "MINIO_ENDPOINT"
      value = var.minio_endpoint
    },
    {
      name  = "MINIO_ACCESS_KEY"
      value = var.minio_access_key
    },
    {
      name  = "MINIO_SECRET_KEY"
      value = var.minio_secret_key
    },
    {
      name  = "IMAGE_REPO"
      value = var.image_repository
    },
    {
      name  = "VERSION"
      value = var.image_tag
    }
  ]
}

resource "helm_release" "s3_file_server" {
  name              = "s3-file-server"
  chart             = "../charts/s3-file-server"
  namespace         = "default"
  dependency_update = true

  dynamic "set" {
    for_each = local.env_vars
    content {
      name  = "env[${index(local.env_vars, set.value)}].name"
      value = set.value.name
    }
  }

  dynamic "set" {
    for_each = local.env_vars
    content {
      name  = "env[${index(local.env_vars, set.value)}].value"
      value = set.value.value
    }
  }
}

variable "s3_bucket" {}
variable "minio_endpoint" {}
variable "minio_access_key" {}
variable "minio_secret_key" {}
variable "image_repository" {}
variable "image_tag" {}