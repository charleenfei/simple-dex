package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/charleenfei/simple-dex/simple-dex/testutil/keeper"
	"github.com/charleenfei/simple-dex/simple-dex/x/simpledex/keeper"
	"github.com/charleenfei/simple-dex/simple-dex/x/simpledex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SimpledexKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
