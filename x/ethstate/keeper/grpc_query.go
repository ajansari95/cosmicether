package keeper

import (
	"context"
	"cosmicether/x/ethstate/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// SlotData implements the Query/SlotData gRPC method for getting a single slotData
func (k Keeper) SlotData(c context.Context, request *types.QuerySlotDataRequest) (*types.QuerySlotDataRequestResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	slotData, found := k.GetSlotData(ctx, request.ContractAddress, request.Slot)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QuerySlotDataRequestResponse{SlotData: &slotData}, nil
}

// ContractData implements the Query/ContractData gRPC method for getting a single contractData
func (k Keeper) ContractData(c context.Context, request *types.QueryContractDataRequest) (*types.QueryContractDataRequestResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	slotDatalist := k.AllContractSlotData(ctx, request.ContractAddress)

	return &types.QueryContractDataRequestResponse{Slots: slotDatalist}, nil
}

func (k Keeper) EthBlock(c context.Context, request *types.QueryEthBlockRequest) (*types.QueryEthBlockRequestResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	ethBlock, found := k.GetBlockData(ctx, request.BlockHeight)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryEthBlockRequestResponse{BlockData: &ethBlock}, nil
}
