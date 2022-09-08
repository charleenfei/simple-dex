package keeper_test

import (
	"testing"

	testkeeper "github.com/charleenfei/simple-dex/simple-dex/testutil/keeper"
	"github.com/charleenfei/simple-dex/simple-dex/x/simpledex/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SimpledexKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
