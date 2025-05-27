package keeper

import (
	"context"

	"github.com/aether-proj/aether/x/fee/types"
)

func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	return k.Params.Set(ctx, *genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	genesis.Params = &params

	return genesis, nil
}
