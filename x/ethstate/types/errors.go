package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ethstate module sentinel errors
var (
	ErrSample          = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrSlotDataNil     = sdkerrors.Register(ModuleName, 1102, "slot data is nil")
	ErrSlotDataInvalid = sdkerrors.Register(ModuleName, 1103, "slot data is invalid")
)
