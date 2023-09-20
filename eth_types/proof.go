package eth_types

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

// Result Object
type QueryRespoonseData struct {
	StorageData      json.RawMessage `json:"storageData"`
	Slot             string          `json:"slot"`
	StorageProofData StorageProof    `json:"storageProofData"`
	BlockRootHash    common.Hash     `json:"blockRootHash"`
}

// Result Object
type QueryBlockHeaderResponseData struct {
	Root   common.Hash `json:"stateRoot"`
	Height uint64      `json:"height"`
}

type Block struct {
	Root   common.Hash `json:"stateRoot"`
	Height string      `json:"number"`
}

func VerifyAccountProof(rootHash common.Hash, key []byte, accounState []byte, proof [][]byte) (bool, error) {

	proofDB := NewMemDB()
	for _, node := range proof {
		key := crypto.Keccak256(node)
		proofDB.Put(key, node)
	}
	path := crypto.Keccak256(key)

	res, err := trie.VerifyProof(rootHash, path, proofDB)
	if err != nil {
		return false, err
	}
	return bytes.Equal(accounState, res), nil
}

func VerifyProof(rootHash common.Hash, key []byte, value []byte, proof [][]byte) (bool, error) {
	paddedKey := common.LeftPadBytes(key, 32)

	proofDB := NewMemDB()
	for _, node := range proof {
		key := crypto.Keccak256(node)
		proofDB.Put(key, node)
	}
	path := crypto.Keccak256(paddedKey)

	res, err := trie.VerifyProof(rootHash, path, proofDB)
	if err != nil {
		return false, err
	}
	return bytes.Equal(value, res), nil
}

func VerifyEthStorageProof(proof *StorageResult, storageHash common.Hash) (bool, error) {
	var err error
	var value []byte

	if len(proof.Value) != 0 {
		value, err = rlp.EncodeToBytes(proof.Value)
		if err != nil {
			return false, err
		}
	}
	strssd := common.Bytes2Hex(proof.Key)
	fmt.Println(strssd)
	return VerifyProof(storageHash, proof.Key, value, proof.Proof)
}

func VerifyEIP1186(proof *StorageProof, stateRoot common.Hash) (bool, error) {
	for _, sp := range proof.StorageProof {
		sp := sp
		if ok, err := VerifyEthStorageProof(&sp, proof.StorageHash); !ok {
			return false, err
		}
	}
	return true, nil
}

func VerifyEthAccountProof(proof *StorageProof, rootHash common.Hash) (bool, error) {
	accountState, err := rlp.EncodeToBytes([]interface{}{
		proof.Nonce, proof.Balance.ToInt(), proof.StorageHash, proof.CodeHash,
	})
	if err != nil {
		return false, err
	}

	return VerifyAccountProof(rootHash, proof.Address.Bytes(), accountState, proof.AccountProof)
}

func VerifyMPTForSlotData(proof *StorageProof, rootHash common.Hash) (bool, error) {

	storageVerified, err := VerifyEIP1186(proof, rootHash)
	if err != nil {
		return false, err
	}

	accountVerified, err := VerifyEthAccountProof(proof, rootHash)
	if err != nil {
		return false, err
	}

	return storageVerified && accountVerified, nil
}
