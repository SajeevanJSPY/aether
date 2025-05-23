package template

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"

	"github.com/aether-proj/aether/x/template/keeper"
	"github.com/aether-proj/aether/x/template/types"
)

var _ depinject.OnePerModuleType = AppModule{}

func (AppModule) IsOnePerModuleType() {}

func init() {
	appconfig.Register(
		&types.Module{},
		appconfig.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In
}

type ModuleOutputs struct {
	depinject.Out

	TemplateKeeper keeper.Keeper
	Module         appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	k := keeper.NewKeeper()
	m := NewAppModule()

	return ModuleOutputs{TemplateKeeper: k, Module: m}
}
