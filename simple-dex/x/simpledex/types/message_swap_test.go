package types

import (
	"testing"

	"simple-dex/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSwap_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSwap
		err  error
	}{
		{
			name: "valid address",
			msg: MsgSwap{
				Sender:    sample.AccAddress(),
				Offer:     sdk.NewCoin("uatom", sdk.NewInt(100)),
				MinAsk:    sdk.NewCoin("uosmo", sdk.NewInt(50)),
				PortId:    "transfer",
				ChannelId: "channel-1",
				Receiver:  "receiverAddr",
			},
			err: nil,
		},
		{
			name: "invalid address",
			msg: MsgSwap{
				Sender:    "invalid_address",
				Offer:     sdk.NewCoin("uatom", sdk.NewInt(100)),
				MinAsk:    sdk.NewCoin("uosmo", sdk.NewInt(50)),
				PortId:    "transfer",
				ChannelId: "channel-1",
				Receiver:  "receiverAddr",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "offer is zero coins",
			msg: MsgSwap{
				Sender:    sample.AccAddress(),
				Offer:     sdk.NewCoin("uatom", sdk.NewInt(0)),
				MinAsk:    sdk.NewCoin("uosmo", sdk.NewInt(50)),
				PortId:    "transfer",
				ChannelId: "channel-1",
				Receiver:  "receiverAddr",
			},
			err: sdkerrors.ErrInsufficientFunds,
		},
		{
			name: "minAsk is zero coins",
			msg: MsgSwap{
				Sender:    sample.AccAddress(),
				Offer:     sdk.NewCoin("uatom", sdk.NewInt(20)),
				MinAsk:    sdk.NewCoin("uosmo", sdk.NewInt(0)),
				PortId:    "transfer",
				ChannelId: "channel-1",
				Receiver:  "receiverAddr",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "empty receiver",
			msg: MsgSwap{
				Sender:    sample.AccAddress(),
				Offer:     sdk.NewCoin("uatom", sdk.NewInt(100)),
				MinAsk:    sdk.NewCoin("uosmo", sdk.NewInt(50)),
				PortId:    "transfer",
				ChannelId: "channel-1",
				Receiver:  "",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
