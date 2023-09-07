package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "cosmicether/testutil/keeper"
	"cosmicether/x/ethstate/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EthstateKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
