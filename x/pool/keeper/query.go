package keeper

import (
	"context"
	"fmt"

	"github.com/aether-proj/aether/x/pool/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper Keeper
}

func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return queryServer{Keeper: keeper}
}

func (q queryServer) GetPoolInfo(ctx context.Context, req *types.QueryGetPoolInfoRequest) (*types.QueryGetPoolInfoResponse, error) {
	if req == nil {
		panic("invalid request")
	}

	pool := q.Keeper.GetPool(ctx, req.PoolId)

	return &types.QueryGetPoolInfoResponse{
		Denom:       pool.Denom,
		TotalAmount: pool.TotalAmount,
		TotalShares: pool.TotalShares,
		TotalUsed:   pool.TotalUsed,
		CreatedAt:   pool.CreatedAt,
	}, nil
}

func (q queryServer) GetAllPoolInfos(ctx context.Context, req *types.QueryGetAllPoolInfosRequest) (*types.QueryGetAllPoolInfosResponse, error) {
	pools_count := q.Keeper.CurrentPoolId(ctx)

	pools := []*types.QueryGetPoolInfoResponse{}

	for i := uint64(1); i <= pools_count; i++ {
		pool := q.Keeper.GetPool(ctx, i)

		pool_info := &types.QueryGetPoolInfoResponse{
			Denom:       pool.Denom,
			TotalAmount: pool.TotalAmount,
			TotalShares: pool.TotalShares,
			TotalUsed:   pool.TotalUsed,
			CreatedAt:   pool.CreatedAt,
		}
		pools = append(pools, pool_info)

	}
	return &types.QueryGetAllPoolInfosResponse{
		Responses: pools,
	}, nil
}

func (q queryServer) GetUserInfo(ctx context.Context, req *types.QueryGetUserInfoRequest) (*types.QueryGetUserInfoResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("the requet is invalid")
	}

	q.Keeper.GetDeposit(ctx, q.Keeper.CurrentDepositId(ctx))

	return &types.QueryGetUserInfoResponse{
		Info: []*types.QueryGetUserInfoResponse_Info{},
	}, nil
}

func (q queryServer) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	params := q.Keeper.GetParams(ctx)

	return &types.QueryParamsResponse{
		MaxPools:          params.MaxPools,
		MinDepositAmount:  params.MinDepositAmount,
		MaxWithdrawAmount: params.MaxWithdrawAmount,
	}, nil
}
