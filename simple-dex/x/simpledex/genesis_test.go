package simpledex_test

import (
	"testing"

	"github.com/charleenfei/simple-dex/simple-dex/x/simpledex"

	"github.com/charleenfei/simple-dex/simple-dex/x/simpledex/types"

	"github.com/charleenfei/simple-dex/simple-dex/testutil/nullify"

	keepertest "github.com/charleenfei/simple-dex/simple-dex/testutil/keeper"
	"github.com/stretchr/testify/require"
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
