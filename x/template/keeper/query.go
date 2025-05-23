package keeper

import (
	"context"

	"github.com/aether-proj/aether/x/template/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct{}

func NewQueryServerImpl() types.QueryServer {
	return queryServer{}
}

func (q queryServer) SayMyName(ctx context.Context, req *types.SayMyNameRequest) (*types.SayMyNameResponse, error) {
	return &types.SayMyNameResponse{
		SayMyName: "heisenberg",
	}, nil
}
