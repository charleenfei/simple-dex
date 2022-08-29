package cli

import (
	"strconv"

	"simple-dex/x/simpledex/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [offer] [min-ask] [portID] [channelID] [receiver]",
		Short: "Broadcast message swap",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOffer := args[0]
			argMinAsk := args[1]
			argPortID := args[2]
			argChannelID := args[3]
			argReceiver := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			offerCoin, err := sdk.ParseCoinNormalized(argOffer)
			if err != nil {
				return err
			}
			askCoin, err := sdk.ParseCoinNormalized(argMinAsk)

			msg := types.NewMsgSwap(
				clientCtx.GetFromAddress().String(),
				offerCoin,
				askCoin,
				argPortID,
				argChannelID,
				argReceiver,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
