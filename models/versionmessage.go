package models

import "encoding/binary"

import "network/utils"

import "encoding/hex"

// VersionMessage is the payload in the network message, if the command == Version
type VersionMessage struct {
	Command          []byte
	Version          uint32
	Serivces         uint32
	Timestamp        uint64
	ReceiverServices uint32
	ReceiverIP       uint32
	ReceiverPort     uint16
	SenderServices   uint32
	SenderIP         uint32
	SenderPort       uint16
	Nonce            uint64
	UserAgent        []byte
	Height           []byte
	Relay            []byte
}

// Serialize will serialize a version message
func (versionMessage *VersionMessage) Serialize() string {
	versionBuf := make([]byte, 4)
	serviceBuf := make([]byte, 4)
	timeStampBuf := make([]byte, 8)

	receiverServiceBuf := make([]byte, 4)
	receiverIPBuf := make([]byte, 4)
	receiverPortBuf := make([]byte, 2)

	senderServiceBuf := make([]byte, 4)
	senderIPBuf := make([]byte, 4)
	senderPortBuf := make([]byte, 2)

	nonceBuf := make([]byte, 8)

	binary.LittleEndian.PutUint32(versionBuf, versionMessage.Version)
	binary.LittleEndian.PutUint32(serviceBuf, versionMessage.Serivces)
	binary.LittleEndian.PutUint64(timeStampBuf, versionMessage.Timestamp)

	binary.LittleEndian.PutUint32(receiverServiceBuf, versionMessage.ReceiverServices)
	// 10 0x00 bytes and 1 0xff bytes then 4 bytes of receiver IP
	binary.BigEndian.PutUint32(receiverIPBuf, versionMessage.ReceiverIP)
	receiverIPBufFull := []byte{
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0xff,
		receiverIPBuf[0],
		receiverIPBuf[1],
		receiverIPBuf[2],
		receiverIPBuf[3],
	}
	binary.BigEndian.PutUint16(receiverPortBuf, versionMessage.ReceiverPort)

	binary.LittleEndian.PutUint32(senderServiceBuf, versionMessage.SenderServices)
	// 10 0x00 bytes and 1 0xff bytes then 4 bytes of receiver IP
	binary.BigEndian.PutUint32(senderIPBuf, versionMessage.SenderIP)
	senderIPBufFull := []byte{
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0xff,
		senderIPBuf[0],
		senderIPBuf[1],
		senderIPBuf[2],
		senderIPBuf[3],
	}
	binary.BigEndian.PutUint16(senderPortBuf, versionMessage.SenderPort)
	binary.BigEndian.PutUint64(nonceBuf, versionMessage.Nonce)

	var result []byte
	result = append(result, versionBuf...)
	result = append(result, serviceBuf...)
	result = append(result, timeStampBuf...)
	result = append(result, receiverServiceBuf...)
	result = append(result, receiverIPBufFull...)
	result = append(result, receiverPortBuf...)
	result = append(result, senderServiceBuf...)
	result = append(result, senderIPBufFull...)
	result = append(result, senderPortBuf...)
	result = append(result, nonceBuf...)
	length := utils.EncodeVarInt(uint64(len(versionMessage.UserAgent)))
	result = append(result, length...)
	result = append(result, versionMessage.UserAgent...)
	result = append(result, versionMessage.Height...)
	return hex.EncodeToString(result)
}
