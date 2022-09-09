package keeper

import (
	"context"
	"errors"
	"time"

	"github.com/charleenfei/simple-dex/simple-dex/x/simpledex/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

// The demo swap function is a very simplified version of a dex
// It will simply burn the offer coins and mint 2x of the offer in the preferred denomination
// IF 2x of the offer is greater than the minimum ask
// IF it is not, then the message will error
// If the transfer is successful, then it will transfer the final tokens to the receiver specified in the message
func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// this is the amount the simpledex will give for the provided offer
	exchangeAmount := msg.Offer.Amount.Add(msg.Offer.Amount)

	// TODO: Handling the message
	// check if the offer is enough to swap for the minAsk
	if exchangeAmount.LT(msg.MinAsk.Amount) {
		return &types.MsgSwapResponse{}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "offer amount must be at least half of minimum ask amount")
	}

	// the following logic mocks the functionality of a dex
	err := k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.MustAccAddressFromBech32(msg.Sender), types.ModuleName, sdk.NewCoins(msg.Offer))
	if err != nil {
		return &types.MsgSwapResponse{}, err
	}

	// burn offer
	err = k.Keeper.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Offer))
	if err != nil {
		return &types.MsgSwapResponse{}, err
	}

	// mint 2x offer in preferred denomination
	exchangeCoin := sdk.NewCoin(msg.MinAsk.Denom, exchangeAmount)
	err = k.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(exchangeCoin))

	var (
		nextSeq uint64
		found   bool
	)

	// send coins to receiver on our own chain or on different chain
	if msg.PortId != "" && msg.ChannelId != "" {
		// if portID/channelID defined:
		// send the "exchanged" coins to the receiver in the message through transfer
		senderAddr := k.Keeper.accountKeeper.GetModuleAddress(types.ModuleName)
		nextSeq, found = k.channelKeeper.GetNextSequenceSend(ctx, msg.PortId, msg.ChannelId)
		if found {
			return &types.MsgSwapResponse{}, errors.New("sequence is not found")
		}
		err = k.Keeper.transferKeeper.SendTransfer(ctx, msg.PortId, msg.ChannelId, exchangeCoin, senderAddr, msg.Receiver, clienttypes.Height{}, uint64(ctx.BlockTime().Add(time.Hour).UnixNano()))
		if err != nil {
			return &types.MsgSwapResponse{}, err
		}
	} else if msg.PortId == "" && msg.ChannelId == "" {
		// if portID/channelID are empty
		// do a regular bank send to the receiver on same chain
		err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(msg.Receiver), sdk.NewCoins(exchangeCoin))
		if err != nil {
			return &types.MsgSwapResponse{}, err
		}
	} else {
		// malformed request
		return &types.MsgSwapResponse{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "portID and channelID must both be defined or both be empty")
	}

	return &types.MsgSwapResponse{Sequence: nextSeq}, nil
}
