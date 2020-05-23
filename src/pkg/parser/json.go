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
	Resources []Resource
}

type resource struct {
	Root      string `json: "Root"`
	Response  json.RawMessage
	Resources json.RawMessage
}

type Response struct {
	Text   string `json:"text"`
	Status int
}

// Read func create Resource struct from config file
func Read(resourcePath string) *[]Resource {
	bytes := read(resourcePath)
	var resources []Resource

	return createResources(bytes, resources)
}

func createResources(bytes []byte, resources []Resource) *[]Resource {

	var _resources []resource

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

		if len(p.Resources) > 0 {
			var rs []Resource
			re := createResources(p.Resources, rs)
			log.Printf("resources: %+v\n", &re)
			log.Printf("nested: %+v", r.Resources)
			r.Resources = *re
		}

		resources = append(resources, r)
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
