package supplychain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/foo/supplychain/x/supplychain/types"
	"github.com/foo/supplychain/x/supplychain/keeper"
)

func handleMsgCreateShipment(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateShipment) (*sdk.Result, error) {
	k.CreateShipment(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
