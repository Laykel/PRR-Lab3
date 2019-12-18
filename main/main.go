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

func checkIfAllSitesAreReady(processId uint8) {
	for i := uint8(0); i < network.Params.NbProcesses; i++ {
		if i != processId {
			recipientPort := strconv.Itoa(network.Params.ProcessAddress[0].Port + int(i))
			recipientAddress := network.Params.ProcessAddress[i].Address + ":" + recipientPort

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
		log.Fatal("Wrong number of arguments. Please pass only the process id.")
	}

	//nbProcesses = network.Params.NbProcesses
	aptitude = network.Params.ProcessAddress[processId].Aptitude

	address := network.Params.ProcessAddress[processId].Address
	port := network.Params.ProcessAddress[processId].Port

	theChosenOne = processId

	network.EchoHaveResponse = true

	fmt.Println("Wait until all sites are ready...")
	go network.Listen(address, port, action)
	go network.ListenTCP(address, port)
	checkIfAllSitesAreReady(processId)
	fmt.Println("All sites are ready. Algorithm will start!")

	go election_algorithm.ChangAndRoberts(processId, aptitude, election, getTheChosenOne, action)

	// TODO Is it the correct way to launch election?
	election <- 1

	for {
		select {

		case <-time.After(2 * 1 * time.Second):
			if theChosenOne != processId {
				if network.EchoHaveResponse {
					log.Printf("Send echo to %d\n", theChosenOne)
					network.SendMeta(network.ElectionMessage{
						MessageType: network.EchoMessageType, ProcessIdSender: processId},
						theChosenOne)

				} else {
					log.Printf("Echo doesn't receive response. Launch election !")
					election <- 1
					theChosenOne = processId
				}
			}

		case theChosenOne = <-getTheChosenOne:
			fmt.Printf("The Chosen One is %d\n", theChosenOne)
		}
	}
}
