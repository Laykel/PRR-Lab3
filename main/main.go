/*
Lab 3 - ...
File: main/main.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Main entrypoint for the mutual exclusion program.

The access to the shared variable is guaranteed to be mutually exclusive
thanks to the Carvalho-Roucairol algorithm.

This file contains the central part of the algorithm, receiving requests
from the client, forwarding them to the network manager and calling implementation functions.
*/
package main

import (
	"../network"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Path to json parameters file
const parametersFile = "main/parameters.json"

// Load parameters from json file
func loadParameters(file string) network.Parameters {
	var params network.Parameters

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
	fmt.Println("Test")
}
