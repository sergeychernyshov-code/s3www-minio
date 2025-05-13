# s3www-minio

## Overview

`s3www-minio` is a lightweight, production-ready HTTP file server built in Go that serves files from a MinIO (S3-compatible) object storage backend. It includes built-in Prometheus metrics for observability and is designed to be deployed to Kubernetes using a Helm chart. Optional components include:

- Integration with distributed MinIO deployment
- Ingress routing with nginx
- CI/CD pipeline with GitHub Actions
- Terraform-based Helm deployments

## Architecture

```
[User] --> [Ingress] --> [s3-file-server Pod] --> [MinIO Service] --> [MinIO Distributed Storage]
                                  |
                                  --> [Prometheus Metrics Exporter]
```

## Deployment

### Prerequisites

- Kubernetes cluster (MicroK8s recommended for local dev)
- Helm 3+
- Terraform CLI (for GitHub Actions deployment)
- GitHub Packages (for image hosting)

### Build & Push Docker Image

```sh
make -C src docker-push
```

This will build the binary, package it into a Docker image, and push it to GitHub Packages.

### Helm Chart Installation

```sh
helm dependency update ./charts/s3-file-server
helm install s3-file-server ./charts/s3-file-server
```

### Terraform Installation (Optional)

A GitHub Actions step uses Terraform to install this Helm chart.

### Ingress Access

The ingress rule assumes the service is exposed as `s3-file-server.local`. For local development, add the following to `/etc/hosts`:

```
127.0.0.1 s3-file-server.local
```

Then test with:

```sh
curl -v -o myfile.txt http://s3-file-server.local/myfile.txt
```

## Configuration

Values in `values.yaml`:

- `minio.accessKey`
- `minio.secretKey`
- `server.bucket`
- `server.region`
- `server.endpoint`

## Observability

Metrics exported on `/metrics` endpoint include:

- `s3_file_downloads_total`
- `s3_file_download_errors_total`

## Operational Considerations

- Ensure secrets are mounted securely.
- Tune retry logic in MinIO client as needed.
- Scale `s3-file-server` based on expected load.

## Authors

Maintained by [Your Name].

## License

MIT