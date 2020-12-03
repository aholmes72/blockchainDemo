package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/foo/supplychain/x/supplychain/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group supplychain queries under a subcommand
	supplychainQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	supplychainQueryCmd.AddCommand(
		flags.GetCommands(
      // this line is used by starport scaffolding # 1
			GetCmdListShipment(queryRoute, cdc),
			GetCmdGetShipment(queryRoute, cdc),
			GetCmdListReceipt(queryRoute, cdc),
			GetCmdGetReceipt(queryRoute, cdc),
		)...,
	)

	return supplychainQueryCmd
}
