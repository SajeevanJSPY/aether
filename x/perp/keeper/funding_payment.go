package keeper

import (
	"context"

	"github.com/aether-proj/aether/x/perp/types"
)

func (k msgServer) UpdateFundingRate(ctx context.Context, req *types.MsgUpdateFundingRate) (*types.MsgUpdateFundingRateResponse, error) {
	return &types.MsgUpdateFundingRateResponse{}, nil
}

func (k msgServer) ApplyFundingPayment(ctx context.Context, req *types.MsgApplyFundingPayment) (*types.MsgApplyFundingPaymentRespone, error) {
	return &types.MsgApplyFundingPaymentRespone{}, nil
}

func (k msgServer) SettleFundingPayment(ctx context.Context, req *types.MsgSettleFundingPayment) (*types.MsgSettleFundingPaymentResponse, error) {
	return &types.MsgSettleFundingPaymentResponse{}, nil
}
