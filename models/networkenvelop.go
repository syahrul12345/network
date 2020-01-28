package models

import (
	"encoding/binary"
	"encoding/hex"
)

var (
	mainnetMagic [4]byte = [4]byte{0xf9, 0xbe, 0xb4, 0xd9}
	testnetMagic [4]byte = [4]byte{0x0b, 0x11, 0x09, 0x07}
)

//NetworkEnvelop is a class that handles the network messages
type NetworkEnvelop struct {
	Magic           [4]byte
	Command         [12]byte
	PayloadLength   uint32
	PayloadChecksum [4]byte
	Payload         []byte
	Testnet         bool
}

// ParseNetworkMessage will parse the network message and return the network envelop
func ParseNetworkMessage(networkMessage string, testnet bool) *NetworkEnvelop {
	rawBytes, _ := hex.DecodeString(networkMessage)
	magic := rawBytes[0:4]
	command := rawBytes[4:16]
	payloadLength := rawBytes[16:20]
	payloadChecksum := rawBytes[20:24]
	payload := rawBytes[24:]

	// Create the required buffers
	var magicBuf [4]byte
	var commandBuf [12]byte
	var payloadChecksumBuf [4]byte

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
	copy(commandBuf[:12], command)
	copy(payloadChecksumBuf[:4], payloadChecksum)
	payloadLengthNum := binary.LittleEndian.Uint32(payloadLength)
	return &NetworkEnvelop{
		Magic:           magicBuf,
		Command:         commandBuf,
		PayloadLength:   payloadLengthNum,
		PayloadChecksum: payloadChecksumBuf,
		Payload:         payload,
	}
}

// Serialize the networkenvelop and gives the string
func (networkEnvelop *NetworkEnvelop) Serialize() string {
	// Need to convert the Payload length back to littleendian byte array
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, networkEnvelop.PayloadLength)
	var resBuf []byte
	resBuf = append(resBuf, networkEnvelop.Magic[:]...)
	resBuf = append(resBuf, networkEnvelop.Command[:]...)
	resBuf = append(resBuf, buf...)
	resBuf = append(resBuf, networkEnvelop.PayloadChecksum[:]...)
	resBuf = append(resBuf, networkEnvelop.Payload...)
	return hex.EncodeToString(resBuf)
}
