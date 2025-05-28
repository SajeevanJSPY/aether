package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/aether-proj/aether/x/fee/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return msgServer{Keeper: keeper}
}

func (m msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if req.Authority != m.Keeper.GetAuthority() {
		return nil, types.ErrInvalidCaller
	}

	params := types.Params{
		TraderFeePercentage:   req.TraderFeePercentage,
		ProtocolFeePercentage: req.ProtocolFeePercentage,
		LpFeePercentage:       req.LpFeePercentage,
	}

	m.Keeper.UpdateParams(ctx, params)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventUpdatedParams, sdk.NewAttribute(types.AttributeTraderFeePercentage, params.TraderFeePercentage.String())))

	return &types.MsgUpdateParamsResponse{}, nil
}
