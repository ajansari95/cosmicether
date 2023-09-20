package keeper

import (
	"fmt"

	"github.com/ajansari95/cosmicether/x/ethquery/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NewEthQuery(
	module string,
	queryType string,
	request []byte,
	callbackID string,
	height uint64,
) *types.EthQuery {
	return &types.EthQuery{
		Id:          GenerateQueryHash(queryType, request, module, fmt.Sprintf("%d", height)),
		QueryType:   queryType,
		Request:     request,
		CallbackId:  callbackID,
		BlockHeight: height,
	}
}

// GetEthQuery returns a ethQuery from its index
func (k Keeper) GetEthQuery(ctx sdk.Context, id string) (types.EthQuery, bool) {
	query := types.EthQuery{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	bz := store.Get(types.GetEthQueryKey(id))
	if bz == nil {
		return query, false
	}
	k.cdc.MustUnmarshal(bz, &query)
	return query, true
}

// SetEthQuery set a specific ethQuery in the store from its index
func (k Keeper) SetEthQuery(ctx sdk.Context, ethQuery types.EthQuery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	bz := k.cdc.MustMarshal(&ethQuery)
	store.Set(types.GetEthQueryKey(ethQuery.Id), bz)
}

// DeleteEthQuery deletes a ethQuery
func (k Keeper) DeleteEthQuery(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), nil)
	store.Delete(types.GetEthQueryKey(id))
}

// IterateEthQueries  iterate through all records
func (k Keeper) IterateEthQueries(ctx sdk.Context, fn func(index int64, ethQuery types.EthQuery) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEthQuery)

	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		ethQuery := types.EthQuery{}
		k.cdc.MustUnmarshal(iterator.Value(), &ethQuery)

		stop := fn(i, ethQuery)

		if stop {
			break
		}
		i++
	}
}

// AllEthQueries returns all ethQuery
func (k Keeper) AllEthQueries(ctx sdk.Context) []types.EthQuery {
	queries := []types.EthQuery{}

	k.IterateEthQueries(ctx, func(_ int64, query types.EthQuery) (stop bool) {
		queries = append(queries, query)
		return false
	})
	return queries
}
