package keeper

import (
	"context"

	"github.com/aether-proj/aether/x/pool/types"
)

func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) {
	if err := k.PoolId.Set(ctx, 0); err != nil {
		panic(err)
	}
	if err := k.DepositId.Set(ctx, 0); err != nil {
		panic(err)
	}

	err := k.Params.Set(ctx, *genState.Params)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}
	genesis.Params = &params

	return genesis
}
