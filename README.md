# Microservice Project: Dynamic Route Setup

This document outlines the process of deploying a Helm chart to a Kubernetes cluster, followed by integrating it with the Linkerd network for service routing based on selected headers.

## Required Tools and Libraries

- Helm: A package manager for Kubernetes, used for deploying our chart.
- Linkerd: A service mesh for Kubernetes, used for network functions like routing and load balancing.
- Go: We'll use a Go script to control the flow of the deployment process.

## Procedure

1. [ ] Enable Kubernetes on Docker Desktop: We'll use this for our demonstration.
2. [ ] Install Linkerd: This will set up the service mesh on our Kubernetes cluster.
3. [ ] Create a Go script to deploy the service: This script will automate the deployment process.