
# s3www-minio

## Overview

`s3www-minio` is a lightweight, HTTP file server built in Go that serves files from a MinIO (S3-compatible) object storage backend. It includes built-in Prometheus metrics for observability and is designed to be deployed to Kubernetes using a Helm chart. Optional components include:

- Integration with distributed MinIO deployment
- Ingress routing with nginx
- CI/CD pipeline with GitHub Actions
- Terraform-based Helm deployments

## 🏗️ Architecture

```
[User] --> [Ingress] --> [s3-file-server Pod] --> [MinIO Service] --> [MinIO Distributed Storage]
                                  |
                                  --> [Prometheus Metrics Exporter]
```

## 📦 Deployment

### ✅ Prerequisites

- Kubernetes cluster (MicroK8s recommended for local dev)
- Helm 3+
- Terraform CLI (for GitHub Actions deployment)
- GitHub Packages (for image hosting)

### 🛠️ Build & Push Docker Image

```sh
make -C src docker-push
```

This will build the binary, package it into a Docker image, and push it to GitHub Packages.

### ☸️ Helm Chart Installation

```sh
helm dependency update ./charts/s3-file-server
helm install s3-file-server ./charts/s3-file-server
```

### Terraform Installation (Optional)

A GitHub Actions step uses Terraform to install this Helm chart.

### 🔗 CI/CD Ingress Access (For Demo purposes)

Ingress access is being done via ngrok proxy that connects to the exposed nginx ingress endpoint in GitHub actions runner, you can see this endpoint during actions run under:  `Output tunnel URL` step, for example:

```
https://05ac-20-57-79-82.ngrok-free.app
```

Please follow this link in your browser and click "Visit Site", this needs to be done due to limitations of ngrok free tier version.

Then test with:

```sh
curl -v -o giphy.gif https://05ac-20-57-79-82.ngrok-free.app/giphy.gif
```

## ⚙️ Configuration

Values in `values.yaml` exported as environmental variables from GitHub secrets, so please make sure these are present in the repositories Secret configuration section under Settings -> Actions -> Secrets (default values are given as an example):

```
env:
  - name: S3_BUCKET
    value: "my-bucket"
  - name: MINIO_ENDPOINT
    value: "s3-file-server-minio.default.svc.cluster.local:9000"
  - name: MINIO_ACCESS_KEY
    value: "minioadmin"
  - name: MINIO_SECRET_KEY
    value: "minioadmin"
```

## 📈 Observability

Metrics exported on `/metrics` endpoint include:

- `s3_file_downloads_total`
- `s3_file_download_errors_total`

## Operational Considerations

- Ensure secrets are configured correctly in repository Settings.
- Scale `s3-file-server` based on expected load.

## TODO

- add tests for the Go file server

## Authors

Maintained by Sergey Chernyshov

## License

GPL-3.0 license