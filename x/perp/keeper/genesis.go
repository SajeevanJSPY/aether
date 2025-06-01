package keeper

import (
	"context"

	"github.com/aether-proj/aether/x/perp/types"
)

func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) {
	err := k.Params.Set(ctx, *genState.Params)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}
	return &types.GenesisState{
		Params: &params,
	}
}
