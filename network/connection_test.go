/*
Lab 3 - Chang and Roberts with failures
File: network/connection_test.go
Authors: Jael Dubey, Luc Wachter
Go version: 1.13.4 (linux/amd64)

Unit tests for the connection file.
*/
package network

import (
	"encoding/gob"
	"net"
	"reflect"
	"testing"
)

// Test networking constants
const (
	address = "127.0.0.1"
	port    = 9706
)

// Channel for received messages
var ch chan ElectionMessage

// Run "server" before running tests
func init() {
	ch = make(chan ElectionMessage)

	go Listen(address, port, ch)
}

// Test that the "server" can be run
func TestListenRun(t *testing.T) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(address),
		Port: port + 1,
	})

	if err != nil {
		t.Error("Error during connection to server: ", err)
	}

	defer conn.Close()
}

// Test the Listen function
func TestListen(t *testing.T) {
	want := ElectionMessage{
		MessageType:      0,
		Elect:            0,
		VisitedProcesses: map[uint8]uint8{2: 3, 3: 4, 4: 1},
	}

	t.Run("Reading simple message should work", func(t *testing.T) {
        // ---------------- Mock Send function
		// Send message to server
		conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
			IP:   net.ParseIP(address),
			Port: port,
		})
		if err != nil || conn == nil {
			t.Error("Error connecting to server: ", err)
			return
		}
		defer conn.Close()

		// Encode message as Gob and send it
		encoder := gob.NewEncoder(conn)
		err = encoder.Encode(want)
		if err != nil {
			t.Error("Error writing payload: ", err)
		}

		// ---------------- Test if server got what I sent
		// Get message from network
		got := <-ch

		// Compare result with wanted result
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Listen() got %v, wanted %v.", got, want)
		}
	})
}

// Test Send function
func TestSendGob(t *testing.T) {
	tests := []struct {
		name    string
		message ElectionMessage
		address string
		port    int
	}{
		{
			"Sending simple message should work",
			ElectionMessage{
				MessageType:      0,
				Elect:            0,
				VisitedProcesses: map[uint8]uint8{2: 3, 3: 4, 4: 1},
			},
			address,
			port,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            // Send message to server
			SendGob(tt.message, tt.address, tt.port)
		})

		// Test if server got what I sent
		got := <-ch

		// Compare result with wanted result
		if !reflect.DeepEqual(got, tt.message) {
			t.Errorf("Listen() got %v, wanted %v.", got, tt.message)
		}
	}
}
