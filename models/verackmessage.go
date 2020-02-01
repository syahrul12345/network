package models

import "encoding/hex"

// VerAckMessage is the message returned as a response when a version message is sent to a node
type VerAckMessage struct {
	Command []byte
}

// CreateVerAckMessage creates a VerAckMessage
func CreateVerAckMessage() *VerAckMessage {
	return &VerAckMessage{
		Command: []byte("verack"),
	}
}

//Serialize a VerAckMessage
func (verackMessage *VerAckMessage) Serialize() string {
	return ""
}

type PingMessage struct {
	Command []byte
	Nonce   []byte
}

// CreatePing will create a PingMessage
func CreatePing(nonce []byte) *PingMessage {
	return &PingMessage{
		Command: []byte("ping"),
		Nonce:   nonce,
	}
}

type PongMessage struct {
	Command []byte
	Nonce   []byte
}

// CreatePong will create a PongMessage
func CreatePong(nonce []byte) *PongMessage {
	return &PongMessage{
		Command: []byte("pong"),
		Nonce:   nonce,
	}
}

// Serialize a pongMewssage
func (pongmessage *PongMessage) Serialize() string {
	return hex.EncodeToString(pongmessage.Nonce)
}
