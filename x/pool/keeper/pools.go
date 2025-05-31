package keeper

import (
	"context"

	"github.com/aether-proj/aether/x/pool/types"
)

func (k Keeper) CurrentPoolId(ctx context.Context) (poolId uint64) {
	poolId, err := k.PoolId.Peek(ctx)
	if err != nil {
		panic(err)
	}
	return poolId
}

func (k Keeper) NextPoolId(ctx context.Context) (poolId uint64) {
	poolId, err := k.PoolId.Next(ctx)
	if err != nil {
		panic(err)
	}
	return poolId
}

func (k Keeper) GetPool(ctx context.Context, poolId uint64) *types.Pool {
	if poolId > k.CurrentPoolId(ctx) {
		panic("pool id out of range")
	}
	val, err := k.Pools.Get(ctx, poolId)
	if err != nil {
		panic(err)
	}
	return &val
}
