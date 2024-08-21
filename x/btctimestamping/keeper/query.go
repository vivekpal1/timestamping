package keeper

import (
	"btc-timestamping/x/btctimestamping/types"
)

var _ types.QueryServer = Keeper{}
