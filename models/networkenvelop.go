package models

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

var (
	mainnetMagic [4]byte = [4]byte{0xf9, 0xbe, 0xb4, 0xd9}
	testnetMagic [4]byte = [4]byte{0x0b, 0x11, 0x09, 0x07}
)

//NetworkEnvelop is a class that handles the network messages
type NetworkEnvelop struct {
	Magic   [4]byte
	Command [12]byte
	Payload []byte
}

//CreateNetworkEnvelop will create a NetworkEnvelop
func CreateNetworkEnvelop(command []byte, payload string, testnet bool) *NetworkEnvelop {
	var magic [4]byte
	if testnet {
		magic = testnetMagic
	} else {
		magic = mainnetMagic
	}
	var commandBuf [12]byte
	copy(commandBuf[:][:12], command)
	payloadBuf, _ := hex.DecodeString(payload)

	return &NetworkEnvelop{
		Magic:   magic,
		Command: commandBuf,
		Payload: payloadBuf,
	}
}

// ParseNetworkMessage will parse the serialized NetworkMessage and return the network envelop
func ParseNetworkMessage(networkMessage string, testnet bool) *NetworkEnvelop {
	rawBytes, _ := hex.DecodeString(networkMessage)
	magic := rawBytes[0:4]
	var magicBuf [4]byte
	// Copy into required buffers and return
	copy(magicBuf[:4], magic)

	// Check if testnet or mainnet magic correct
	if testnet {
		if hex.EncodeToString(magicBuf[:]) != hex.EncodeToString(testnetMagic[:]) {
			return nil
		}
	} else {
		if hex.EncodeToString(magicBuf[:]) != hex.EncodeToString(mainnetMagic[:]) {
			return nil
		}
	}
	//Get the command of the network message
	command := rawBytes[4:16]
	var commandBuf [12]byte
	copy(commandBuf[:12], command)

	// Get the payload of the ntwork message
	payloadLength := binary.LittleEndian.Uint32(rawBytes[16:20])
	payload := rawBytes[24 : 24+payloadLength]
	// Get the checksum
	payloadChecksum := hex.EncodeToString(rawBytes[20:24])
	// Calculate the actual checksum
	first := sha256.Sum256(payload)
	second := sha256.Sum256(first[:])
	calculatedChecksum := hex.EncodeToString(second[:][0:4])
	if payloadChecksum != calculatedChecksum {
		fmt.Println("checksum does not match")
		return nil
	}
	return &NetworkEnvelop{
		Magic:   magicBuf,
		Command: commandBuf,
		Payload: payload,
	}

}

// Serialize the networkenvelop and gives the string
func (networkEnvelop *NetworkEnvelop) Serialize() string {
	// Need to convert the Payload length back to littleendian byte array
	magicBuf := make([]byte, 4)
	copy(magicBuf, networkEnvelop.Magic[:])

	var buf []byte
	buf = append(buf, magicBuf...)
	buf = append(buf, networkEnvelop.Command[:]...)
	// Get the byte array of the payload length
	payloadLength := make([]byte, 4)
	binary.LittleEndian.PutUint32(payloadLength, uint32(len(networkEnvelop.Payload)))
	buf = append(buf, payloadLength...)

	first := sha256.Sum256(networkEnvelop.Payload)
	second := sha256.Sum256(first[:])
	calculatedChecksum := second[:][0:4]
	buf = append(buf, calculatedChecksum...)
	buf = append(buf, networkEnvelop.Payload...)
	return hex.EncodeToString(buf)
}
