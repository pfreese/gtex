package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pfreese/gtex/pkg/ioutils"
)

func parseJSON() {
	// Parse the json file
	jsonF := "../configs/gtex.json"
	config, err := ioutils.ParseJSON(jsonF)
	if err != nil {
		fmt.Println(fmt.Sprintf("error parsing JSON %s", jsonF))
		fmt.Print(err.Error())
	}

	log.Println(config)
}

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	logDir := filepath.Join(pwd, "logs")
	fmt.Println(logDir)
	ioutils.CreateDirIfNotExist(logDir)

	// Mark the timestamp in the log name: month_day_year_hour.minute.second
	const layout = "01_02_2006_03.04.05"

	tm := time.Now().Format(layout)
	fmt.Println(tm)
	logBasename := fmt.Sprintf("log.%s.txt", tm)
	file, err := os.OpenFile(filepath.Join(logDir, logBasename), os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Print("Logging to a file in Go!")
	//test case
	log.Println("check to make sure it works")

	parseJSON()
}

