package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Resource is rooting definition
type Resource struct {
	Root string
	Response
}

type resource struct {
	Root     string `json: "Root"`
	Response json.RawMessage
}

type Response struct {
	Text   string `json:"text"`
	Status int
}

// Read func create Resource struct from config file
func Read(resourcePath string) *[]Resource {
	var bytes = read(resourcePath)
	var _resources []resource
	var resources []Resource
	err := json.Unmarshal(bytes, &_resources)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range _resources {
		var response Response
		err := json.Unmarshal(p.Response, &response)
		if err != nil {
			log.Fatal(err)
		}

		r := Resource{
			Root:     p.Root,
			Response: response,
		}

		resources = append(resources, r)
	}
	log.Printf("Read: %+v\n", resources)
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
