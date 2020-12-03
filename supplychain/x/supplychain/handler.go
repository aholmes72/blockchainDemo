package supplychain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/foo/supplychain/x/supplychain/keeper"
	"github.com/foo/supplychain/x/supplychain/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateShipment:
			return handleMsgCreateShipment(ctx, k, msg)
		case types.MsgSetShipment:
			return handleMsgSetShipment(ctx, k, msg)
		case types.MsgDeleteShipment:
			return handleMsgDeleteShipment(ctx, k, msg)
		case types.MsgCreateReceipt:
			return handleMsgCreateReceipt(ctx, k, msg)
		case types.MsgSetReceipt:
			return handleMsgSetReceipt(ctx, k, msg)
		case types.MsgDeleteReceipt:
			return handleMsgDeleteReceipt(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
