#!/bin/sh

operator-sdk build vromero/activemq-artemis-operator
docker push vromero/activemq-artemis-operator

kubectl create -f deploy/rbac.yaml
kubectl create -f deploy/crd.yaml
kubectl create -f deploy/operator.yaml


