package keeper

import (
	"context"
	"cosmicether/x/ethquery/types"
	"errors"
	"fmt"
	"sort"

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

func (m msgServer) SubmitQueryResponse(c context.Context, response *types.MsgSubmitQueryResponse) (*types.MsgSubmitQueryResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	q, found := m.GetEthQuery(ctx, response.QueryId)
	if !found {
		m.Logger(ctx).Debug("query not found", "QueryID", response.QueryId)

		return &types.MsgSubmitQueryResponseResponse{}, nil
	}

	delete_flag := true

	keys := make([]string, 0)
	for k := range m.callbacks {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	callbackExecuted := false

	for _, key := range keys {
		module := m.callbacks[key]
		if module.Has(q.CallbackId) {
			err := module.Call(ctx, q.CallbackId, response.Result, q)
			callbackExecuted = true
			if err != nil {
				if !errors.Is(err, types.ErrSucceededNoDelete) {
					m.Logger(ctx).Error("error in callback", "error", err, "msg", response.QueryId, "result", response.Result, "type", q.QueryType, "params", q.Request)
					return nil, err
				}
				delete_flag = false
			}

			break
		}

	}
	if !callbackExecuted && q.CallbackId != "" {
		m.Logger(ctx).Error("callback expected but not found", "callbackId", q.CallbackId, "msg", response.QueryId, "type", q.QueryType)
		return nil, fmt.Errorf("expected callback %s, but did not find it", q.CallbackId)
	}

	if delete_flag {
		m.DeleteEthQuery(ctx, q.Id)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})
	return &types.MsgSubmitQueryResponseResponse{}, nil
}
