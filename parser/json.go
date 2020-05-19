package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Resource struct {
	Root     string
	Response string
}

func Read() *[]Resource {
	var bytes = read()
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

func read() []byte {
	bytes, err := ioutil.ReadFile("./root.json")
	if err != nil {
		panic(err)
	}
	return bytes
}

func main() {
	Read()
}
