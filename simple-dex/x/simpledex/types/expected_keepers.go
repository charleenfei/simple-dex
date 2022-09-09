package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
	GetModuleAddress(moduleName string) sdk.AccAddress
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// ChannelKeeper defines the expected interface needed to retrieve next sequence send.
type ChannelKeeper interface {
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
}

// TransferKeeper defines the expected interface needed to send transfer packets
type TransferKeeper interface {
	SendTransfer(ctx sdk.Context, sourcePort, sourceChannel string, token sdk.Coin, sender sdk.AccAddress,
		receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
	) error
}
