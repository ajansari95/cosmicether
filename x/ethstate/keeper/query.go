package keeper

import (
	"cosmicether/x/ethstate/types"
)

var _ types.QueryServer = Keeper{}
