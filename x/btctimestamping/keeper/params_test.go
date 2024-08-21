package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "btc-timestamping/testutil/keeper"
	"btc-timestamping/x/btctimestamping/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BtctimestampingKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
