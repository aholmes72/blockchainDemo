package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreateShipment{}, "supplychain/CreateShipment", nil)
		cdc.RegisterConcrete(MsgSetShipment{}, "supplychain/SetShipment", nil)
		cdc.RegisterConcrete(MsgDeleteShipment{}, "supplychain/DeleteShipment", nil)
		cdc.RegisterConcrete(MsgCreateReceipt{}, "supplychain/CreateReceipt", nil)
		cdc.RegisterConcrete(MsgSetReceipt{}, "supplychain/SetReceipt", nil)
		cdc.RegisterConcrete(MsgDeleteReceipt{}, "supplychain/DeleteReceipt", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
