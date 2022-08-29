package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"simple-dex/x/simpledex/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper  types.AccountKeeper
		bankKeeper     types.BankKeeper
		transferKeeper types.TransferKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	transferKeeper types.TransferKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		paramstore:     ps,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		transferKeeper: transferKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
