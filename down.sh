#!/bin/sh

kubectl delete -f deploy/operator.yaml
kubectl delete -f deploy/rbac.yaml
kubectl delete -f deploy/crd.yaml
kubectl delete -f deploy/cr.yaml
