package main

import (
	"../election_algorithm"
	"../network"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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

func checkIfAllSitesAreReady(processId uint8, nbSites uint8, address string, initialPort int) {
	for i := uint8(0); i < nbSites; i++ {
		if i != processId {
			recipientPort := strconv.Itoa(initialPort + int(i))
			recipientAddress := address + ":" + recipientPort

			network.AreYouThere(recipientAddress)
		}
	}
}

func main() {
	network.Params = loadParameters(parametersFile)

	// Create channels to communicate with the Network routine
	action := make(chan network.ElectionMessage)
	election := make(chan uint8)
	getTheChosenOne := make(chan uint8)

	//var nbProcesses uint8
	var processId uint8
	var aptitude uint8
	var theChosenOne uint8

	if len(os.Args) == 2 {
		tmp, _ := strconv.Atoi(os.Args[1])
		processId = uint8(tmp)
	} else {
		processId = 0
	}

	//nbProcesses = Params.NbProcesses
	aptitude = network.Params.ProcessAddress[processId-1].Aptitude

	address := network.Params.ProcessAddress[processId-1].Address
	port := network.Params.ProcessAddress[processId-1].Port

	theChosenOne = processId

	fmt.Println("Wait until all sites are ready...")
	go network.Listen(address, port, action)
	//checkIfAllSitesAreReady(processId, nbProcesses, address, Params.ProcessAddress[0].Port)
	//time.Sleep(30 * time.Second)
	fmt.Println("All sites are ready. Algorithm will start ! ")

	// TODO ChangAndRoberts should not need address and port
	go election_algorithm.ChangAndRoberts(processId, aptitude, election, getTheChosenOne, action)

	// TODO Is it the correct way to launch election?
	election <- 1

	for {
		select {

		// TODO echo the chosen one periodically
		case <-time.After(2 * 1 * time.Second):
			println("salut")

		case theChosenOne = <-getTheChosenOne:
			fmt.Printf("The Chosen One is %d\n", theChosenOne)
		}
	}
}
