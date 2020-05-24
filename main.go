package main

import (
	"log"
	"net/http"
	"path/filepath"

	"pkg/build"
	"pkg/parser"
)

func main() {
	configPath, err := filepath.Abs("./config/root.json")
	if err != nil {
		log.Fatalln(err)
	}
	resources := parser.Read(configPath)
	http.HandleFunc("/", helloHandler)
	//	fmt.Println(resources)
	build.Build(resources)
	http.ListenAndServe(":8888", nil)
}

// TODO: This function return rooting list
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World! HOME\n"))
}
