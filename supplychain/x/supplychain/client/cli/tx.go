package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/foo/supplychain/x/supplychain/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	supplychainTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	supplychainTxCmd.AddCommand(flags.PostCommands(
    // this line is used by starport scaffolding # 1
		GetCmdCreateShipment(cdc),
		GetCmdSetShipment(cdc),
		GetCmdDeleteShipment(cdc),
		GetCmdCreateReceipt(cdc),
		GetCmdSetReceipt(cdc),
		GetCmdDeleteReceipt(cdc),
	)...)

	return supplychainTxCmd
}
