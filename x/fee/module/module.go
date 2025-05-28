package template

import (
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/aether-proj/aether/x/fee/keeper"
	"github.com/aether-proj/aether/x/fee/types"
)

var (
	_ module.AppModuleBasic   = AppModule{}
	_ module.HasGenesisBasics = AppModule{}

	_ appmodule.AppModule        = AppModule{}
	_ module.HasConsensusVersion = AppModule{}
	_ module.HasGenesis          = AppModule{}
)

type AppModule struct {
	keeper keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		keeper: k,
	}
}

func (AppModule) IsAppModule() {}

func (AppModule) Name() string {
	return types.ModuleName
}

func (AppModule) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {}

func (AppModule) RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registrar)
}

func (am AppModule) RegisterServices(registrar grpc.ServiceRegistrar) error {
	types.RegisterQueryServer(registrar, keeper.NewQueryServerImpl(am.keeper))
	return nil
}

func (am AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form.
func (am AppModule) ValidateGenesis(cdc codec.JSONCodec, encodingConfig client.TxEncodingConfig, gs json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(gs, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return genState.Validate()
}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) {
	var genState types.GenesisState
	cdc.MustUnmarshalJSON(gs, &genState)

	am.keeper.InitGenesis(ctx, genState)
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := am.keeper.ExportGenesis(ctx)

	return cdc.MustMarshalJSON(genState)
}

func (AppModule) ConsensusVersion() uint64 { return 1 }
