package network

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	// Describe test cases
	tests := []struct {
		name    string
		message interface{}
		want    []byte
	}{
		{
			"Test MessageCS message encoding",
			MessageCS{RequestMessageType, 28, 1},
			[]byte{0, 28, 0, 0, 0, 1},
		},
		{
			"Test MessageCS message encoding",
			MessageCS{ReleaseMessageType, 56, 3},
			[]byte{1, 56, 0, 0, 0, 3},
		},
		{
			"Test SetVariable message encoding",
			SetVariable{SetValueMessageType, 456},
			[]byte{2, 0, 0, 1, 200},
		},
	}

	// Run the test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Encode(test.message)

			// Compare result with wanted result
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("encode() got %v, want %v", got, test.want)
			}
		})
	}
}

func TestDecodeMessage(t *testing.T) {
	// Describe test cases
	tests := []struct {
		name   string
		buffer []byte
		want   MessageCS
	}{
		{
			"Test decoding request message",
			[]byte{0, 28, 0, 0, 0, 12},
			MessageCS{RequestMessageType, 28, 12},
		},
		{
			"Test decoding release message",
			[]byte{1, 34, 0, 0, 0, 12},
			MessageCS{ReleaseMessageType, 34, 12},
		},
	}

	// Run test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := DecodeMessage(test.buffer)

			// Compare result with wanted result
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("DecodeMessage() got %v, want %v", got, test.want)
			}
		})
	}
}

func TestDecodeSetVariable(t *testing.T) {
	// Describe test cases
	tests := []struct {
		name   string
		buffer []byte
		want   SetVariable
	}{
		{
			"Test decoding SetVariable message",
			[]byte{2, 0, 0, 0, 12},
			SetVariable{SetValueMessageType, 12},
		},
	}

	// Run test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := DecodeSetVariable(test.buffer)

			// Compare result with wanted result
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("DecodeMessage() got %v, want %v", got, test.want)
			}
		})
	}
}
