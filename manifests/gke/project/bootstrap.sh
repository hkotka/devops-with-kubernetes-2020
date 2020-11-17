# Install sealed secrets and apply secrets
kubectl apply -f https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.13.1/controller.yaml
rm secrets/*
sleep 30
kubeseal --scope cluster-wide -o yaml <mysecret-project.json >secrets/postgres-project-secrets.yaml
kubectl apply -k .
