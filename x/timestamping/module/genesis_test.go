package timestamping_test

import (
	"testing"

	keepertest "btc-timestamping/testutil/keeper"
	"btc-timestamping/testutil/nullify"
	timestamping "btc-timestamping/x/timestamping/module"
	"btc-timestamping/x/timestamping/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TimestampingKeeper(t)
	timestamping.InitGenesis(ctx, k, genesisState)
	got := timestamping.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
