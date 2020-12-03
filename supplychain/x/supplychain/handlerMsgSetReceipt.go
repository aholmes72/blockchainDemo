package supplychain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/foo/supplychain/x/supplychain/types"
	"github.com/foo/supplychain/x/supplychain/keeper"
)

func handleMsgSetReceipt(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetReceipt) (*sdk.Result, error) {
	var receipt = types.Receipt{
		Creator: msg.Creator,
		ID:      msg.ID,
    	ShippingNumber: msg.ShippingNumber,
    	ReceiptDate: msg.ReceiptDate,
	}
	if !msg.Creator.Equals(k.GetReceiptOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetReceipt(ctx, receipt)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
