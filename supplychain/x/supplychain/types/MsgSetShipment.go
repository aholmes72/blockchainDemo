package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetShipment{}

type MsgSetShipment struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ShippingNumber string `json:"ShippingNumber" yaml:"ShippingNumber"`
  ShipTo string `json:"ShipTo" yaml:"ShipTo"`
  Item string `json:"Item" yaml:"Item"`
  ShipDate string `json:"ShipDate" yaml:"ShipDate"`
}

func NewMsgSetShipment(creator sdk.AccAddress, id string, ShippingNumber string, ShipTo string, Item string, ShipDate string) MsgSetShipment {
  return MsgSetShipment{
    ID: id,
		Creator: creator,
    ShippingNumber: ShippingNumber,
    ShipTo: ShipTo,
    Item: Item,
    ShipDate: ShipDate,
	}
}

func (msg MsgSetShipment) Route() string {
  return RouterKey
}

func (msg MsgSetShipment) Type() string {
  return "SetShipment"
}

func (msg MsgSetShipment) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetShipment) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetShipment) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}