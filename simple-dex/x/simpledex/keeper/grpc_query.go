package keeper

import (
	"github.com/charleenfei/simple-dex/simple-dex/x/simpledex/types"
)

var _ types.QueryServer = Keeper{}
