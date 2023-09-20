package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/ajansari95/cosmicether/testutil/keeper"
	"github.com/ajansari95/cosmicether/x/ethquery/keeper"
	"github.com/ajansari95/cosmicether/x/ethquery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EthqueryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
