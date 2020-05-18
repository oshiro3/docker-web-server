package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8888", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}
