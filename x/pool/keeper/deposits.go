package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/aether-proj/aether/x/pool/types"
)

func (k Keeper) ValidateSender(ctx context.Context, sender string) (sdk.AccAddress, error) {
	addr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}

	if k.AccountKeeper.GetAccount(ctx, addr) == nil {
		return nil, fmt.Errorf("account %s does not exist", sender)
	}

	return addr, nil
}

func (k Keeper) CurrentDepositId(ctx context.Context) (depositId uint64) {
	id, err := k.DepositId.Peek(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (k Keeper) NextDepositId(ctx context.Context) (depositId uint64) {
	id, err := k.DepositId.Next(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (k Keeper) SetDeposit(ctx context.Context, depositor string, amount math.Int, poolId uint64) error {
	nextDepositId := k.NextDepositId(ctx)

	minDepositAmount := k.GetParams(ctx).MinDepositAmount

	if amount.LT(minDepositAmount) {
		return fmt.Errorf("the given amount `%d` is less than the minimum deposit amount `%d`", amount, minDepositAmount)
	}

	newDeposit := types.Deposit{
		Depositor: depositor,
		PoolId:    poolId,
		Shares:    amount,
	}

	err := k.Deposits.Set(ctx, nextDepositId, newDeposit)
	if err != nil {
		panic(err)
	}
	return nil
}

func (k Keeper) GetDeposit(ctx context.Context, depositId uint64) *types.Deposit {
	if depositId > k.CurrentDepositId(ctx) {
		panic("deposit id out of range")
	}
	val, err := k.Deposits.Get(ctx, depositId)
	if err != nil {
		panic(err)
	}
	return &val
}
