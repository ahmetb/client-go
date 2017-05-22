/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"flag"
	"log"
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
	if *kubeconfig == "" {
		panic("-kubeconfig not specified")
	}

	d := yaml.NewYAMLOrJSONDecoder(os.Stdin, 1024*32)
	var ns apiv1.Namespace
	if err := d.Decode(&ns); err != nil {
		panic(err)
	}
	log.Printf("parsed Namespace %q (%d labels)", ns.Name, len(ns.GetObjectMeta().GetLabels()))

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	nsClient := clientset.CoreV1().Namespaces()

	if _, err = nsClient.Create(&ns); errors.IsAlreadyExists(err) {
		if _, err = nsClient.Update(&ns); err != nil {
			panic(err)
		}
		log.Printf("updated namespace %q", ns.Name)
	} else if err != nil {
		panic(err)
	} else {
		log.Printf("created namespace %q", ns.Name)
	}
}
