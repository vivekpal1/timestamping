package keeper

import (
	"context"

	"btc-timestamping/x/timestamping/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTimestamp(goCtx context.Context, msg *types.MsgCreateTimestamp) (*types.MsgCreateTimestampResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateTimestampResponse{}, nil
}
