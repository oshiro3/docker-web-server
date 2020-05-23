package build

import (
	"fmt"
	"net/http"

	"pkg/parser"
)

// Build Resources
func Build(resources *[]parser.Resource) {
	build(*resources, "")
}

func build(resources []parser.Resource, s string) {
	for _, resource := range resources {
		addHandler(resource, s)
	}
}

func createFunc(res parser.Response) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(res.Status)
		w.Write([]byte(fmt.Sprintf("%+v\n", res.Text)))
	}
}

func addHandler(r parser.Resource, parent string) {
	rooting := fmt.Sprintf("%s/%s", parent, r.Root)
	// log.Printf("%s\n", rooting)
	// log.Printf("%+v\n", r.Resources)
	if len(r.Resources) > 0 {
		build(r.Resources, rooting)
	}
	http.HandleFunc(rooting, createFunc(r.Response))
}
