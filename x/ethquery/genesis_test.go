package ethquery_test

import (
	"testing"

	keepertest "github.com/ajansari95/cosmicether/testutil/keeper"
	"github.com/ajansari95/cosmicether/testutil/nullify"
	"github.com/ajansari95/cosmicether/x/ethquery"
	"github.com/ajansari95/cosmicether/x/ethquery/types"
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
