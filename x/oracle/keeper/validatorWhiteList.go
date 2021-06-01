package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/Sifchain/sifnode/x/oracle/types"
)

// GetAllWhiteList set the validator list for a network.
func (k Keeper) GetAllWhiteList(ctx sdk.Context) types.ValidatorWhiteList {
	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator()

	key := networkDescriptor.GetPrefix()
	store.Set(key, k.cdc.MustMarshalBinaryBare(&validatorList))
}

// SetOracleWhiteList set the validator list for a network.
func (k Keeper) SetOracleWhiteList(ctx sdk.Context, networkDescriptor types.NetworkDescriptor, validatorList types.ValidatorWhiteList) {
	store := ctx.KVStore(k.storeKey)
	key := networkDescriptor.GetPrefix()
	store.Set(key, k.cdc.MustMarshalBinaryBare(&validatorList))
}

// ExistsOracleWhiteList check if the key exist
func (k Keeper) ExistsOracleWhiteList(ctx sdk.Context, networkDescriptor types.NetworkDescriptor) bool {
	key := networkDescriptor.GetPrefix()
	return k.Exists(ctx, key)
}

// GetOracleWhiteList return validator list
func (k Keeper) GetOracleWhiteList(ctx sdk.Context, networkDescriptor types.NetworkDescriptor) types.ValidatorWhiteList {
	store := ctx.KVStore(k.storeKey)
	// key := types.WhiteListValidatorPrefix
	key := networkDescriptor.GetPrefix()
	bz := store.Get(key)
	validators := &types.ValidatorWhiteList{}
	k.cdc.MustUnmarshalBinaryBare(bz, validators)

	return *validators
}

// GetAllValidators return validator list
func (k Keeper) GetAllValidators(ctx sdk.Context, networkDescriptor types.NetworkDescriptor) []sdk.ValAddress {
	store := ctx.KVStore(k.storeKey)
	// key := types.WhiteListValidatorPrefix
	key := networkDescriptor.GetPrefix()
	bz := store.Get(key)
	valAddresses := &stakingtypes.ValAddresses{}
	k.cdc.MustUnmarshalBinaryBare(bz, valAddresses)

	vl := make([]sdk.ValAddress, len(valAddresses.Addresses))
	for i, entry := range valAddresses.Addresses {
		addr, err := sdk.ValAddressFromBech32(entry)
		if err != nil {
			panic(err)
		}
		vl[i] = addr
	}

	return vl
}

// ValidateAddress is a validator in whitelist
func (k Keeper) ValidateAddress(ctx sdk.Context, networkDescriptor types.NetworkDescriptor, address sdk.ValAddress) bool {
	if !k.ExistsOracleWhiteList(ctx, networkDescriptor) {
		return false
	}
	valList := k.GetAllValidators(ctx, networkDescriptor)

	for _, validator := range valList {
		if validator.Equals(address) {
			return true
		}
	}
	return false
}

// UpdateOracleWhiteList validator's power
func (k Keeper) UpdateOracleWhiteList(ctx sdk.Context, networkDescriptor types.NetworkDescriptor, validator sdk.ValAddress, power uint32) {
	valList := k.GetOracleWhiteList(ctx, networkDescriptor)
	valList.GetWhiteList()[validator.String()] = power
	k.SetOracleWhiteList(ctx, networkDescriptor, valList)
}
