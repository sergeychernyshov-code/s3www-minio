provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

resource "helm_release" "s3_file_server" {
  name              = "s3-file-server"
  chart             = "../charts/s3-file-server"
  namespace         = "default"
  dependency_update = true

  set {
    name  = "env[0].name"
    value = "S3_BUCKET"
  }

  set {
    name  = "env[0].value"
    value = var.s3_bucket
  }

  set {
    name  = "env[1].name"
    value = "MINIO_ENDPOINT"
  }

  set {
    name  = "env[1].value"
    value = var.minio_endpoint
  }

  set {
    name  = "env[2].name"
    value = "MINIO_ACCESS_KEY"
  }

  set {
    name  = "env[2].value"
    value = var.minio_access_key
  }

  set {
    name  = "env[3].name"
    value = "MINIO_SECRET_KEY"
  }

  set {
    name  = "env[3].value"
    value = var.minio_secret_key
  }
}

variable "s3_bucket" {}
variable "minio_endpoint" {}
variable "minio_access_key" {}
variable "minio_secret_key" {}
