package keeper

import (
	"context"
	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/aether-proj/aether/x/fee/types"
)

type FeeI interface {
	// calculating the fee, and update the totally collected fee
	CalculateFee(ctx context.Context, amount math.LegacyDec) (fee math.LegacyDec, afterFee math.LegacyDec)
	// get the totally collected fee
	GetCollectedTotalFee(ctx context.Context) math.LegacyDec
}

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec

	authority []byte

	Schema    collections.Schema
	Params    collections.Item[types.Params]
	TotalFees collections.Item[types.TotalFee]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,
) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		authority:    authority,
		addressCodec: addressCodec,

		Params:    collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		TotalFees: collections.NewItem(sb, types.TotalFeeKey, "totalfees", codec.CollValue[types.TotalFee](cdc)),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

func (k Keeper) CalculateFee(ctx context.Context, amount math.LegacyDec) (fee math.LegacyDec, afterFee math.LegacyDec) {
	fee = amount.Mul(k.GetTraderFeePercentage(ctx))
	amountAfterFeeReduction := amount.Sub(fee)

	total_collected_fee := k.GetCollectedTotalFee(ctx)
	updated_collected_fee := total_collected_fee.Add(fee)

	err := k.TotalFees.Set(ctx, types.TotalFee{Amount: updated_collected_fee})
	if err != nil {
		panic("failed to set the total collected fee")
	}

	return fee, amountAfterFeeReduction
}

func (k Keeper) GetAuthority() string {
	return string(k.authority)
}

func (k Keeper) GetParams(ctx context.Context) types.Params {
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic("failed to get the params")
	}
	return params
}

func (k Keeper) GetTraderFeePercentage(ctx context.Context) math.LegacyDec {
	return k.GetParams(ctx).TraderFeePercentage
}

func (k Keeper) GetProtocolFeePercentage(ctx context.Context) math.LegacyDec {
	return k.GetParams(ctx).ProtocolFeePercentage
}

func (k Keeper) GetLpFeePercentage(ctx context.Context) math.LegacyDec {
	return k.GetParams(ctx).LpFeePercentage
}

func (k Keeper) GetCollectedTotalFee(ctx context.Context) math.LegacyDec {
	fees, err := k.TotalFees.Get(ctx)
	if err != nil {
		panic(err)
	}
	return fees.Amount
}

func (k Keeper) UpdateParams(ctx context.Context, params types.Params) {
	err := k.Params.Set(ctx, params)

	if err != nil {
		panic("failed to set the protocol params")
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
