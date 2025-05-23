package template

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"github.com/aether-proj/aether/x/template/types"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "SayMyName",
					Use:       "say-my-name",
					Short:     "saying my name",
				},
			},
		},
	}
}
