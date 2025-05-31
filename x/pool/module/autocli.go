package pool

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"github.com/aether-proj/aether/x/pool/types"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetPoolInfo",
					Use:       "get-pool-info",
					Short:     "get the pool information",
				},
				{
					RpcMethod: "GetAllPoolInfos",
					Use:       "get-pool-infos",
					Short:     "get all the pool information",
				},
				{
					RpcMethod: "GetUserInfo",
					Use:       "get-user-info",
					Short:     "get user information",
				},
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "get the protocol params",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: types.Msg_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreatePool",
					Use:       "create-pool",
					Short:     "create a pool",
				},
				{
					RpcMethod: "RemovePool",
					Use:       "remove-pool",
					Short:     "remove pool",
				},
				{
					RpcMethod: "Deposit",
					Use:       "deposit",
					Short:     "deposit tokens to a pool",
				},
				{
					RpcMethod: "Withdraw",
					Use:       "withdraw",
					Short:     "withdraw tokens from a pool",
				},
				{
					RpcMethod: "UpdateParams",
					Use:       "update-params",
					Short:     "update the protocol params",
				},
			},
		},
	}
}
