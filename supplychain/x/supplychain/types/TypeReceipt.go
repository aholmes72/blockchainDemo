package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Receipt struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    ShippingNumber string `json:"ShippingNumber" yaml:"ShippingNumber"`
    ReceiptDate string `json:"ReceiptDate" yaml:"ReceiptDate"`
}