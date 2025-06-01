package types

import (
	"cosmossdk.io/errors"
)

var (
	ErrInvalidRequest = errors.Register(ModuleName, 1200, "the request must be valid")
)
