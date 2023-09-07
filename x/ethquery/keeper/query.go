package keeper

import (
	"cosmicether/x/ethquery/types"
)

var _ types.QueryServer = Keeper{}
