provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

resource "helm_release" "s3_file_server" {
  name       = "s3-file-server"
  chart      = "../charts/s3-file-server"
  namespace  = "default"
  create_namespace = true

  set {
    name  = "ingress.enabled"
    value = "true"
  }

  set {
    name  = "ingress.hostname"
    value = "s3-file-server.local"
  }
}