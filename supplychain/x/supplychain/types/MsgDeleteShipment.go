package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteShipment{}

type MsgDeleteShipment struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteShipment(id string, creator sdk.AccAddress) MsgDeleteShipment {
  return MsgDeleteShipment{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteShipment) Route() string {
  return RouterKey
}

func (msg MsgDeleteShipment) Type() string {
  return "DeleteShipment"
}

func (msg MsgDeleteShipment) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteShipment) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteShipment) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}