package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/math"
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
	if req == nil {
		panic("request is invalid")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(req.Authority)
	if err != nil {
		return nil, err
	}

	// get the pool id for
	poolId := m.Keeper.NextPoolId(ctx)

	baseDenom := req.BaseDenom
	quoteDenom := req.QuoteDenom
	if baseDenom == "" {
		panic("base denom cannot be empty")
	}
	if quoteDenom == "" {
		panic("quote denom cannot be empty")
	}
	totalAmount := sdk.NewDecCoin(baseDenom, math.NewInt(0))

	pool := &types.Pool{
		Creator:        creator.String(),
		TotalAmount:    &totalAmount,
		QuoteDenom:     req.QuoteDenom,
		TotalShares:    math.NewInt(0),
		ReservedAmount: math.NewInt(0),
		CreatedAt:      time.Now(),
	}

	m.Keeper.SetPool(ctx, poolId, *pool)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventPoolCreated,
		sdk.NewAttribute(types.AttributePoolId, fmt.Sprintf("%d", poolId)),
		sdk.NewAttribute(types.AttributePoolCreator, creator.String()),
	))

	return &types.MsgCreatePoolResponse{
		PoolId: poolId,
	}, nil
}

func (m msgServer) RemovePool(goCtx context.Context, req *types.MsgRemovePool) (*types.MsgRemovePoolResponse, error) {

	return &types.MsgRemovePoolResponse{}, nil
}

func (m msgServer) Deposit(goCtx context.Context, req *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	if req == nil {
		panic("request is invalid")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	depositor, err := m.Keeper.ValidateSender(ctx, req.Depositor)
	if err != nil {
		return nil, err
	}

	poolInfo := m.Keeper.GetPool(ctx, req.PoolId)
	balance := m.Keeper.BankKeeper.GetBalance(ctx, depositor, poolInfo.TotalAmount.GetDenom())

	if poolInfo.IsPaused {
		return nil, types.ErrPoolPaused
	}

	depositAmount := req.DepositAmount

	if depositAmount.GT(balance.Amount) {
		return nil, fmt.Errorf("your bank balance is lesser than the deposit amount")
	}

	amount := sdk.NewCoin(poolInfo.TotalAmount.GetDenom(), depositAmount)

	// send tokens from user to the protocol
	m.Keeper.BankKeeper.SendCoinsFromAccountToModule(ctx, depositor, types.ModuleName, []sdk.Coin{amount})

	shares := m.Keeper.CalculateShares(poolInfo.TotalAmount.Amount.TruncateInt(), poolInfo.TotalShares, depositAmount)

	m.Keeper.SetDeposit(ctx, depositor.String(), shares, req.PoolId)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventDeposited,
		sdk.NewAttribute(types.AttributePoolId, fmt.Sprintf("%d", req.PoolId)),
		sdk.NewAttribute(types.AttributeShares, shares.String()),
	))

	return &types.MsgDepositResponse{
		Shares: shares,
	}, nil
}

func (m msgServer) Withdraw(goCtx context.Context, req *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	return &types.MsgWithdrawResponse{}, nil
}

func (m msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	return &types.MsgUpdateParamsResponse{}, nil
}
