package types

import "cosmossdk.io/collections"

const (
	ModuleName    = "perp"
	StoreKey      = ModuleName
	GovModuleName = "gov"
)

var (
	ParamsKey   = collections.NewPrefix(1)
	PositionKey = collections.NewPrefix(2)
)
