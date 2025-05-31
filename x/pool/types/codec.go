package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "aether/pool/create-pool", nil)
	cdc.RegisterConcrete(&MsgRemovePool{}, "aether/pool/remove-pool", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "aether/pool/deposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "aether/pool/withdraw", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "aether/pool/update-params", nil)
}

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
		&MsgRemovePool{},
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
