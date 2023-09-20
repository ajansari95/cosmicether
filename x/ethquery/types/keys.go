package types

const (
	// ModuleName defines the module name
	ModuleName = "ethquery"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ethquery"
)

var (
	KeyPrefixEthQuery = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetEthQueryKey(id string) []byte {
	return append(KeyPrefixEthQuery, append([]byte(id))...)
}
