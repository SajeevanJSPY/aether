package types

import (
	"errors"

	"cosmossdk.io/math"
)

const (
	MaxMaxPools     = 50
	DefaultMaxPools = 20
)

var (
	DefaultMinDepositAmount  = math.NewInt(1000)
	MaxMinDepositAmount      = math.NewInt(3000)
	DefaultMaxWithdrawAmount = math.NewInt(100000)
	MaxMaxWithdrawAmount     = math.NewInt(300000)
)

func DefaultParams() *Params {
	return &Params{
		MaxPools:          DefaultMaxPools,
		MinDepositAmount:  DefaultMinDepositAmount,
		MaxWithdrawAmount: DefaultMaxWithdrawAmount,
	}
}

func (p Params) Validate() error {
	if p.MaxPools > MaxMaxPools {
		return errors.New("max pools exceeds the max max pools")
	}

	if p.MinDepositAmount.GT(MaxMinDepositAmount) {
		return errors.New("min deposit amount exceeds the max deposit amount")
	}

	if p.MaxWithdrawAmount.GT(MaxMaxWithdrawAmount) {
		return errors.New("max withdraw amount exceeds the max withdraw amount")
	}

	return nil
}
