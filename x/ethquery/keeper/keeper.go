package keeper

import (
	"fmt"

	"github.com/ajansari95/cosmicether/x/ethquery/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		callbacks  map[string]types.QueryCallbacks
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		callbacks:  make(map[string]types.QueryCallbacks),
	}
}

func (k *Keeper) SetCallbackHandler(module string, handler types.QueryCallbacks) error {
	_, found := k.callbacks[module]
	if found {
		return fmt.Errorf("callback handler already set for %s", module)
	}
	k.callbacks[module] = handler.RegisterCallbacks()
	return nil
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) MakeRequest(
	ctx sdk.Context,
	module string,
	queryType string,
	request []byte,
	callbackID string,
	height uint64,
) {
	k.Logger(ctx).Info("MakeRequest",
		"queryType", queryType,
		"request", request,
		"callbackID", callbackID,
		"height", height)

	key := GenerateQueryHash(queryType, request, module, fmt.Sprintf("%d", height))
	existing, found := k.GetEthQuery(ctx, key)
	if !found {
		if module != "" && callbackID != "" {
			if _, exists := k.callbacks[module]; !exists {
				err := fmt.Errorf("no callback handler registered for module %s", module)
				k.Logger(ctx).Error(err.Error())
				panic(err)
			}
			if exists := k.callbacks[module].Has(callbackID); !exists {
				err := fmt.Errorf("no callback %s registered for module %s", callbackID, module)
				k.Logger(ctx).Error(err.Error())
				panic(err)
			}
		}
		newQuery := k.NewEthQuery(module, queryType, request, callbackID, height)
		k.SetEthQuery(ctx, *newQuery)
	} else {
		k.Logger(ctx).Info("re-request", "LastHeight", existing.BlockHeight)
		existing.BlockHeight = height
		k.SetEthQuery(ctx, existing)
	}
}

func GenerateQueryHash(queryType string, request []byte, module string, height string) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(module+queryType+height), request...)))
}
