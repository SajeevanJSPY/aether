package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidCaller = errorsmod.Register(ModuleName, 1100, "invalid caller")
)
