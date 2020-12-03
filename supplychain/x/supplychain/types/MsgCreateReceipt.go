package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateReceipt{}

type MsgCreateReceipt struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ShippingNumber string `json:"ShippingNumber" yaml:"ShippingNumber"`
  ReceiptDate string `json:"ReceiptDate" yaml:"ReceiptDate"`
}

func NewMsgCreateReceipt(creator sdk.AccAddress, ShippingNumber string, ReceiptDate string) MsgCreateReceipt {
  return MsgCreateReceipt{
		Creator: creator,
    ShippingNumber: ShippingNumber,
    ReceiptDate: ReceiptDate,
	}
}

func (msg MsgCreateReceipt) Route() string {
  return RouterKey
}

func (msg MsgCreateReceipt) Type() string {
  return "CreateReceipt"
}

func (msg MsgCreateReceipt) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateReceipt) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateReceipt) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}