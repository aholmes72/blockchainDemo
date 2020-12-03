package keeper

import (
  // this line is used by starport scaffolding # 1
	"github.com/foo/supplychain/x/supplychain/types"
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for supplychain clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListShipment:
			return listShipment(ctx, k)
		case types.QueryGetShipment:
			return getShipment(ctx, path[1:], k)
		case types.QueryListReceipt:
			return listReceipt(ctx, k)
		case types.QueryGetReceipt:
			return getReceipt(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown supplychain query endpoint")
		}
	}
}
