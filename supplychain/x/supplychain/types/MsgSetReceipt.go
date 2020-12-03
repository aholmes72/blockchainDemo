package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetReceipt{}

type MsgSetReceipt struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ShippingNumber string `json:"ShippingNumber" yaml:"ShippingNumber"`
  ReceiptDate string `json:"ReceiptDate" yaml:"ReceiptDate"`
}

func NewMsgSetReceipt(creator sdk.AccAddress, id string, ShippingNumber string, ReceiptDate string) MsgSetReceipt {
  return MsgSetReceipt{
    ID: id,
		Creator: creator,
    ShippingNumber: ShippingNumber,
    ReceiptDate: ReceiptDate,
	}
}

func (msg MsgSetReceipt) Route() string {
  return RouterKey
}

func (msg MsgSetReceipt) Type() string {
  return "SetReceipt"
}

func (msg MsgSetReceipt) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetReceipt) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetReceipt) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}