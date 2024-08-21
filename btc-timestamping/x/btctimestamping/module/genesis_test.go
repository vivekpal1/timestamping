package btctimestamping_test

import (
	"testing"

	keepertest "btc-timestamping/testutil/keeper"
	"btc-timestamping/testutil/nullify"
	btctimestamping "btc-timestamping/x/btctimestamping/module"
	"btc-timestamping/x/btctimestamping/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BtctimestampingKeeper(t)
	btctimestamping.InitGenesis(ctx, k, genesisState)
	got := btctimestamping.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
