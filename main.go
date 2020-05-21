package main

import (
	"fmt"
	"net/http"

	"pkg/build"
	"pkg/parser"
)

func main() {
	resources := parser.Read("./config/root.json")
	fmt.Printf("%+v\n", *resources)
	http.HandleFunc("/", helloHandler)
	build.Build(resources)
	http.ListenAndServe(":8888", nil)
}

// TODO: This function return rooting list
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World! HOME\n"))
}
