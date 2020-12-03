package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteReceipt{}

type MsgDeleteReceipt struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteReceipt(id string, creator sdk.AccAddress) MsgDeleteReceipt {
  return MsgDeleteReceipt{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteReceipt) Route() string {
  return RouterKey
}

func (msg MsgDeleteReceipt) Type() string {
  return "DeleteReceipt"
}

func (msg MsgDeleteReceipt) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteReceipt) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteReceipt) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}