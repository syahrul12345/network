package models

import (
	"encoding/binary"
	"fmt"
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
		Version:   70005,
		Serivces:  0,
		Timestamp: 0,
		Nonce:     binary.BigEndian.Uint64(buf),
	}
	fmt.Println(v.Serialize())
}
