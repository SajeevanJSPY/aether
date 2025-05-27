package keeper

import (
	"context"

	"github.com/aether-proj/aether/x/fee/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return msgServer{Keeper: keeper}
}

func (m msgServer) UpdateParams(ctx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if req.Authority != m.Keeper.GetAuthority() {
		return nil, types.ErrInvalidCaller
	}

	params := types.Params{
		TraderFeePercentage:   req.TraderFeePercentage,
		ProtocolFeePercentage: req.ProtocolFeePercentage,
		LpFeePercentage:       req.LpFeePercentage,
	}

	m.Keeper.UpdateParams(ctx, params)

	return &types.MsgUpdateParamsResponse{}, nil
}
