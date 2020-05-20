package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Resource is rooting definition
type Resource struct {
	Root     string
	Response string
}

// Read func create Resource struct from config file
func Read(resourcePath string) *[]Resource {
	var bytes = read(resourcePath)
	var resources []Resource
	err := json.Unmarshal(bytes, &resources)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, p := range resources {
		fmt.Printf("Root is: %s\n", p.Root)
		fmt.Printf("Response is: %s\n", p.Response)
	}
	return &resources
}

func read(resourcePath string) []byte {
	bytes, err := ioutil.ReadFile(resourcePath)
	if err != nil {
		panic(err)
	}
	return bytes
}

func main() {
	Read("../../config/root.json")
}
