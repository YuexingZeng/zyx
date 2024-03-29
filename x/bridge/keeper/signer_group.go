package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"zyx/x/bridge/types"
)

// SetSignerGroup set a specific signerGroup in the store from its index
func (k Keeper) SetSignerGroup(ctx sdk.Context, signerGroup types.SignerGroup) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerGroupKeyPrefix))
	b := k.cdc.MustMarshal(&signerGroup)
	store.Set(types.SignerGroupKey(
		signerGroup.Name,
	), b)
}

// GetSignerGroup returns a signerGroup from its index
func (k Keeper) GetSignerGroup(
	ctx sdk.Context,
	name string,

) (val types.SignerGroup, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerGroupKeyPrefix))

	b := store.Get(types.SignerGroupKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSignerGroup removes a signerGroup from the store
func (k Keeper) RemoveSignerGroup(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerGroupKeyPrefix))
	store.Delete(types.SignerGroupKey(
		name,
	))
}

// GetAllSignerGroup returns all signerGroup
func (k Keeper) GetAllSignerGroup(ctx sdk.Context) (list []types.SignerGroup) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerGroupKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SignerGroup
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
