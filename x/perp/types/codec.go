package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgOpenPosition{}, "aether/perp/open-position", nil)
	cdc.RegisterConcrete(&MsgClosePosition{}, "aether/perp/close-position", nil)
	cdc.RegisterConcrete(&MsgLiquidate{}, "aether/perp/liquidate", nil)
	cdc.RegisterConcrete(&MsgUpdateFundingRate{}, "aether/perp/update-funding-rate", nil)
	cdc.RegisterConcrete(&MsgApplyFundingPayment{}, "aether/perp/apply-funding-payment", nil)
	cdc.RegisterConcrete(&MsgSettleFundingPayment{}, "aether/perp/settle-funding-payment", nil)
}

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOpenPosition{},
		&MsgClosePosition{},
		&MsgLiquidate{},
		&MsgUpdateFundingRate{},
		&MsgApplyFundingPayment{},
		&MsgSettleFundingPayment{},
	)

	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
