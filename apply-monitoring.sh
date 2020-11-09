#!/bin/bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add stable https://charts.helm.sh/stable
helm repo add loki https://grafana.github.io/loki/charts
helm repo update
helm install prometheus-community/kube-prometheus-stack --generate-name --namespace prometheus
helm upgrade --install loki --namespace=loki-stack loki/loki-stack