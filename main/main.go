/*
Lab 3 - Chang and Roberts with failures
File: main/main.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Main entrypoint for the election algorithm program.
*/
package main

import (
    "../network"
    "encoding/json"
    "log"
    "os"
    "strconv"
)

// Path to json parameters file
const parametersFile = "main/parameters.json"

// Read constants from parameters file
type Parameters struct {
	NbProcesses    uint8     `json:"nb_of_processes"`
	ProcessAddress []Process `json:"processes"`
}

type Process struct {
	Address  string `json:"address"`
	Aptitude uint8  `json:"aptitude"`
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

	// Create channels to communicate with the Network routine
	election := make(chan uint8)
	getTheChosenOne := make(chan uint8)
	announcement := make(chan uint8)
	result := make(chan uint8)

	var nbProcesses uint8
	var processId uint8
	var aptitude uint8
	var state uint8
	var theChosenOne uint8

	if len(os.Args) == 2 {
		tmp, _ := strconv.Atoi(os.Args[1])
		processId = uint8(tmp)
	} else {
		processId = 0
	}

	nbProcesses = Params.NbProcesses
	aptitude = Params.ProcessAddress[processId].Aptitude

	// TODO Launch network go routine

	for {
		select {

		case <-election:
			// TODO Send ANNOUNCEMENT{(processId, aptitude)}
			state = network.AnnouncementMessageType
		case <-getTheChosenOne:

		case <-announcement:
			// TODO Receive ANNOUNCEMENT and remove next line
			var list network.Announce

			_, ok := list.VisitedProcesses[processId]
			if ok {
				// TODO keyOfMax = Bad English?
				var maxApt, keyOfMax uint8
				for k, v := range list.VisitedProcesses {
					if maxApt < v {
						maxApt = v
						keyOfMax = k
					}
				}
				theChosenOne = keyOfMax
				// TODO Send RESULT(theChosenOne, {processId})
				state = network.ResultMessageType
			} else {
				list.VisitedProcesses[processId] = aptitude
				// TODO Send ANNOUNCEMENT(list.VisitedProcesses)
				state = network.AnnouncementMessageType
			}
		case <-result:
			// TODO Receive RESULT and remove next line
			var list network.Result

			ok := list.VisitedProcesses[processId]
			if ok {
				break
			} else if state == network.ResultMessageType && theChosenOne != list.Elect {
				// TODO Send ANNOUNCEMENT({processId, aptitude})
				state = network.AnnouncementMessageType
			} else if state == network.AnnouncementMessageType {
				theChosenOne = list.Elect
				list.VisitedProcesses[processId] = true
				// TODO Send RESULT(theChoseOne, list.VisitedProcesses)
				state = network.ResultMessageType
			}
		}
	}
}
