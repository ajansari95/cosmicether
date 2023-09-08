package keeper

import (
	"context"
	"cosmicether/x/ethquery/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) SubmitQueryResponse(ctx context.Context, response *types.MsgSubmitQueryResponse) (*types.MsgSubmitQueryResponseResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
