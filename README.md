# ActiveMQ Artemis Operator

## Overview

This ActiveMQ Artemis operator is an implementation done using the [operator-sdk][operator_sdk] tools and APIs. The SDK CLI `operator-sdk` generates the project layout and controls the development life cycle. In addition, this implementation replaces the use of [client-go][client_go] with the SDK APIs to watch, query, and mutate Kubernetes resources.

## Quick Start

The quick start guide walks through the process of building the ActiveMQ Artemis operator image using the SDK CLI, setting up the RBAC, deploying operators, and creating an ActiveMQ Artemis cluster.

### Prerequisites

- [dep][dep_tool] version v0.4.1+.
- [go][go_tool] version v1.10+.
- [docker][docker_tool] version 17.03+.
- [kubectl][kubectl_tool] version v1.9.0+.
- Access to a kubernetes v.1.9.0+ cluster.

**Note**: This guide uses [minikube][minikube_tool] version v0.25.0+ as the local kubernetes cluster and quay.io for the public registry.

### Install the Operator SDK CLI

First, checkout and install the operator-sdk CLI:

```sh
$ mkdir -p $GOPATH/src/github.com/operator-framework/operator-sdk
$ cd $GOPATH/src/github.com/operator-framework/operator-sdk
$ git checkout tags/v0.0.5
$ dep ensure
$ go install github.com/operator-framework/operator-sdk/commands/operator-sdk
```

### Initial Setup

Checkout this ActiveMQ Artemis Operator repository:

```sh
$ mkdir -p $GOPATH/src/github.com/vromero
$ cd $GOPATH/src/github.com/vromero
$ git clone https://github.com/vromero/activemq-artemis-operator.git
$ cd vromero/activemq-artemis-operator
```

Vendor the dependencies:

```sh
$ dep ensure
```

### Build and run the operator

Build the ActiveMQ Artemis operator image and push it to a public registry:

```sh
$ export IMAGE=vromero/activemq-artemis-operator:latest
$ operator-sdk build $IMAGE
$ docker push $IMAGE
```

Setup RBAC for the ActiveMQ Artemis operator:

```sh
$ kubectl create -f deploy/rbac.yaml
```

Deploy the ActiveMQ Artemis operator:

```sh
$ kubectl create -f deploy/operator.yaml
```
### Deploying a Vault cluster

Create an ActiveMQ Artemis:

```sh
$ kubectl create -f deploy/cr.yaml
```

Verify that the Activemq Artemis cluster is up:

```sh
$ kubectl get pods 
NAME                       READY     STATUS    RESTARTS   AGE
example-654658f5fc-2wdlq   1/2       Running   0          1m
example-654658f5fc-7ztzf   1/2       Running   0          1m
```

### ActiveMQ Artemis Guide

Once the cluster is up, see the [ActiveMQ Artemis Usage Guide][guide] on how to manage and use the cluster.

[client_go]:https://github.com/kubernetes/client-go
[operator_sdk]:https://github.com/operator-framework/operator-sdk
[dep_tool]:https://golang.github.io/dep/docs/installation.html
[go_tool]:https://golang.org/dl/
[docker_tool]:https://docs.docker.com/install/
[kubectl_tool]:https://kubernetes.io/docs/tasks/tools/install-kubectl/
[minikube_tool]:https://github.com/kubernetes/minikube#installation
[guide]:/GUIDE.md
