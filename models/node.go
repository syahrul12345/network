package models

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
)

// SimpleNode is a bitcoin node
type SimpleNode struct {
	Testnet bool
	Logging bool
	Conn    net.Conn
}

// CreateSimpleNode
func CreateSimpleNode(host string, port string, testnet bool, logging bool) *SimpleNode {
	addr := host + ":" + port
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Failed to connect to server")
	}
	return &SimpleNode{
		Testnet: testnet,
		Logging: logging,
		Conn:    conn,
	}
}

// SendVersionMessage will cause the simplenode sends a version message
func (simpleNode *SimpleNode) SendVersionMessage(message VersionMessage) {
	envelop := CreateNetworkEnvelop(message.Command, message.Serialize(), simpleNode.Testnet)
	if simpleNode.Logging {
		fmt.Printf("Sending %s\n", envelop.Serialize())
	}
	_, err := simpleNode.Conn.Write([]byte(envelop.Serialize()))
	if err != nil {
		fmt.Println("Failed to write to remote server..")
		fmt.Println(err.Error())
	}
}

// SendVerAck will cause the simplenode sends a version message
func (simpleNode *SimpleNode) SendVerAck(message VerAckMessage) {
	envelop := CreateNetworkEnvelop(message.Command, message.Serialize(), simpleNode.Testnet)
	if simpleNode.Logging {
		fmt.Printf("Sending %s\n", envelop.Serialize())
	}
	_, err := simpleNode.Conn.Write([]byte(envelop.Serialize()))
	if err != nil {
		fmt.Println("Failed to write to remote server..")
		fmt.Println(err.Error())
	}
}

// SendPong will cause the simplenode sends a version message
func (simpleNode *SimpleNode) SendPong(message PongMessage) {
	envelop := CreateNetworkEnvelop(message.Command, message.Serialize(), simpleNode.Testnet)
	if simpleNode.Logging {
		fmt.Printf("Sending %s\n", envelop.Serialize())
	}
	_, err := simpleNode.Conn.Write([]byte(envelop.Serialize()))
	if err != nil {
		fmt.Println("Failed to write to remote server..")
		fmt.Println(err.Error())
	}
}

// Read a message from the connection
func (simpleNode *SimpleNode) Read() *NetworkEnvelop {
	result, err := ioutil.ReadAll(simpleNode.Conn)
	if simpleNode.Logging {
		fmt.Println("receiving a network envelop...")
	}
	if err != nil {
		fmt.Println("Failed to read from the connection")
		return nil
	}
	envelop := ParseNetworkMessage(hex.EncodeToString(result), simpleNode.Testnet)
	return envelop
}

// WaitForMessage will read responses coming from the remote server
func (simpleNode *SimpleNode) WaitForMessage(messages []interface{}) {
	command := ""
	classes := []string{}
	defer simpleNode.Conn.Close()
	// Parse whatever is given and put it in the list
	for _, message := range messages {
		classes = append(classes, string(message.(VerAckMessage).Command))
	}
	for {
		if !stringInSlice(command, classes) {
			res, err := ioutil.ReadAll(simpleNode.Conn)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Failed to read from the connection")
			}
			message := hex.EncodeToString(res)
			envelop := ParseNetworkMessage(message, true)
			command = string(envelop.Command[:])
			if command == "version" {
				simpleNode.SendVerAck(*CreateVerAckMessage())
			} else if command == "ping" {
				simpleNode.SendPong(*CreatePong(envelop.Payload))
			}
		}
		fmt.Println("END")
	}
}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Handshake will do a handshake
func (simpleNode *SimpleNode) Handshake() {
	versionMessage := &VersionMessage{
		Command:      []byte("version"),
		Version:      70015,
		Serivces:     0,
		Timestamp:    0,
		ReceiverPort: 8333,
		SenderPort:   8333,
		Nonce:        1,
		UserAgent:    []byte("/programmingbitcoin:0.1/"),
		LatestBlock:  0,
		Relay:        false,
	}
	simpleNode.SendVersionMessage(*versionMessage)

	simpleNode.WaitForMessage([]interface{}{
		*CreateVerAckMessage(),
	})
}
