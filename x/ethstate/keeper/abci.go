package keeper

import (
	"encoding/json"

	"github.com/ajansari95/cosmicether/eth_types"
	"github.com/ajansari95/cosmicether/x/ethstate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

const blockInterval = 10

func (k *Keeper) BeginBlocker(ctx sdk.Context) {

	if ctx.BlockHeight()%blockInterval == 0 {
		k.Logger(ctx).Info("Verifying the slotData", "height", ctx.BlockHeight())
		k.IterateSlotDatas(ctx, func(index int64, slotData types.SlotData) (stop bool) {
			if slotData.Verified == false {
				k.Logger(ctx).Info("Verifying the slotData", "height", ctx.BlockHeight())
				blockData, found := k.GetBlockData(ctx, slotData.Height)
				if found {

					var rootHash common.Hash

					err := json.Unmarshal(blockData.Data, &rootHash)
					if err != nil {
						k.Logger(ctx).Error("Error while unmarshalling the blockData", "error", err)
						return
					}

					var proof eth_types.StorageProof

					err = json.Unmarshal(slotData.Proof, &proof)

					verified, err := eth_types.VerifyMPTForSlotData(&proof, rootHash)
					if err != nil {
						k.Logger(ctx).Error("Error while verifying the slotData", "error", err)
						return
					}
					if verified {
						slotData.Verified = true
						k.SetSlotData(ctx, &slotData)
						k.Logger(ctx).Info("SlotData verified", "height", ctx.BlockHeight())
					}

				}

			}
			return false
		})
	}
}
