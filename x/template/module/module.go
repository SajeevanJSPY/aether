package template

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/aether-proj/aether/x/template/keeper"
	"github.com/aether-proj/aether/x/template/types"
)

var (
	_ module.AppModuleBasic = (*AppModule)(nil)
	_ module.AppModule      = (*AppModule)(nil)
)

type AppModule struct{}

func NewAppModule() AppModule {
	return AppModule{}
}

func (AppModule) IsAppModule() {}

func (AppModule) Name() string {
	return types.ModuleName
}

func (AppModule) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

func (AppModule) RegisterInterfaces(registrar codectypes.InterfaceRegistry) {}

func (am AppModule) RegisterServices(registrar grpc.ServiceRegistrar) error {
	types.RegisterQueryServer(registrar, keeper.NewQueryServerImpl())
	return nil
}

func (AppModule) ConsensusVersion() uint64 { return 1 }

func (am AppModule) BeginBlock(_ context.Context) error {
	return nil
}

func (am AppModule) EndBlock(_ context.Context) error {
	return nil
}
