package main

import (
	"net/http"

	"pkg/parser"
)

func main() {
	parser.Read("./config/root.json")
	http.HandleFunc("/", HelloHandler)
	// http.ListenAndServe(":8888", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}
