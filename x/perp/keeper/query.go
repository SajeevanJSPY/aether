package keeper

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/aether-proj/aether/x/perp/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct{}

func NewQueryServerImpl() types.QueryServer {
	return queryServer{}
}

func (q queryServer) GetPosition(ctx context.Context, req *types.QueryGetPositionRequest) (*types.QueryGetPositionResponse, error) {
	decCoin := sdk.NewDecCoin("aeth", math.NewInt(1000000))

	return &types.QueryGetPositionResponse{
		Trader: "heisenberg",
		IsLong: true,
		Size_:  "size debug",
		Margin: &decCoin,
	}, nil
}

func (q queryServer) GetLiquidity(ctx context.Context, req *types.QueryLiquidityRequest) (*types.QueryLiquidityResponse, error) {
	return &types.QueryLiquidityResponse{
		Shares: math.LegacyNewDec(100000),
	}, nil
}

func (q queryServer) GetLiquidationCandidates(ctx context.Context, req *types.QueryGetLiquidationCandidatesRequest) (*types.QueryGetLiquidationCandidatesResponse, error) {
	return &types.QueryGetLiquidationCandidatesResponse{
		Candidates: []string{
			"heisenberg",
			"eren yeager",
		},
	}, nil
}

func (q queryServer) GetFundingRate(ctx context.Context, req *types.QueryGetFundingRateRequest) (*types.QueryGetFundingRateResponse, error) {
	return &types.QueryGetFundingRateResponse{}, nil
}

func (q queryServer) GetPositionFundingInfo(ctx context.Context, req *types.QueryGetPositionFundingInfoRequest) (*types.QueryGetPositionFundingInfoResponse, error) {
	return &types.QueryGetPositionFundingInfoResponse{}, nil
}

func (q queryServer) GetAccruedFunding(ctx context.Context, req *types.QueryGetAccruedFundingRequest) (*types.QueryGetAccruedFundingResponse, error) {
	return &types.QueryGetAccruedFundingResponse{}, nil
}
