package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSwap = "swap"

var _ sdk.Msg = &MsgSwap{}

func NewMsgSwap(sender string, offer sdk.Coin, minAsk sdk.Coin, portID, channelID, receiver string) *MsgSwap {
	return &MsgSwap{
		Sender:    sender,
		Offer:     offer,
		MinAsk:    minAsk,
		PortId:    portID,
		ChannelId: channelID,
		Receiver:  receiver,
	}
}

func (msg *MsgSwap) Route() string {
	return RouterKey
}

func (msg *MsgSwap) Type() string {
	return TypeMsgSwap
}

func (msg *MsgSwap) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSwap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	if strings.TrimSpace(msg.Receiver) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "receiver is empty")
	}
	if msg.Offer.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "offer coins is zero")
	}
	if msg.MinAsk.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "min ask is zero coins")
	}
	return nil
}
