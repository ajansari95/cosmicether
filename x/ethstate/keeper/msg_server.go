package keeper

import (
	"context"
	"cosmicether/x/ethstate/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// SubmitSlotData handles MsgSubmitSlotData by storing the data in the store.
func (m msgServer) SubmitSlotData(c context.Context, msg *types.MsgSubmitSlotData) (*types.MsgSubmitSlotDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	msg.SlotData.Verified = false

	m.Keeper.SetSlotData(ctx, msg.SlotData)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName,
			),
		),
		sdk.NewEvent(
			types.EventTypeSubmitSlotData,
			sdk.NewAttribute(types.AttributeKeySlotData, msg.SlotData.String()),
			sdk.NewAttribute(types.AttributeKeySlotDataHeight, fmt.Sprintf("%v", msg.SlotData.Height)),
		),
	})

	//create ethquery request for block height

	return &types.MsgSubmitSlotDataResponse{}, nil
}

// GetSlotDataFromEth handles MsgGetSlotDataFromEth by creating a query on ethquery module for the data.
func (m msgServer) GetSlotDataFromEth(c context.Context, msg *types.MsgGetSlotDataFromEth) (*types.MsgGetSlotDataFromEthResponse, error) {
	//ctx := sdk.UnwrapSDKContext(c)

	height := msg.Height
	//create ethquery request for block height
	fmt.Println(height)

	//create ethquery request for slot data and the heigh
	fmt.Println(msg.ContractAddress, msg.Slot, msg.Height)

	return &types.MsgGetSlotDataFromEthResponse{}, nil
}
