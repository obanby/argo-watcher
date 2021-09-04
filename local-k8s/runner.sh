#!/bin/bash
CLUSTER_NAME=k8s-argo-cluster
# TODO add flag -d to delete and stop
kind delete cluster --name $CLUSTER_NAME
kind create cluster --config ./multi-node-config.yml --name $CLUSTER_NAME
