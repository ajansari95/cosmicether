package keeper

import (
	"context"
	"cosmicether/x/ethstate/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) GetSlotDataFromEth(ctx context.Context, eth *types.MsgGetSlotDataFromEth) (*types.MsgGetSlotDataFromEthResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) SubmitSlotData(ctx context.Context, data *types.MsgSubmitSlotData) (*types.MsgSubmitSlotDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
