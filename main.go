package main

import (
	"fmt"
	"net/http"

	"pkg/parser"
)

func main() {
	parser.Read("./config/root.json")
	// http.HandleFunc("/", HelloHandler)
	for i := 0; i < 5; i++ {
		addHandler(i)
	}
	http.ListenAndServe(":8888", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World! HOME\n"))
}

func createFunc(i int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello, World! No.%d\n", i)))
	}
}

func addHandler(i int) {
	rooting := fmt.Sprintf("/dummy%d", i)
	http.HandleFunc(rooting, createFunc(i))
}
