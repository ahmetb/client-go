# client-go Usage Examples

This directory contains examples that cover examples of various use cases
and functionality for Go client library for Kubernetes.

#### How to run these examples

To run these examples, make sure you have `go` tool in your PATH and have a
Kubernetes cluster ready. You can use
[minikube](https://kubernetes.io/docs/getting-started-guides/minikube/) to
create a local Kubernetes cluster on your development machine.

```
$ go get k8s.io/client-go/...
$ cd $HOME/src/k8s.io/client-go/examples
$ go build -o app <EXAMPLE-DIRECTORY>
$ ./app
```

## Table of Contents

- Authenticating
  - [Inside the cluster](./in-cluster)
  - [Outside the cluster](./out-of-cluster)
