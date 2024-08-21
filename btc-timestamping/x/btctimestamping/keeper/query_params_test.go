package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "btc-timestamping/testutil/keeper"
	"btc-timestamping/x/btctimestamping/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.BtctimestampingKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
