package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidData     = sdkerrors.Register(ModuleName, 1, "invalid data")
	ErrInvalidHash     = sdkerrors.Register(ModuleName, 2, "invalid hash")
	ErrInvalidBTCHeight = sdkerrors.Register(ModuleName, 3, "invalid BTC height")
)
