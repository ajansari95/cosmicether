package keeper_test

import (
	"testing"

	testkeeper "github.com/ajansari95/cosmicether/testutil/keeper"
	"github.com/ajansari95/cosmicether/x/ethquery/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EthqueryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
