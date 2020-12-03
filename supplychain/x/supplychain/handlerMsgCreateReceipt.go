package supplychain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/foo/supplychain/x/supplychain/types"
	"github.com/foo/supplychain/x/supplychain/keeper"
)

func handleMsgCreateReceipt(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateReceipt) (*sdk.Result, error) {
	k.CreateReceipt(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
