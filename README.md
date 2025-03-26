# K8Backup - Kubernetes Backup & Restore CLI

## Overview
K8Backup is a command-line tool built with Golang, Cobra, and the Kubernetes API to streamline backing up and restoring Kubernetes resources with a single command. It supports backups for Pods, Deployments, and Persistent Volumes (PVCs), ensuring quick recovery in case of failures.

## Features
- **Backup & Restore** Kubernetes resources (Pods, Deployments, PVCs)
- **List & Delete** backup objects
- **Supports Minikube volume snapshots**
- **Simple CLI commands** for efficient usage

## Prerequisites
### 1. Kubeconfig File
The `.kube/config` file is required to authenticate and interact with your Kubernetes cluster. It contains the cluster configuration, credentials, and context information.

**To obtain and set up your kubeconfig file:**
```sh
# Copy kubeconfig from remote cluster to local machine (if applicable)
scp user@remote-server:/path/to/kubeconfig ~/.kube/config

# Set environment variable (if using a custom kubeconfig path)
export KUBECONFIG=/custom/path/to/kubeconfig
```

### 2. Minikube (For Volume Snapshots)
If using Minikube for PVC snapshots, enable the required addons:
```sh
minikube addons enable csi-hostpath-driver
minikube addons enable volumesnapshots
```

---

## Usage
### General Syntax
```sh
K8Backup <resource> <action> --name "resource_name" --namespace "namespace" --path "kubeconfig_path"
```

### Backup Commands
```sh
# Backup a Pod
K8Backup pod backup --name "pod_name" --namespace "namespace" --path "kubeconfig_path"

# Backup a Deployment
K8Backup deployment backup --name "deployment_name" --namespace "namespace" --path "kubeconfig_path"

# Backup a Persistent Volume Claim (PVC)
K8Backup volume backup --name "pvc_name" --namespace "namespace" --path "kubeconfig_path"
```

### Restore Commands
```sh
# Restore a Pod
K8Backup pod restore --object "backup_object" --path "kubeconfig_path" --name "new_pod_name"

# Restore a Deployment
K8Backup deployment restore --object "backup_object" --path "kubeconfig_path" --name "new_deployment_name"

# Restore a Persistent Volume Claim (PVC)
K8Backup volume restore --object "snapshot_name" --path "kubeconfig_path" --name "new_pvc_name"
```

### Utility Commands
```sh
# List all backup objects
K8Backup list

# Delete a backup object
K8Backup delete --file "backup_filename"
```

---

## Resources
- **Kubernetes Documentation**: https://kubernetes.io/docs/
- **Cobra Documentation**: https://github.com/spf13/cobra

## Contributing
Contributions are welcome! Feel free to submit issues or open pull requests.

