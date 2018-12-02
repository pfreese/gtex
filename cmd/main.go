package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pfreese/gtex/pkg/ioutils"
)

// The .json file constaining settings
const jsonF string = "../configs/gtex.json"

func main() {

	// Make a 'logs' directory in which to write log files
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logDir := filepath.Join(pwd, "logs")
	ioutils.CreateDirIfNotExist(logDir)

	// Set up the log file
	// Mark the timestamp in the log file name: month_day_year_hour.minute.second
	const layout = "01_02_2006_03.04.05"

	tm := time.Now().Format(layout)
	logBasename := fmt.Sprintf("log.%s.txt", tm)
	logFile, err := os.OpenFile(filepath.Join(logDir, logBasename),
		os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Print("Logging to a file in Go!")
	// test case
	log.Println("check to make sure it works")

	// Parse the .json file
	configuration, nil := ioutils.ParseJSON(jsonF)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure all required data & metadata files exist locally
	ioutils.EnsureFilesDownloaded(configuration)

	// Finished successfully!
	fmt.Sprintf("FINISHED! See: %s", logFile)
}

