package keeper

import (
	"btc-timestamping/x/timestamping/types"
)

var _ types.QueryServer = Keeper{}
