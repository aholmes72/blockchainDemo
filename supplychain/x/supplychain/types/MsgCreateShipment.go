package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateShipment{}

type MsgCreateShipment struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ShippingNumber string `json:"ShippingNumber" yaml:"ShippingNumber"`
  ShipTo string `json:"ShipTo" yaml:"ShipTo"`
  Item string `json:"Item" yaml:"Item"`
  ShipDate string `json:"ShipDate" yaml:"ShipDate"`
}

func NewMsgCreateShipment(creator sdk.AccAddress, ShippingNumber string, ShipTo string, Item string, ShipDate string) MsgCreateShipment {
  return MsgCreateShipment{
		Creator: creator,
    ShippingNumber: ShippingNumber,
    ShipTo: ShipTo,
    Item: Item,
    ShipDate: ShipDate,
	}
}

func (msg MsgCreateShipment) Route() string {
  return RouterKey
}

func (msg MsgCreateShipment) Type() string {
  return "CreateShipment"
}

func (msg MsgCreateShipment) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateShipment) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateShipment) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}