package keeper

import (
	"context"

	"github.com/aether-proj/aether/x/perp/types"
)

func (k msgServer) OpenPosition(goCtx context.Context, req *types.MsgOpenPosition) (*types.MsgOpenPositionResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	return &types.MsgOpenPositionResponse{}, nil
}

func (k msgServer) ClosePosition(ctx context.Context, req *types.MsgClosePosition) (*types.MsgClosePositionResponse, error) {
	if req == nil {
		return nil, types.ErrInvalidRequest
	}

	return &types.MsgClosePositionResponse{}, nil
}

func (k msgServer) Liquidate(ctx context.Context, req *types.MsgLiquidate) (*types.MsgLiquidateResponse, error) {
	return &types.MsgLiquidateResponse{}, nil
}
