package supplychain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/foo/supplychain/x/supplychain/types"
	"github.com/foo/supplychain/x/supplychain/keeper"
)

// Handle a message to delete name
func handleMsgDeleteShipment(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteShipment) (*sdk.Result, error) {
	if !k.ShipmentExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetShipmentOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteShipment(ctx, msg.ID)
	return &sdk.Result{}, nil
}
