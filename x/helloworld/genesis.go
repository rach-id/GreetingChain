package helloworld

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sweexordious/./x/helloworld/types"
	"github.com/sweexordious/x/helloworld/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
//func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
//	// TODO: Define logic for when you would like to initialize a new genesis
//}
//
//// ExportGenesis writes the current store values
//// to a genesis file, which can be imported again
//// with InitGenesis
//func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
//	// TODO: Define logic for exporting state
//	return types.NewGenesisState()
//}