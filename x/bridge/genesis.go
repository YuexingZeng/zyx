package bridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"zyx/x/bridge/keeper"
	"zyx/x/bridge/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the signerGroup
	for _, elem := range genState.SignerGroupList {
		k.SetSignerGroup(ctx, elem)
	}
	// Set all the callerGroup
	for _, elem := range genState.CallerGroupList {
		k.SetCallerGroup(ctx, elem)
	}
	// Set all the deposit
	for _, elem := range genState.DepositList {
		k.SetDeposit(ctx, elem)
	}
	// Set all the withdraw
	for _, elem := range genState.WithdrawList {
		k.SetWithdraw(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.SignerGroupList = k.GetAllSignerGroup(ctx)
	genesis.CallerGroupList = k.GetAllCallerGroup(ctx)
	genesis.DepositList = k.GetAllDeposit(ctx)
	genesis.WithdrawList = k.GetAllWithdraw(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
