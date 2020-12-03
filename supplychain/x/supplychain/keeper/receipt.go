package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/foo/supplychain/x/supplychain/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetReceiptCount get the total number of receipt
func (k Keeper) GetReceiptCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ReceiptCountPrefix)
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

// SetReceiptCount set the total number of receipt
func (k Keeper) SetReceiptCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.ReceiptCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateReceipt creates a receipt
func (k Keeper) CreateReceipt(ctx sdk.Context, msg types.MsgCreateReceipt) {
	// Create the receipt
	count := k.GetReceiptCount(ctx)
    var receipt = types.Receipt{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        ShippingNumber: msg.ShippingNumber,
        ReceiptDate: msg.ReceiptDate,
    }

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ReceiptPrefix + receipt.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(receipt)
	store.Set(key, value)

	// Update receipt count
    k.SetReceiptCount(ctx, count+1)
}

// GetReceipt returns the receipt information
func (k Keeper) GetReceipt(ctx sdk.Context, key string) (types.Receipt, error) {
	store := ctx.KVStore(k.storeKey)
	var receipt types.Receipt
	byteKey := []byte(types.ReceiptPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &receipt)
	if err != nil {
		return receipt, err
	}
	return receipt, nil
}

// SetReceipt sets a receipt
func (k Keeper) SetReceipt(ctx sdk.Context, receipt types.Receipt) {
	receiptKey := receipt.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(receipt)
	key := []byte(types.ReceiptPrefix + receiptKey)
	store.Set(key, bz)
}

// DeleteReceipt deletes a receipt
func (k Keeper) DeleteReceipt(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ReceiptPrefix + key))
}

//
// Functions used by querier
//

func listReceipt(ctx sdk.Context, k Keeper) ([]byte, error) {
	var receiptList []types.Receipt
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ReceiptPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var receipt types.Receipt
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &receipt)
		receiptList = append(receiptList, receipt)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, receiptList)
	return res, nil
}

func getReceipt(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	receipt, err := k.GetReceipt(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, receipt)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetReceiptOwner(ctx sdk.Context, key string) sdk.AccAddress {
	receipt, err := k.GetReceipt(ctx, key)
	if err != nil {
		return nil
	}
	return receipt.Creator
}


// Check if the key exists in the store
func (k Keeper) ReceiptExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ReceiptPrefix + key))
}
