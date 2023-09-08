package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"math/big"
)

func main() {
	ethURL := "https://mainnet.infura.io/v3/bf6a9a06927f473095017ecf7cf2d2ce"

	client, err := ethclient.Dial(ethURL)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	contractAddress := "0x4d224452801ACEd8B2F0aebE155379bb5D594381"
	storagePosition := "0x0000000000000000000000000000000000000000000000000000000000000002"
	blockNumber := big.NewInt(18077318)
	hexBlock := fmt.Sprintf("0x%x", blockNumber)
	fmt.Println(hexBlock)

	clientG := gethclient.New(client.Client())
	if err != nil {
		fmt.Println("error", err)
		return
	}

	data, err := client.StorageAt(context.TODO(), common.HexToAddress(contractAddress), common.HexToHash(storagePosition), blockNumber)

	proofs, err := clientG.GetProof(context.Background(), common.HexToAddress(contractAddress), []string{storagePosition}, blockNumber)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	client.BlockByNumber()

	batch := []rpc.BatchElem{
		{
			Method: "eth_getStorageAt",
			Args:   []interface{}{common.HexToAddress(contractAddress), common.HexToHash(storagePosition), fmt.Sprintf("0x%x", blockNumber)},
			Result: &json.RawMessage{},
		},
		{
			Method: "eth_getProof",
			Args:   []interface{}{common.HexToAddress(contractAddress), common.HexToHash(storagePosition), fmt.Sprintf("0x%x", blockNumber)},
			Result: &json.RawMessage{},
		},
		{
			Method: "eth_getBlockByNumber",
			Args:   []interface{}{fmt.Sprintf("0x%x", blockNumber)},
			Result: &json.RawMessage{},
		},
	}

	err = client.Client().BatchCall(batch)
	if err != nil {
		panic(err)
	}
	//var result string
	//err = client.Client().Call(&result, "eth_getProof", []interface{}{common.HexToAddress(contractAddress), []string{storagePosition}, fmt.Sprintf("0x%x", blockNumber)}...)
	//if err != nil {
	//	panic(err)
	//}

	torageRoot, err := client.BlockByNumber(context.Background(), blockNumber)
	fmt.Println(torageRoot)
	//dstr := string(data)
	//dstrp := common.Hex2Bytes(dstr)

	paddedBytes := append([]byte(storagePosition[:2]), common.LeftPadBytes([]byte(storagePosition)[2:], 32)...)
	cpp := string(paddedBytes)
	cpps := common.FromHex(cpp)

	val := data
	dbA := trie.NewDatabase(rawdb.NewMemoryDatabase())
	testtrie, err := trie.New(trie.TrieID(types.EmptyRootHash), dbA)
	err = testtrie.Update(crypto.Keccak256(cpps), val)
	if err != nil {
		panic(err)
	}

	rootHash := testtrie.Hash()

	fmt.Println(rootHash)

	//storageHash := proofs.StorageHash
	//storageProof := proofs.StorageProof[0]
	//
	//key := common.LeftPadBytes([]byte(storageProof.Key), 32)
	//value,err:= rlp.EncodeToBytes(storageProof.Value)
	//
	//proofTrie ;=

	//if storageHash == roo

	fmt.Printf("Data at storage position %s: %s\n", storagePosition, data)

	fmt.Println("client", client, err)
	fmt.Println("proofs", proofs)
}

//proofsfunc main() {
//	ethURL := "https://mainnet.infura.io/v3/bf6a9a06927f473095017ecf7cf2d2ce"
//
//	client, err := ethclient.Dial(ethURL)
//	if err != nil {
//		fmt.Println("error", err)
//		return
//	}
//
//}
