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

const (
	flagPortID    = "port-id"
	flagChannelID = "channel-id"
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
			argReceiver := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			offerCoin, err := sdk.ParseCoinNormalized(argOffer)
			if err != nil {
				return err
			}
			askCoin, err := sdk.ParseCoinNormalized(argMinAsk)

			portID, err := cmd.Flags().GetString(flagPortID)
			if err != nil {
				return err
			}
			channelID, err := cmd.Flags().GetString(flagChannelID)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwap(
				clientCtx.GetFromAddress().String(),
				offerCoin,
				askCoin,
				portID,
				channelID,
				argReceiver,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagPortID, "", "PortID to send exchanged coins on")
	cmd.Flags().String(flagChannelID, "", "ChannelID to send exchanged coins on")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
