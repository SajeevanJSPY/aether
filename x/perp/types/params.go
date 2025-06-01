package types

import "cosmossdk.io/math"

var (
	DefaultMaintenanceMarginRatio = math.LegacyNewDec(10000)
	DefaultLiquidationPenalty     = math.LegacyNewDec(9000)
)

func DefaultParams() *Params {
	return &Params{
		MaintenanceMarginRatio: DefaultMaintenanceMarginRatio,
		LiquidationPenalty:     DefaultLiquidationPenalty,
	}
}

func (Params) Validate() error {
	return nil
}
