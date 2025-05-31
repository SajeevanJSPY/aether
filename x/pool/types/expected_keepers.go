package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
}
