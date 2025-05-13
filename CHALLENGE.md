# Challenge: Fully Self-Contained CI Environment Using GitHub

## Overview

This project demonstrates how to create a fully self-contained environment that requires nothing beyond GitHub to run. The entire project, including a multi-node MINIO installation and a Kubernetes environment, is deployed and executed within a GitHub Actions runner.

## Key Objectives

- **Self-Containment:** No external services or infrastructure are needed beyond GitHub.
- **Full E2E Testing:** The approach proves that complete end-to-end tests can be executed using only GitHub tools.

## Technical Highlights

- **GitHub Runner as Environment:** The entire stack is spun up inside a GitHub Actions runner.
- **Kubernetes & MINIO:** A distributed MINIO cluster with 4 nodes is deployed using Kubernetes.
- **Custom Helm Chart:** MINIO is deployed via a custom Helm chart, available at:
  [minio-distributed-chart](https://github.com/sergeychernyshov-code/minio-distributed-chart)
- **GitHub Pages Integration:** The Helm chart is built and published using GitHub Pages functionality.
- **Init Job:** On startup, an Init Job downloads a file from the assignment to initialize the environment.

## Debugging Approach

- **SSH Access:** To facilitate debugging, the `runner-open-ssh` GitHub Action is used to connect directly to the runner nodes.
- **Ingress Exposure:** `ngrok` is used to expose the Kubernetes ingress controller to the outside world for easier access and troubleshooting.

---

This setup proves that even complex, cloud-native systems can be tested in a completely isolated and GitHub-native way.
