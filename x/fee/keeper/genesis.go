package keeper

import (
	"context"

	"cosmossdk.io/math"
	"github.com/aether-proj/aether/x/fee/types"
)

func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) {
	// set the default value for the total fee collected
	if err := k.TotalFees.Set(ctx, types.TotalFee{Amount: math.LegacyZeroDec()}); err != nil {
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
