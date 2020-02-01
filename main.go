package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	// versionMessage := &models.VersionMessage{
	// 	Command:      []byte("version"),
	// 	Version:      70015,
	// 	Serivces:     0,
	// 	Timestamp:    0,
	// 	ReceiverPort: 8333,
	// 	SenderPort:   8333,
	// 	Nonce:        1,
	// 	UserAgent:    []byte("/programmingbitcoin:0.1/"),
	// 	LatestBlock:  0,
	// 	Relay:        false,
	// }
	// networkEnvelop := models.CreateNetworkEnvelop(versionMessage.Command, versionMessage.Serialize(), true)

	// // Lets dial here
	// address := "testnet.programmingbitcoin.com:18333"
	// conn, err := net.Dial("tcp", address)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer conn.Close()
	// _, err = conn.Write([]byte(networkEnvelop.Serialize()))
	// if err != nil {
	// 	fmt.Println("error writing..")
	// 	fmt.Println(err)
	// }
	// result, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	fmt.Println("error reading...")
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)
	conn, err := net.Dial("tcp", "testnet.programmingbitcoin.com:18333")
	// conn, err := net.Dial("tcp", service)
	checkError(err)

	defer conn.Close()

	_, err = conn.Write([]byte("getheaders / HTTP/2.0\r\n\r\n"))
	checkError(err)
	fmt.Println("waiting for read..")
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
