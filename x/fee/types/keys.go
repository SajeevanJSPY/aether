package types

import (
	"cosmossdk.io/collections"
)

const (
	ModuleName    = "fee"
	StoreKey      = ModuleName
	GovModuleName = "gov"
)

var (
	ParamsKey   = collections.NewPrefix(1)
	TotalFeeKey = collections.NewPrefix(2)
)
