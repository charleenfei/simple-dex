package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "simple-dex/testutil/keeper"
	"simple-dex/x/simpledex/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SimpledexKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
