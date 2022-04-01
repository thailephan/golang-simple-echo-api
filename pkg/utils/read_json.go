package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// https://tutorialedge.net/golang/parsing-json-with-golang/
func ReadJson[O any](path string, out O) (err error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()

	if err != nil {
		return
	}

	fmt.Println("Success read json file")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &out)
	return nil
}