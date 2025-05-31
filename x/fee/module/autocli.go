package fee

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"github.com/aether-proj/aether/x/fee/types"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "get the protocol params",
				},
				{
					RpcMethod: "TotalFeeCollected",
					Use:       "total-fee-collected",
					Short:     "get the collected fee",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: types.Msg_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Use:       "update-params",
					Short:     "update the protocol params",
				},
			},
		},
	}
}
