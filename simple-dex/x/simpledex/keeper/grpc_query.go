package keeper

import (
	"simple-dex/x/simpledex/types"
)

var _ types.QueryServer = Keeper{}
