package keeper

import (
	"github.com/aether-proj/aether/x/perp/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

func NewMsgServer(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}
