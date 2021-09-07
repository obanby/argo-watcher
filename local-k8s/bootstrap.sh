#!/bin/bash

if ! command -v kind &> /dev/null
then
    echo "kind could not be found. Please install and add it to PATH"
    exit 1
fi

read -p "Enter cluster name: " CLUSTER_NAME
read -p "Enter your public IP for K8S API to listen on: " API_IP
read -p "Enter port you would like to bind to: " API_PORT

cat <<EOF > kind-$CLUSTER_NAME-config.yml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
  - role: worker
  - role: worker
networking:
  # WARNING: It is _strongly_ recommended that you keep this the default
  # (127.0.0.1) for security reasons. However it is possible to change this.
  apiServerAddress: $API_IP
  # By default the API server listens on a random open port.
  # You may choose a specific port but probably don't need to in most cases.
  # Using a random port makes it easier to spin up multiple clusters.
  apiServerPort: $API_PORT
EOF

kind delete cluster --name $CLUSTER_NAME
kind create cluster --config ./kind-$CLUSTER_NAME-config.yml --name $CLUSTER_NAME
