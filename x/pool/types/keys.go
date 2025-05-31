package types

import (
	"cosmossdk.io/collections"
)

const (
	ModuleName    = "pool"
	StoreKey      = ModuleName
	GovModuleName = "gov"
)

var (
	ParamsKey            = collections.NewPrefix(1)
	PoolIdKey            = collections.NewPrefix(2)
	PoolsKey             = collections.NewPrefix(3)
	DepositIdKey         = collections.NewPrefix(4)
	DepositsKey          = collections.NewPrefix(5)
	DepositorIndexPrefix = collections.NewPrefix(6)
)
