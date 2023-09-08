package types

import (
	"fmt"
)

const (
	// ModuleName defines the module name
	ModuleName = "ethstate"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ethstate"
)

var (
	KeyPrefixSlotData  = []byte{0x01}
	KeyPrefixBlockData = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetSlotKey(contractAddress string, slot string) []byte {
	return append(KeyPrefixSlotData, append([]byte(contractAddress), []byte(slot)...)...)
}

func GetBlockKey(blockNumber uint64) []byte {
	return append(KeyPrefixBlockData, []byte(fmt.Sprintf("%v", blockNumber))...)
}
