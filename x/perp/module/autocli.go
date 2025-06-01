package template

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"github.com/aether-proj/aether/x/perp/types"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetPosition",
					Use:       "get-position",
					Short:     "get the information about a position",
				},
				{
					RpcMethod: "GetLiquidity",
					Use:       "get-liquidity",
					Short:     "get liquidity information",
				},
				{
					RpcMethod: "GetLiquidationCandidates",
					Use:       "get-liquidity-candidates",
					Short:     "get the liquidity candidates",
				},
				{
					RpcMethod: "GetFundingRate",
					Use:       "get-funding-rate",
					Short:     "get funding rate",
				},
				{
					RpcMethod: "GetPositionFundingInfo",
					Use:       "get-position-funding-info",
					Short:     "get funding rate for a position",
				},
				{
					RpcMethod: "GetAccruedFunding",
					Use:       "get-accrued-funding",
					Short:     "get accrued funding",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: types.Msg_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "OpenPosition",
					Use:       "open-position",
					Short:     "open up a position",
				},
				{
					RpcMethod: "ClosePosition",
					Use:       "close-position",
					Short:     "close a position",
				},
				{
					RpcMethod: "Liquidate",
					Use:       "liquidate",
					Short:     "liquidate a position",
				},
				{
					RpcMethod: "UpdateFundingRate",
					Use:       "update-funding-rate",
					Short:     "update the funding rate",
				},
				{
					RpcMethod: "ApplyFundingPayment",
					Use:       "apply-funding-payment",
					Short:     "apply the funding payment",
				},
				{
					RpcMethod: "SettleFundingPayment",
					Use:       "settle-funding-payment",
					Short:     "settle the funding payment",
				},
			},
		},
	}
}
