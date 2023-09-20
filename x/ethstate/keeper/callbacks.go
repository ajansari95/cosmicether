package keeper

import (
	"encoding/json"

	"github.com/ajansari95/cosmicether/eth_types"
	ethquerymoduletypes "github.com/ajansari95/cosmicether/x/ethquery/types"
	"github.com/ajansari95/cosmicether/x/ethstate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Callback func(*Keeper, sdk.Context, []byte, ethquerymoduletypes.EthQuery) error

type Callbacks struct {
	k         *Keeper
	callbacks map[string]Callback
}

var _ ethquerymoduletypes.QueryCallbacks = Callbacks{}

func (k *Keeper) CallbackHandler() Callbacks {
	return Callbacks{
		k:         k,
		callbacks: make(map[string]Callback),
	}
}

// Call calls callback handler
func (c Callbacks) Call(ctx sdk.Context, id string, args []byte, query ethquerymoduletypes.EthQuery) error {
	return c.callbacks[id](c.k, ctx, args, query)
}

func (c Callbacks) Has(id string) bool {
	_, found := c.callbacks[id]
	return found
}

func (c Callbacks) AddCallback(id string, fn interface{}) ethquerymoduletypes.QueryCallbacks {
	c.callbacks[id] = fn.(Callback)
	return c
}

func (c Callbacks) RegisterCallbacks() ethquerymoduletypes.QueryCallbacks {
	a := c.
		AddCallback("getstorageat", Callback(GetStorageCallback)).
		AddCallback("getblock", Callback(GetBlockCallback))

	return a.(Callbacks)
}

// Callback Handler

func GetStorageCallback(k *Keeper, ctx sdk.Context, args []byte, query ethquerymoduletypes.EthQuery) error {
	var storageData eth_types.QueryRespoonseData
	err := json.Unmarshal(args, &storageData)
	if err != nil {
		return err
	}
	var slotData types.SlotData

	height := storageData.StorageProofData.Height.Uint64()

	proof, err := json.Marshal(storageData.StorageProofData)
	if err != nil {
		return err
	}

	Data := storageData.StorageData

	slotData = types.SlotData{
		Height:          height,
		Slot:            storageData.Slot,
		ContractAddress: storageData.StorageProofData.Address.Hex(),
		Proof:           proof,
		Verified:        false,
		Data:            Data,
	}
	k.SetSlotData(ctx, &slotData)

	return nil
}

func GetBlockCallback(k *Keeper, ctx sdk.Context, args []byte, query ethquerymoduletypes.EthQuery) error {
	var blockData eth_types.QueryBlockHeaderResponseData
	err := json.Unmarshal(args, &blockData)
	if err != nil {
		return err
	}

	var block types.BlockData

	height := blockData.Height
	data, err := json.Marshal(blockData.Root)
	if err != nil {
		return err
	}

	block = types.BlockData{
		Height: height,
		Data:   data,
	}

	k.SetBlockData(ctx, &block)

	return nil
}
