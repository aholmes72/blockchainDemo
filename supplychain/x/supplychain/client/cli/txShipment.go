package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/foo/supplychain/x/supplychain/types"
)

func GetCmdCreateShipment(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-shipment [ShippingNumber] [ShipTo] [Item] [ShipDate]",
		Short: "Creates a new shipment",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsShippingNumber := string(args[0] )
			argsShipTo := string(args[1] )
			argsItem := string(args[2] )
			argsShipDate := string(args[3] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateShipment(cliCtx.GetFromAddress(), string(argsShippingNumber), string(argsShipTo), string(argsItem), string(argsShipDate))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetShipment(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-shipment [id]  [ShippingNumber] [ShipTo] [Item] [ShipDate]",
		Short: "Set a new shipment",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsShippingNumber := string(args[1])
			argsShipTo := string(args[2])
			argsItem := string(args[3])
			argsShipDate := string(args[4])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetShipment(cliCtx.GetFromAddress(), id, string(argsShippingNumber), string(argsShipTo), string(argsItem), string(argsShipDate))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteShipment(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-shipment [id]",
		Short: "Delete a new shipment by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteShipment(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
