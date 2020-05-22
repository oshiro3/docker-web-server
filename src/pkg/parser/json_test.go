package parser

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonReader(t *testing.T) {
	testJSONStr := []byte(`[{"Root": "testRoot1", "Response": "I'm Response"}]`)
	tmpf, _ := ioutil.TempFile("", "exam")
	defer os.Remove(tmpf.Name())

	if _, err := tmpf.Write(testJSONStr); err != nil {
		log.Fatal(err)
	}

	res := Read(tmpf.Name())
	assert.Equal(t, (*res)[0].Root, "testRoot1")
	assert.Equal(t, (*res)[0].Response, "I'm Response")
}
