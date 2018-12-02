package ioutils

import (
	"encoding/json"
	"os"
)

// Data and annotation files were downloaded from:
//
const (
	sampleAnnot string = "GTEx_v7_Annotations_SampleAttributesDS.txt"
)

// Configuration tracks json config attributes
type Configuration struct {
	DataDIR    string
}


func ParseJSON(jsonF string) (Configuration, error) {
	file, _ := os.Open(jsonF)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, nil
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
