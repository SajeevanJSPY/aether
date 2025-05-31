package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/aether-proj/aether/x/pool/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return msgServer{Keeper: keeper}
}

func (m msgServer) CreatePool(goCtx context.Context, req *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	poolId := m.Keeper.NextPoolId(ctx)

	return &types.MsgCreatePoolResponse{
		PoolId: poolId,
	}, nil
}

func (m msgServer) RemovePool(goCtx context.Context, req *types.MsgRemovePool) (*types.MsgRemovePoolResponse, error) {

	return &types.MsgRemovePoolResponse{}, nil
}

func (m msgServer) Deposit(goCtx context.Context, req *types.MsgDeposit) (*types.MsgDepositResponse, error) {

	return &types.MsgDepositResponse{}, nil
}

func (m msgServer) Withdraw(goCtx context.Context, req *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	return &types.MsgWithdrawResponse{}, nil
}

func (m msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	return &types.MsgUpdateParamsResponse{}, nil
}
