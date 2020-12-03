package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Shipment struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    ShippingNumber string `json:"ShippingNumber" yaml:"ShippingNumber"`
    ShipTo string `json:"ShipTo" yaml:"ShipTo"`
    Item string `json:"Item" yaml:"Item"`
    ShipDate string `json:"ShipDate" yaml:"ShipDate"`
}