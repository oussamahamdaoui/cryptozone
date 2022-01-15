package cryptozone_test

import (
	"testing"

	keepertest "github.com/oussamahamdaoui/cryptozone/testutil/keeper"
	"github.com/oussamahamdaoui/cryptozone/testutil/nullify"
	"github.com/oussamahamdaoui/cryptozone/x/cryptozone"
	"github.com/oussamahamdaoui/cryptozone/x/cryptozone/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CryptozoneKeeper(t)
	cryptozone.InitGenesis(ctx, *k, genesisState)
	got := cryptozone.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
