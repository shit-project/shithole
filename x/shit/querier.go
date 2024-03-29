package rand

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// Query endpoints
const (
	QuerySeedInfo  = "seed"
	QueryRoundInfo = "round"
	QueryRoundIDs  = "round_ids"
)

// NewQuerier -
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {

		switch path[0] {
		//case QuerySeedInfo:
		//	return querySeedInfo(ctx, path[1:], req, keeper)

		case QueryRoundInfo:
			return queryRoundInfo(ctx, path[1:], req, keeper)

		case QueryRoundIDs:
			return queryRoundIDs(ctx, req, keeper)

		default:
			return nil, sdk.ErrUnknownRequest("Unknown shit query endpoint")
		}
	}
}

/*
// querySeedInfo -
func querySeedInfo(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	height := path[0]

	seed := keeper.GetSeed(ctx, height)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, seed)
	if err2 != nil {
		panic("Cannot marshal JSON")
	}

	return bz, nil
}
*/

// queryRoundInfo -
func queryRoundInfo(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]

	round := keeper.GetRound(ctx, id)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, round)
	if err2 != nil {
		panic("Cannot marshal JSON")
	}

	return bz, nil
}

// queryRoundIDs -
func queryRoundIDs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var roundIDs QueryResRoundIDs

	iterator := keeper.GetIDsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		roundIDs = append(roundIDs, id)
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, roundIDs)
	if err2 != nil {
		panic("Cannot marshal JSON.")
	}

	return bz, nil

}
