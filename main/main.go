/*
Lab 3 - Chang and Roberts with failures
File: main/main.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Main entrypoint for the election algorithm program.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Path to json parameters file
const parametersFile = "main/parameters.json"

// Read constants from parameters file
type Parameters struct {
	InitialPort    uint16 `json:"initial_port"`
	NbProcesses    uint8  `json:"nb_of_processes"`
	ProcessAddress string `json:"process_address1"`
}

var Params Parameters

// Load parameters from json file
func loadParameters(file string) Parameters {
	var params Parameters

	// Read parameters file
	configFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err.Error())
	} else if configFile == nil {
		log.Fatal("Could not open parameters file.")
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&params)

	return params
}

// Main entrypoint for the mutual exclusion program
func main() {
	Params = loadParameters(parametersFile)

	fmt.Println("Test " + Params.ProcessAddress)

	// Run the UDP receiver
	// go Listen()
}
