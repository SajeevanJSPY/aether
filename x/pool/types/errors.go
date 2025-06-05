package types

import "cosmossdk.io/errors"

var (
	ErrPoolAlreadyExist = errors.Register(ModuleName, 1300, "pool already exist")
	ErrPoolPaused = errors.Register(ModuleName, 1301, "pool is paused")
)
