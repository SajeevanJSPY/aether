package keeper

import (
	"context"

	"github.com/aether-proj/aether/x/fee/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper Keeper
}

func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return queryServer{Keeper: keeper}
}

func (q queryServer) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	return &types.QueryParamsResponse{
		TraderFeePercentage:   q.Keeper.GetTraderFeePercentage(ctx),
		ProtocolFeePercentage: q.Keeper.GetProtocolFeePercentage(ctx),
		LpFeePercentage:       q.Keeper.GetLpFeePercentage(ctx),
	}, nil
}

func (q queryServer) TotalFeeCollected(ctx context.Context, req *types.QueryTotalFeeCollectedRequest) (*types.QueryTotalFeeCollectedResponse, error) {
	return &types.QueryTotalFeeCollectedResponse{
		CollectedFees: q.Keeper.GetCollectedTotalFee(ctx),
	}, nil
}
