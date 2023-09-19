package keeper

import (
	"encoding/json"
	"fmt"

	ethquerykeeper "github.com/ajansari95/cosmicether/x/ethquery/keeper"
	"github.com/ajansari95/cosmicether/x/ethstate/types"
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
		ICQKeeper  ethquerykeeper.Keeper
	}
)
type queryGetStorageAt struct {
	ContractAddress string `json:"contractAddress"`
	Slot            string `json:"slot"`
	Height          uint64 `json:"height"`
}

type queryGetBlockHeader struct {
	Height uint64 `json:"height"`
}

func NewKeeper(
	cdc codec.BinaryCodec,
	icqKeeper ethquerykeeper.Keeper,
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
		ICQKeeper:  icqKeeper,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateBlockHeaderQuery(ctx sdk.Context, height uint64) error {
	//create ethquery request for block height
	request := queryGetBlockHeader{Height: height}

	reqbz, err := json.Marshal(request)
	if err != nil {
		return err
	}
	k.ICQKeeper.MakeRequest(ctx, types.ModuleName, "eth_getBlockByNumber", reqbz, "getblock", uint64(ctx.BlockHeight()))
	return nil
}
