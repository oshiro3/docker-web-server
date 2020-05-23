package build

import (
	"fmt"
	"net/http"

	"pkg/parser"
)

// Build Resources
func Build(resources *[]parser.Resource) {
	for _, resource := range *resources {
		addHandler(resource)
	}
}

func createFunc(res parser.Response) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%+v\n", res)))
		fmt.Printf("%+v\n", r)
	}
}

func addHandler(resource parser.Resource) {
	rooting := fmt.Sprintf("/%s", resource.Root)
	http.HandleFunc(rooting, createFunc(resource.Response))
}
