package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ajansari95/cosmicether/x/ethstate/types"
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
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
		sdk.NewEvent(
			types.EventTypeSubmitSlotData,
			sdk.NewAttribute(types.AttributeKeySlotData, msg.SlotData.String()),
			sdk.NewAttribute(types.AttributeKeySlotDataHeight, fmt.Sprintf("%v", msg.SlotData.Height)),
		),
	})

	//create ethquery request for block height
	err := m.Keeper.CreateBlockHeaderQuery(ctx, msg.SlotData.Height)
	if err != nil {
		return nil, err
	}

	return &types.MsgSubmitSlotDataResponse{}, nil
}

// GetSlotDataFromEth handles MsgGetSlotDataFromEth by creating a query on ethquery module for the data.
func (m msgServer) GetSlotDataFromEth(c context.Context, msg *types.MsgGetSlotDataFromEth) (*types.MsgGetSlotDataFromEthResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	height := msg.Height

	fmt.Println(height)

	request := queryGetStorageAt{
		ContractAddress: msg.ContractAddress,
		Slot:            msg.Slot,
		Height:          height,
	}

	reqbz, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	//create ethquery request for slot data and the height

	err = m.Keeper.CreateBlockHeaderQuery(ctx, height)
	if err != nil {
		return nil, err
	}

	m.ICQKeeper.MakeRequest(ctx, types.ModuleName, "eth_getStorageAt", reqbz, "getstorageat", uint64(ctx.BlockHeight()))

	m.Logger(ctx).Info("Created query on ethquery module for slot data and height);")

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
		sdk.NewEvent(
			types.EventTypeGetSlotDataFromEth,
			sdk.NewAttribute(types.AttributeKeySlotData, msg.Slot),
			sdk.NewAttribute(types.AttributeKeyContractAddress, msg.ContractAddress),
			sdk.NewAttribute(types.AttributeKeySlotDataHeight, fmt.Sprintf("%v", msg.Height)),
		),
	})

	return &types.MsgGetSlotDataFromEthResponse{}, nil
}
