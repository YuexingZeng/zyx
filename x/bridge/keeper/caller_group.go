package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"zyx/x/bridge/types"
)

// SetCallerGroup set a specific callerGroup in the store from its index
func (k Keeper) SetCallerGroup(ctx sdk.Context, callerGroup types.CallerGroup) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CallerGroupKeyPrefix))
	b := k.cdc.MustMarshal(&callerGroup)
	store.Set(types.CallerGroupKey(
		callerGroup.Name,
	), b)
}

// GetCallerGroup returns a callerGroup from its index
func (k Keeper) GetCallerGroup(
	ctx sdk.Context,
	name string,

) (val types.CallerGroup, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CallerGroupKeyPrefix))

	b := store.Get(types.CallerGroupKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCallerGroup removes a callerGroup from the store
func (k Keeper) RemoveCallerGroup(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CallerGroupKeyPrefix))
	store.Delete(types.CallerGroupKey(
		name,
	))
}

// GetAllCallerGroup returns all callerGroup
func (k Keeper) GetAllCallerGroup(ctx sdk.Context) (list []types.CallerGroup) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CallerGroupKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CallerGroup
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
