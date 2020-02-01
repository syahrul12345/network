package models

import (
	"encoding/binary"
	"testing"
)

func TestParseVersionMessage(t *testing.T) {
	buf := []byte{
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
	}
	v := &VersionMessage{
		Version:      70015,
		Serivces:     0,
		Timestamp:    0,
		ReceiverPort: 8333,
		SenderPort:   8333,
		Nonce:        binary.BigEndian.Uint64(buf),
		UserAgent:    []byte("/programmingbitcoin:0.1/"),
		LatestBlock:  0,
		Relay:        false,
	}
	get := v.Serialize()
	want := "7f11010000000000000000000000000000000000000000000000000000000000000000000000ffff00000000208d000000000000000000000000000000000000ffff00000000208d0000000000000000182f70726f6772616d6d696e67626974636f696e3a302e312f0000000000"
	if get != want {
		t.Errorf("Expected the version message to be %s but got %s", want, get)
	}
}
