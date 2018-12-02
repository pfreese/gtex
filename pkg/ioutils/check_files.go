package ioutils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Data and annotation files were downloaded from V7 of:
// https://gtexportal.org/home/datasets
const (
	annotations string = "GTEx_v7_Annotations_SampleAttributesDS.txt"
	geneMedianTPM string = "GTEx_Analysis_2016-01-15_v7_RNASeQCv1.1.8_gene_median_tpm.gct.gz"
	geneReads string = "GTEx_Analysis_2016-01-15_v7_RNASeQCv1.1.8_gene_reads.gct.gz"
	sampleAnnot string = "GTEx_v7_Annotations_SampleAttributesDS.txt"
)

// DataFiles contains absolute paths of files
type DataFiles struct {
	Annotations	string
	GeneMedianTPM	string
	GeneReads	string
	SampleAnnot	string
}

// Configuration tracks json config attributes
type Configuration struct {
	DataDIR    string
}

// ParseJSON parses a settings .json file and saves fields in a Configuration struct
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

func EnsureFilesDownloaded(config Configuration) DataFiles {
	dataFiles := DataFiles{}

	// Check for annotations
	annotationsF := filepath.Join(config.DataDIR, annotations)
	if _, err := os.Stat(annotationsF); os.IsNotExist(err) {
		panic(fmt.Sprintf("Error does not exist:", annotationsF))
	} else {
		log.Printf("sampleAnnot file exists: %s", annotationsF)
		dataFiles.Annotations = annotationsF
	}

	// Check for geneMedianTPM
	geneMedianTPMF := filepath.Join(config.DataDIR, geneMedianTPM)
	if _, err := os.Stat(annotationsF); os.IsNotExist(err) {
		panic(fmt.Sprintf("Error does not exist:", geneMedianTPMF))
	} else {
		log.Printf("GeneMedianTPM file exists: %s", geneMedianTPMF)
		dataFiles.GeneMedianTPM = geneMedianTPMF
	}

	// Check for geneReads
	geneReadsF := filepath.Join(config.DataDIR, geneReads)
	if _, err := os.Stat(annotationsF); os.IsNotExist(err) {
		panic(fmt.Sprintf("Error does not exist:", geneReadsF))
	} else {
		log.Printf("geneReads file exists: %s", geneReadsF)
		dataFiles.GeneReads = geneMedianTPMF
	}

	// Check for sampleAnnot
	sampleAnnotF := filepath.Join(config.DataDIR, sampleAnnot)
	if _, err := os.Stat(sampleAnnotF); os.IsNotExist(err) {
		panic(fmt.Sprintf("Error does not exist:", sampleAnnotF))
	} else {
		log.Printf("sampleAnnot file exists: %s", sampleAnnotF)
		dataFiles.SampleAnnot = sampleAnnotF
	}

	return dataFiles
}

// CreateDirIfNotExist creates the dir if it does not already exist
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
