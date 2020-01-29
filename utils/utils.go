package utils

import (
	"encoding/binary"
	"math/big"
	"strconv"
)

func EncodeVarInt(input uint64) []byte {
	// Some of the nubmers huge, use bigint
	someNumber := strconv.FormatUint(input, 10)
	bigInt, _ := big.NewInt(0).SetString(someNumber, 10)
	check1, _ := big.NewInt(0).SetString("fd", 16)
	check2, _ := big.NewInt(0).SetString("10000", 16)
	check3, _ := big.NewInt(0).SetString("100000000", 16)
	check4, _ := big.NewInt(0).SetString("10000000000000000", 16)
	if bigInt.Cmp(check1) < 0 {
		// If it's less than 256, it can be fit in one byte but
		return bigInt.Bytes()
	} else if bigInt.Cmp(check2) < 0 {
		numberInt, _ := strconv.ParseUint(someNumber, 10, 16)
		numberByte := make([]byte, 2)
		binary.LittleEndian.PutUint16(numberByte, uint16(numberInt))
		return append([]byte{0xfd}, numberByte...)
	} else if bigInt.Cmp(check3) < 0 {
		numberInt, _ := strconv.ParseUint(someNumber, 10, 32)
		numberByte := make([]byte, 4)
		binary.LittleEndian.PutUint32(numberByte, uint32(numberInt))
		return append([]byte{0xfe}, numberByte...)
	} else if bigInt.Cmp(check4) < 0 {
		numberInt, _ := strconv.ParseUint(someNumber, 10, 64)
		numberByte := make([]byte, 8)
		binary.LittleEndian.PutUint64(numberByte, numberInt)
		return append([]byte{0xff}, numberByte...)

	}
	return []byte{}
}
