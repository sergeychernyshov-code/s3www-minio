s3www-minio System Documentation

Overview

This microservice provides HTTP access to files stored in an S3-compatible object store (MinIO). It is built with:

Go (uses minio-go client)

MinIO (distributed via Helm dependency)

Kubernetes (via Helm chart)

NGINX Ingress (optionally using ngrok)

Prometheus Metrics (file downloads and error counts)

GitHub Actions for CI/CD

Terraform for infrastructure and Helm chart management

Architecture

         Ingress (NGINX/ngrok)
               |
         s3-file-server (Go)
               |
         MinIO Server (S3-compatible)

Configuration

Environment Variables

Variable

Description

Required

Example

MINIO_ENDPOINT

MinIO service address

✅

minio.default.svc:9000

MINIO_BUCKET

Target bucket name

✅

my-bucket

MINIO_ACCESS_KEY

MinIO access key

✅

admin

MINIO_SECRET_KEY

MinIO secret key

✅

password123

PORT

Port the HTTP server binds to

❌

8080

Deployment

1. Install MicroK8s (for local testing)

sudo snap install microk8s --classic
microk8s enable dns storage ingress helm3 prometheus

2. Build and Push Docker Image

cd src
make build
make docker-build
make docker-push

3. Helm Chart Deployment

Helm chart includes:

Service

Ingress

Deployment

Metrics service

Dependency: minio-distributed

cd helm/s3-file-server
helm dependency update
helm install s3-file-server . -n default

Or deploy using Terraform (see below).

Kubernetes Access

Access via Ingress

Ensure your /etc/hosts has:

127.0.0.1 s3-file-server.local

Access:

curl -v -o giphy.gif http://s3-file-server.local/giphy.gif

If using ngrok:

ngrok http 80

Update your Ingress or test with ngrok URL.

GitHub Actions CI/CD

CI/CD builds, pushes, and deploys via Terraform.

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/setup-go@v5
        with:
          go-version: 1.23.x

      - run: make -C src build
      - run: make -C src docker-build
      - run: make -C src docker-push
      - run: terraform init && terraform apply -auto-approve

Prometheus Metrics

Service exposes metrics at /metrics.

s3_file_downloads_total

s3_file_errors_total

Example alert:

- alert: HighFileErrorRate
  expr: rate(s3_file_errors_total[5m]) > 5
  for: 2m
  labels:
    severity: warning
  annotations:
    summary: High error rate accessing files

Helm Chart Layout

helm/s3-file-server/
├── charts/
├── templates/
│   ├── deployment.yaml
│   ├── ingress.yaml
│   ├── service.yaml
├── Chart.yaml
├── values.yaml
└── requirements.yaml

Security Considerations

Secrets managed via K8s secrets or Terraform

MinIO is internal unless explicitly exposed

Use HTTPS + Cert-Manager in production

File access can be secured with signed URLs

Terraform Usage

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

resource "helm_release" "s3-file-server" {
  name       = "s3-file-server"
  chart      = "./helm/s3-file-server"
  namespace  = "default"
}

Operational Tips

kubectl logs deploy/s3-file-server to check logs

kubectl port-forward svc/s3-file-server 8080:80

Final Validation Checklist



Contributors

DevOps: CI/CD, Terraform

Backend: Go service, metrics

QA: Helm testing, Ingress validation

Versioning

Follows Semantic Versioning. Maintains CHANGELOG.md.

For questions or contributions, open an issue or PR on GitHub.

