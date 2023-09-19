package eth_types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// QuantityBytes marshals/unmarshals as a JSON string in hex with 0x prefix encoded
// as a QUANTITY.  The empty slice marshals as "0x0".
type QuantityBytes []byte

// MarshalText implements encoding.TextMarshaler
func (q QuantityBytes) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("0x%v",
		new(big.Int).SetBytes(q).Text(16))), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (q *QuantityBytes) UnmarshalText(input []byte) error {
	input = bytes.TrimPrefix(input, []byte("0x"))
	v, ok := new(big.Int).SetString(string(input), 16)
	if !ok {
		return fmt.Errorf("invalid hex input")
	}
	*q = v.Bytes()
	return nil
}

// SliceData marshals/unmarshals as a JSON vector of strings with in hex with
// 0x prefix.
type SliceData [][]byte

// MarshalJSON  implements encoding.TextMarshaler
func (s SliceData) MarshalJSON() ([]byte, error) {
	bs := make([]hexutil.Bytes, len(s))
	for i, b := range s {
		bs[i] = b
	}
	return json.Marshal(bs)
}

// UnmarshalJSON implements encoding.TextUnmarshaler.
func (s *SliceData) UnmarshalJSON(data []byte) error {
	var bs []hexutil.Bytes
	if err := json.Unmarshal(data, &bs); err != nil {
		return err
	}
	*s = make([][]byte, len(bs))
	for i, b := range bs {
		(*s)[i] = b
	}
	return nil
}

// StorageProof allows unmarshaling the object returned by `eth_getProof`.
type StorageProof struct {
	Height       *big.Int        `json:"height"`
	Address      common.Address  `json:"address"`
	Balance      *hexutil.Big    `json:"balance"`
	CodeHash     common.Hash     `json:"codeHash"`
	Nonce        hexutil.Uint64  `json:"nonce"`
	StateRoot    common.Hash     `json:"stateRoot"`
	StorageHash  common.Hash     `json:"storageHash"`
	AccountProof SliceData       `json:"accountProof"`
	StorageProof []StorageResult `json:"storageProof"`
}

// StorageResult is an object from StorageProof that contains a proof of
// storage.
type StorageResult struct {
	Key   QuantityBytes `json:"key"`
	Value QuantityBytes `json:"value"`
	Proof SliceData     `json:"proof"`
}

// MemDB is an ethdb.KeyValueReader implementation which is not thread safe and
// assumes that all keys are common.Hash.
type MemDB struct {
	kvs map[common.Hash][]byte
}

// NewMemDB creates a new empty MemDB
func NewMemDB() *MemDB {
	return &MemDB{
		kvs: make(map[common.Hash][]byte),
	}
}

// Has returns true if the MemBD contains the key
func (m *MemDB) Has(key []byte) (bool, error) {
	h := common.BytesToHash(key)
	_, ok := m.kvs[h]
	return ok, nil
}

// Get returns the value of the key, or nil if it's not found
func (m *MemDB) Get(key []byte) ([]byte, error) {
	h := common.BytesToHash(key)
	value, ok := m.kvs[h]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}
	return value, nil
}

// Put sets or updates the value at key
func (m *MemDB) Put(key []byte, value []byte) {
	h := common.BytesToHash(key)
	m.kvs[h] = value
}
