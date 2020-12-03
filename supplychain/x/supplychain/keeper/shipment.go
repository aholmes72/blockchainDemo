package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/foo/supplychain/x/supplychain/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetShipmentCount get the total number of shipment
func (k Keeper) GetShipmentCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ShipmentCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetShipmentCount set the total number of shipment
func (k Keeper) SetShipmentCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ShipmentCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateShipment creates a shipment
func (k Keeper) CreateShipment(ctx sdk.Context, msg types.MsgCreateShipment) {
	// Create the shipment
	count := k.GetShipmentCount(ctx)
    var shipment = types.Shipment{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        ShippingNumber: msg.ShippingNumber,
        ShipTo: msg.ShipTo,
        Item: msg.Item,
        ShipDate: msg.ShipDate,
    }

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ShipmentPrefix + shipment.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(shipment)
	store.Set(key, value)

	// Update shipment count
    k.SetShipmentCount(ctx, count+1)
}

// GetShipment returns the shipment information
func (k Keeper) GetShipment(ctx sdk.Context, key string) (types.Shipment, error) {
	store := ctx.KVStore(k.storeKey)
	var shipment types.Shipment
	byteKey := []byte(types.ShipmentPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &shipment)
	if err != nil {
		return shipment, err
	}
	return shipment, nil
}

// SetShipment sets a shipment
func (k Keeper) SetShipment(ctx sdk.Context, shipment types.Shipment) {
	shipmentKey := shipment.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(shipment)
	key := []byte(types.ShipmentPrefix + shipmentKey)
	store.Set(key, bz)
}

// DeleteShipment deletes a shipment
func (k Keeper) DeleteShipment(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ShipmentPrefix + key))
}

//
// Functions used by querier
//

func listShipment(ctx sdk.Context, k Keeper) ([]byte, error) {
	var shipmentList []types.Shipment
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ShipmentPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var shipment types.Shipment
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &shipment)
		shipmentList = append(shipmentList, shipment)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, shipmentList)
	return res, nil
}

func getShipment(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	shipment, err := k.GetShipment(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, shipment)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetShipmentOwner(ctx sdk.Context, key string) sdk.AccAddress {
	shipment, err := k.GetShipment(ctx, key)
	if err != nil {
		return nil
	}
	return shipment.Creator
}


// Check if the key exists in the store
func (k Keeper) ShipmentExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ShipmentPrefix + key))
}
