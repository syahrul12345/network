package models

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	message := "f9beb4d976657261636b000000000000000000005df6e0e2"
	envelop := ParseNetworkMessage(message, false)
	command := string(envelop.Command[:])
	command = strings.Trim(command, "\x00")
	payload := envelop.Payload
	payloadString := hex.EncodeToString(payload)
	if command != "verack" {
		t.Errorf("Expected the comamnd string to be: %s but got %s", "verack", command)
	}
	if payloadString != "" {
		t.Errorf("Expected the payload to be %v but got %v", []byte{}, payload)
	}
	message = "f9beb4d976657273696f6e0000000000650000005f1a69d2721101000100000000000000bc8f5e5400000000010000000000000000000000000000000000ffffc61b6409208d010000000000000000000000000000000000ffffcb0071c0208d128035cbc97953f80f2f5361746f7368693a302e392e332fcf05050001"
	envelop = ParseNetworkMessage(message, false)
	command = string(envelop.Command[:])
	command = strings.Trim(command, "\x00")
	payload = envelop.Payload
	payloadString = hex.EncodeToString(payload)
	if command != "version" {
		t.Errorf("Expected the command string to be %s but got %s", "version", command)
	}
	if payloadString != message[48:] {
		t.Errorf("Expected the payload to be %s but got %s", message[48:], payloadString)
	}

}

func TestSerialize(t *testing.T) {
	message := "f9beb4d976657261636b000000000000000000005df6e0e2"
	envelop := ParseNetworkMessage(message, false)
	get := envelop.Serialize()
	if get != message {
		t.Errorf("Expected the envelop to serialize into %s but got %s", message, get)
	}
	message = "f9beb4d976657273696f6e0000000000650000005f1a69d2721101000100000000000000bc8f5e5400000000010000000000000000000000000000000000ffffc61b6409208d010000000000000000000000000000000000ffffcb0071c0208d128035cbc97953f80f2f5361746f7368693a302e392e332fcf05050001"
	envelop = ParseNetworkMessage(message, false)
	get = envelop.Serialize()
	if get != message {
		t.Errorf("Expected the envelop to serialize into %s but got %s", message, get)
	}
}
