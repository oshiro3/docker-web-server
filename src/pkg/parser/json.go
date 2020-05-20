package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Resource struct {
	Root     string
	Response string
}

func Read(resource_path string) *[]Resource {
	var bytes = read(resource_path)
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

func read(resource_path string) []byte {
	bytes, err := ioutil.ReadFile(resource_path)
	if err != nil {
		panic(err)
	}
	return bytes
}

func main() {
	Read("../../config/root.json")
}
