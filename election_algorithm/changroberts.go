/*
Lab 3 - Chang and Roberts with failures
File: main/changroberts.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Main entrypoint for the election algorithm program.
*/
package election_algorithm

import (
	"../network"
)

func Itob(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

// Main entrypoint for the mutual exclusion program
func ChangAndRoberts(processId uint8,
	aptitude uint8,
	election chan uint8,
	getTheChosenOne chan uint8,
	action chan network.ElectionMessage) {

	theChosenOne := processId
	state := network.ResultMessageType

	for {
		select {

		case msg := <-action:
			switch msg.MessageType {

			case network.AnnouncementMessageType:
				list := msg

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

					message := network.ElectionMessage{
						MessageType:      network.ResultMessageType,
						Elect:            theChosenOne,
						VisitedProcesses: make(map[uint8]uint8),
						ProcessIdSender:  processId,
					}
					message.VisitedProcesses[processId] = 1
					network.SendElectionMessage(message)

					state = network.ResultMessageType
				} else {
					list.VisitedProcesses[processId] = aptitude

					message := network.ElectionMessage{
						MessageType:      network.AnnouncementMessageType,
						Elect:            0,
						VisitedProcesses: list.VisitedProcesses,
						ProcessIdSender:  processId,
					}
					network.SendElectionMessage(message)

					state = network.AnnouncementMessageType
				}

			case network.ResultMessageType:
				list := msg

				ok := list.VisitedProcesses[processId]
				if Itob(int(ok)) {
					break
				} else if state == network.ResultMessageType && theChosenOne != list.Elect {

					message := network.ElectionMessage{
						MessageType:      network.AnnouncementMessageType,
						Elect:            0,
						VisitedProcesses: make(map[uint8]uint8),
						ProcessIdSender:  processId,
					}
					message.VisitedProcesses[processId] = aptitude
					network.SendElectionMessage(message)

					state = network.AnnouncementMessageType
				} else if state == network.AnnouncementMessageType {
					theChosenOne = list.Elect
					list.VisitedProcesses[processId] = 1

					message := network.ElectionMessage{
						MessageType:      network.ResultMessageType,
						Elect:            theChosenOne,
						VisitedProcesses: list.VisitedProcesses,
						ProcessIdSender:  processId,
					}
					network.SendElectionMessage(message)

					state = network.ResultMessageType
				}
			}

		case <-election:
			message := network.ElectionMessage{
				MessageType:      network.AnnouncementMessageType,
				Elect:            0,
				VisitedProcesses: make(map[uint8]uint8),
				ProcessIdSender:  processId,
			}
			message.VisitedProcesses[processId] = aptitude
			network.SendElectionMessage(message)

			state = network.AnnouncementMessageType
		}

		if state == network.ResultMessageType {
			getTheChosenOne <- theChosenOne
		}
	}
}
