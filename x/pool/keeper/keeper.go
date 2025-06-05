package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/aether-proj/aether/x/pool/types"
)

type DepositorIndexes struct {
	Depositor *indexes.Unique[sdk.AccAddress, uint64, types.Deposit]
}

func (a DepositorIndexes) IndexesList() []collections.Index[uint64, types.Deposit] {
	return []collections.Index[uint64, types.Deposit]{a.Depositor}
}

func NewDepositorIndexes(sb *collections.SchemaBuilder) DepositorIndexes {
	return DepositorIndexes{
		Depositor: indexes.NewUnique(
			sb, types.DepositorIndexPrefix, "deposits_by_address",
			sdk.AccAddressKey, collections.Uint64Key,
			func(_ uint64, v types.Deposit) (sdk.AccAddress, error) {
				return sdk.AccAddress(v.Depositor), nil
			},
		),
	}
}

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec

	authority []byte

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper

	Schema    collections.Schema
	Params    collections.Item[types.Params]
	PoolId    collections.Sequence
	Pools     collections.Map[uint64, types.Pool]
	DepositId collections.Sequence
	Deposits  *collections.IndexedMap[uint64, types.Deposit, DepositorIndexes]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,
	accountKeeper types.AccountKeeper,
) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService:  storeService,
		cdc:           cdc,
		addressCodec:  addressCodec,
		authority:     authority,
		AccountKeeper: accountKeeper,
		Params:        collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		PoolId:        collections.NewSequence(sb, types.PoolIdKey, "pool_id"),
		Pools:         collections.NewMap(sb, types.PoolsKey, "pools", collections.Uint64Key, codec.CollValue[types.Pool](cdc)),
		DepositId:     collections.NewSequence(sb, types.DepositIdKey, "deposit_id"),
		Deposits:      collections.NewIndexedMap(sb, types.DepositsKey, "deposits", collections.Uint64Key, codec.CollValue[types.Deposit](cdc), NewDepositorIndexes(sb)),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

func (k Keeper) GetAuthority() []byte {
	return k.authority
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetParams(ctx context.Context) types.Params {
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}
	return params
}
