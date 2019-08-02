package app

import (
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

// DONTCOVER

// NewShitAppUNSAFE -
func NewShitAppUNSAFE(logger log.Logger, db dbm.DB, invCheckPeriod uint) (sapp *ShitApp, keyMain, keyStaking *sdk.KVStoreKey, stakingKeeper staking.Keeper) {

	sapp = NewShitApp(logger, db, invCheckPeriod)
	return sapp, sapp.keyMain, sapp.keyStaking, sapp.stakingKeeper
}
