package tokenfactory

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"tokenfactory/x/tokenfactory/keeper"
	"tokenfactory/x/tokenfactory/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the denom
	for _, elem := range genState.DenomList {
		k.SetDenom(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.Params.Set(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}

	genesis.DenomList = k.GetAllDenom(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
