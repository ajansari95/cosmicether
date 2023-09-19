package keeper

import (
	"github.com/ajansari95/cosmicether/x/ethstate/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSlotData returns a slotData from its index
func (k *Keeper) GetSlotData(ctx sdk.Context, contractAddress string, slot string) (types.SlotData, bool) {
	slotData := types.SlotData{}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	bz := store.Get(types.GetSlotKey(contractAddress, slot))
	if bz == nil {
		return slotData, false
	}
	k.cdc.MustUnmarshal(bz, &slotData)
	return slotData, true
}

// SetSlotData set a specific slotData in the store from its index
func (k *Keeper) SetSlotData(ctx sdk.Context, slotData *types.SlotData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	bz := k.cdc.MustMarshal(slotData)
	store.Set(types.GetSlotKey(slotData.ContractAddress, slotData.Slot), bz)
}

// DeleteSlotData deletes a slotData
func (k *Keeper) DeleteSlotData(ctx sdk.Context, contractAddress string, slot string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	store.Delete(types.GetSlotKey(contractAddress, slot))
}

// IteratePrefixedSlotData iterate through all records with a prefix.
func (k *Keeper) IteratePrefixedSlotData(ctx sdk.Context, prefixBytes []byte, fn func(index int64, slotData types.SlotData) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixSlotData)

	iterator := sdk.KVStorePrefixIterator(store, prefixBytes)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		slotData := types.SlotData{}
		k.cdc.MustUnmarshal(iterator.Value(), &slotData)

		stop := fn(i, slotData)

		if stop {
			break
		}
		i++
	}
}

// IterateSlotDatas iterate through all records.
func (k *Keeper) IterateSlotDatas(ctx sdk.Context, fn func(index int64, slotData types.SlotData) (stop bool)) {
	k.IteratePrefixedSlotData(ctx, nil, fn)
}

// AllContractSlotData returns all slotData associated to contractAddress
func (k *Keeper) AllContractSlotData(ctx sdk.Context, contractAddress string) []types.SlotData {
	var list []types.SlotData
	k.IteratePrefixedSlotData(ctx, []byte(contractAddress), func(_ int64, slotData types.SlotData) bool {
		list = append(list, slotData)
		return false
	})
	return list
}
