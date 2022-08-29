package simpledex_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "simple-dex/testutil/keeper"
	"simple-dex/testutil/nullify"
	"simple-dex/x/simpledex"
	"simple-dex/x/simpledex/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SimpledexKeeper(t)
	simpledex.InitGenesis(ctx, *k, genesisState)
	got := simpledex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
