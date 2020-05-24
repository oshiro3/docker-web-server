package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonReader(t *testing.T) {
	// testJSONStr := []byte(`[{"Root": "testRoot1", "Response":[{ "Text": "I'm Response", "Status": 200}]]`)
	testJSONStr := []byte(`
[
  {
    "Root": "testRoot1",
    "Response": [
      {
        "method": "GET",
        "text": "I'm Response",
        "status": 500
      },
      {
        "method": "POST",
        "text": "Receive POST req",
        "status": 301
      }
    ],
    "Children": [
      {
        "Root": "TestRootChirdren1",
        "Response": [
          {
            "text": "Children",
            "status": 404
          }
        ]
      }
    ]
  }
]`)
	tmpf, _ := ioutil.TempFile("", "exam")
	defer os.Remove(tmpf.Name())

	if _, err := tmpf.Write(testJSONStr); err != nil {
		log.Fatal(err)
	}

	res := Read(tmpf.Name())
	fmt.Printf("%+v\n", res)
	assert.Equal(t, (*res)[0].Root, "testRoot1")
	assert.Equal(t, (*res)[0].Response[0].Text, "I'm Response")
	assert.Equal(t, (*res)[0].Response[0].Status, 500)
	assert.Equal(t, (*res)[0].Response[1].Text, "Receive POST req")
	assert.Equal(t, (*res)[0].Response[1].Status, 301)
	assert.Equal(t, (*res)[0].Children[0].Root, "TestRootChirdren1")
}
