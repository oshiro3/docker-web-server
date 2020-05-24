package build

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"pkg/parser"
)

// Build Resources
func Build(r *[]parser.Resource) {
	build(*r, "")
}

func build(rs []parser.Resource, s string) {
	for _, r := range rs {
		addHandler(r, s)
	}
}

func createFunc(res []parser.Response) func(http.ResponseWriter, *http.Request) {
	findByMethod := func(arr []parser.Response, method string) (parser.Response, error) {
		if method == "" {
			method = http.MethodGet
		}
		for _, v := range arr {
			if v.Method == method {
				return v, nil
			}
		}
		return parser.Response{Text: "NotFound", Status: 404}, errors.New("No implemented method")
	}
	return func(w http.ResponseWriter, req *http.Request) {
		v, err := findByMethod(res, req.Method)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Response: %+v\n", v)
		w.WriteHeader(v.Status)
		w.Write([]byte(fmt.Sprintf("%+v\n", v.Text)))
	}
}

func addHandler(r parser.Resource, parent string) {
	rooting := fmt.Sprintf("%s/%s", parent, r.Root)
	// log.Printf("%s\n", rooting)
	// log.Printf("%+v\n", r.Resources)
	if len(r.Children) > 0 {
		build(r.Children, rooting)
	}
	http.HandleFunc(rooting, createFunc(r.Response))
}
