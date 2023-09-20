package keeper

import (
	"github.com/ajansari95/cosmicether/x/ethstate/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetBlockData returns a blockData from its index
func (k *Keeper) GetBlockData(ctx sdk.Context, blockNumber uint64) (types.BlockData, bool) {
	blockData := types.BlockData{}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	bz := store.Get(types.GetBlockKey(blockNumber))
	if bz == nil {
		return blockData, false
	}
	k.cdc.MustUnmarshal(bz, &blockData)
	return blockData, true
}

// SetBlockData set a specific blockData in the store from its index
func (k *Keeper) SetBlockData(ctx sdk.Context, blockData *types.BlockData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	bz := k.cdc.MustMarshal(blockData)
	store.Set(types.GetBlockKey(blockData.Height), bz)
}

// DeleteBlockData deletes a blockData from the store
func (k *Keeper) DeleteBlockData(ctx sdk.Context, blockNumber uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	store.Delete(types.GetBlockKey(blockNumber))
}
