package ethquery_test

import (
	keepertest "cosmicether/testutil/keeper"
	"cosmicether/testutil/nullify"
	"cosmicether/x/ethquery"
	"cosmicether/x/ethquery/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EthqueryKeeper(t)
	ethquery.InitGenesis(ctx, *k, genesisState)
	got := ethquery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
