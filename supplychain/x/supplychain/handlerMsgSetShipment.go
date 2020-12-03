package supplychain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/foo/supplychain/x/supplychain/types"
	"github.com/foo/supplychain/x/supplychain/keeper"
)

func handleMsgSetShipment(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetShipment) (*sdk.Result, error) {
	var shipment = types.Shipment{
		Creator: msg.Creator,
		ID:      msg.ID,
    	ShippingNumber: msg.ShippingNumber,
    	ShipTo: msg.ShipTo,
    	Item: msg.Item,
    	ShipDate: msg.ShipDate,
	}
	if !msg.Creator.Equals(k.GetShipmentOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetShipment(ctx, shipment)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
