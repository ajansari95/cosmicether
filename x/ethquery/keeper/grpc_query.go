package keeper

import (
	"context"
	"cosmicether/x/ethquery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// Queries implements the QueryServer interface
func (k Keeper) Queries(c context.Context, request *types.QueryRequestsRequest) (*types.QueryRequestsResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	queries := k.AllEthQueries(ctx)

	return &types.QueryRequestsResponse{Quereis: queries}, nil
}
