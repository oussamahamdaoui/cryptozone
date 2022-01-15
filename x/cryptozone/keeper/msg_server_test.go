package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/oussamahamdaoui/cryptozone/testutil/keeper"
	"github.com/oussamahamdaoui/cryptozone/x/cryptozone/keeper"
	"github.com/oussamahamdaoui/cryptozone/x/cryptozone/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CryptozoneKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
