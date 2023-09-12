package types

// DONTCOVER

import (
	"errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ethquery module sentinel errors
var (
	ErrSample            = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrSucceededNoDelete = errors.New("query succeeded; do not not execute default behavior")
)
