package types

import (
	"fmt"

	"cosmossdk.io/math"
)

var (
	DefaultFeePercentage         = math.LegacyMustNewDecFromStr("0.01")
	DefaultProtocolFeePercentage = math.LegacyMustNewDecFromStr("0.35")
	DefaultLpFeePercentage       = math.LegacyMustNewDecFromStr("0.65")
)

func DefaultParams() *Params {
	return &Params{
		TraderFeePercentage:   DefaultFeePercentage,
		ProtocolFeePercentage: DefaultProtocolFeePercentage,
		LpFeePercentage:       DefaultLpFeePercentage,
	}
}

func (p Params) Validate() error {
	if p.TraderFeePercentage.IsNegative() || p.TraderFeePercentage.GT(math.LegacyOneDec()) || p.ProtocolFeePercentage.IsZero() {
		return fmt.Errorf("fee percentage must be between 0 and 1, got: %s", p.TraderFeePercentage)
	}

	if p.ProtocolFeePercentage.IsNegative() || p.ProtocolFeePercentage.GT(math.LegacyOneDec()) || p.ProtocolFeePercentage.IsZero() {
		return fmt.Errorf("fee percentage must be between 0 and 1, got: %s", p.ProtocolFeePercentage)
	}

	if p.LpFeePercentage.IsNegative() || p.LpFeePercentage.GT(math.LegacyOneDec()) || p.LpFeePercentage.IsZero() {
		return fmt.Errorf("fee percentage must be between 0 and 1, got: %s", p.LpFeePercentage)
	}

	return nil
}
