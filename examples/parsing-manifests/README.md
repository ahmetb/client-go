# Parsing YAML/JSON manifests

This example demonstrates how to parse YAML or JSON manifest files
describing Kubernetes resources and deploying them using client-go
library.

## Running this example

This sample application parses the YAML/JSON files passed through
standard input (STDIN) into a [Namespace][1] object. Then, it
creates (or updates) the namespace with the given specification in
the manifest file.

First, compile this application:

```
cd parsing-manifests
go build -o ./app
```

Then deploy the contents of `namespace.yaml` by running:

```
./app -kubeconfig=$HOME/.kube/config < namespace.yaml
```

You will see output:

```
parsed namespace "client-go-demo" (0 labels)
updated namespace "client-go-demo"
```

Then inspect the `namespace-updated.json`, which is an updated
version of the existing namespace, in JSON manifest format. To
update existing namespace with this manifest:


```
./app -kubeconfig=$HOME/.kube/config < namespace-updated.json
```

You will see output:

```
parsed namespace "client-go-demo" (1 labels)
updated namespace "client-go-demo"
```

You can run `kubectl get ns client-go-demo -o=yaml` to verify the updated
namespace has the labels specified in the `namespace-updated.json`.

> **NOTE:** The `yaml.YAMLOrJSONDecoder` type does not do any schema validation
> on the given JSON/YAML files. If the manifest file has typos or invalid
> fields, parsing manifest will not return any errors. You should use `kubectl`
> command to apply manifests reliably.

[1]: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
