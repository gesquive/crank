package passgen

import "bytes"
import "crypto/rand"
import "encoding/binary"

func getRandomUInt16Array(length int) []uint16 {
	sz := 2
	rands := make([]byte, length*sz)
	_, err := rand.Read(rands)
	if err != nil {
		return []uint16{}
	}

	ints := make([]uint16, length)
	var store uint16
	for i := range ints {
		idx := sz * i
		binary.Read(bytes.NewBuffer(rands[idx:idx+sz]), binary.LittleEndian, &store)
		ints[i] = store
	}
	return ints
}

func getRandomUInt32Array(length int) []uint32 {
	sz := 4
	rands := make([]byte, length*sz)
	_, err := rand.Read(rands)
	if err != nil {
		return []uint32{}
	}

	ints := make([]uint32, length)
	var store uint32
	for i := range ints {
		idx := sz * i
		binary.Read(bytes.NewBuffer(rands[idx:idx+sz]), binary.LittleEndian, &store)
		ints[i] = store
	}
	return ints
}

func getRandomUInt64Array(length int) []uint64 {
	sz := 8
	rands := make([]byte, length*sz)
	_, err := rand.Read(rands)
	if err != nil {
		return []uint64{}
	}

	ints := make([]uint64, length)
	var store uint64
	for i := range ints {
		idx := sz * i
		binary.Read(bytes.NewBuffer(rands[idx:idx+sz]), binary.LittleEndian, &store)
		ints[i] = store
	}
	return ints
}
