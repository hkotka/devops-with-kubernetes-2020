#!/bin/bash
./build-images.sh
k3d cluster create mycluster --agents 2 -p "30000-32767:30000-32767@server[0]"
docker exec k3d-mycluster-agent-0 mkdir -p /tmp/kube
docker exec k3d-mycluster-agent-0 mkdir -p /tmp/kube2
kubectl apply -f manifests/local/namespaces/
kubectl apply -f manifests/local/volumes/
kubectl apply -f manifests/local/configs/
# Install sealed secrets and apply secrets
kubectl apply -f https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.13.1/controller.yaml
rm manifests/local/secrets/*
sleep 30
kubeseal --scope cluster-wide -o yaml <mysecret-project.json >manifests/local/secrets/postgres-project-secrets.yaml
kubeseal --scope cluster-wide -o yaml <mysecret-mainapp.json >manifests/local/secrets/postgres-mainapp-secrets.yaml
kubectl apply -f manifests/local/secrets/
# Start apps
kubectl apply -f manifests/local/apps
# Install monitoring - Grafana, Prometheus, Loki
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add stable https://charts.helm.sh/stable
helm repo add loki https://grafana.github.io/loki/charts
helm repo update
helm install prometheus-community/kube-prometheus-stack --generate-name --namespace prometheus
helm upgrade --install loki --namespace=loki-stack loki/loki-stack
# Update Traefik node ports to http 30000 and https 30443
kubectl patch svc traefik -n kube-system -p '{"spec": {"ports": [{"port": 443,"nodePort": 30443,"name": "https"},{"port": 80,"nodePort": 30000,"name": "http"}],"type": "LoadBalancer"}}'