package tx

import (
	"github.com/Pylons-tech/pylons/x/pylons/msgs"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

// FiatItem is the client cli command for creating item
func FiatItem() *cobra.Command {
	var msgFI = &msgs.MsgFiatItem{}

	cmd := &cobra.Command{
		Use:   "fiat-item [args]",
		Short: "create item and assign it to sender",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			byteValue, err := ReadFile(args[0])
			if err != nil {
				return err
			}
			err = msgFI.Unmarshal(byteValue)
			if err != nil {
				return err
			}

			msgFI.Sender = clientCtx.GetFromAddress().String()

			err = msgFI.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), []sdk.Msg{msgFI}...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.PersistentFlags().StringVar(&msgFI.CookbookID, "cookbookID", "", "The ID of the cookbook for this item")

	return cmd
}
