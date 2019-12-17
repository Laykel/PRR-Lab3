/*
Lab 3 - Chang and Roberts with failures
File: network/protocol.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Describe networking values, messages structure for the protocol and provide
encoding and decoding functions for messages
*/
package network

const (
	AnnouncementMessageType = 0
	ResultMessageType       = 1
	AcknowledgeMessageType  = 2
)

type Process struct {
    Address  string `json:"address"`
    Port     int    `json:"port"`
    Aptitude uint8  `json:"aptitude"`
}

// Read constants from parameters file
type Parameters struct {
    NbProcesses    uint8     `json:"nb_of_processes"`
    ProcessAddress []Process `json:"processes"`
}

var Params Parameters

// Message for an election
// Can either be an Announcement message, a Result message or an Ack
type ElectionMessage struct {
	MessageType      uint8
	VisitedProcesses map[uint8]uint8 // Process number - aptitude
	Elect            uint8
	ProcessIdSender  uint8
}

// Compute recipient IP and port and send message
func SendMessage(message ElectionMessage) {
    nextProcess := (message.ProcessIdSender + 1) % Params.NbProcesses

    // Send acknowledge message
    senderIP := Params.ProcessAddress[nextProcess].Address
    senderPort := Params.ProcessAddress[nextProcess].Port

    SendGob(message, senderIP, senderPort)
}
