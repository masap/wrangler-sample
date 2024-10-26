package main

import (
	"github.com/rancher/wrangler-sample/pkg/apis/samplecontroller.k8s.io/v1alpha1"
	controllergen "github.com/rancher/wrangler/v3/pkg/controller-gen"
	"github.com/rancher/wrangler/v3/pkg/controller-gen/args"
)

func main() {
	controllergen.Run(args.Options{
		OutputPackage: "github.com/rancher/wrangler-sample/pkg/generated",
		Boilerplate:   "hack/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"samplecontroller.k8s.io": {
				Types: []interface{}{
					v1alpha1.Foo{},
				},
				GenerateTypes: true,
			},
		},
	})
}
