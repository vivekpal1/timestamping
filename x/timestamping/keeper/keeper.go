package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/vivekpal1/timestamping/x/timestamping/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetTimestamp(ctx sdk.Context, timestamp types.Timestamp) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&timestamp)
	store.Set(GetTimestampIDBytes(timestamp.Hash), b)
}

func (k Keeper) GetTimestamp(ctx sdk