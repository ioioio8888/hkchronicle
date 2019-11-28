package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ioioio8888/hkchronicle/x/hkchronicle/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the hkchronicle Querier
const (
	QueryResolve      = "resolve"
	QueryWhois        = "whois"
	QueryNames        = "names"
	QueryEventResolve = "eresolve"
	QueryWhoseEvent   = "whoseevent"
	QueryAllEvents    = "allevents"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		// case QueryResolve:
		// 	return queryResolve(ctx, path[1:], req, keeper)
		// case QueryWhois:
		// 	return queryWhois(ctx, path[1:], req, keeper)
		// case QueryNames:
		// 	return queryNames(ctx, req, keeper)
		case QueryEventResolve:
			return queryEventResolve(ctx, path[1:], req, keeper)
		case QueryWhoseEvent:
			return queryWhoseEvent(ctx, path[1:], req, keeper)
		case QueryAllEvents:
			return queryAllEvents(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown hkchronicle query endpoint")
		}
	}
}

// // nolint: unparam
// func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
// 	value := keeper.ResolveName(ctx, path[0])

// 	if value == "" {
// 		return []byte{}, sdk.ErrUnknownRequest("could not resolve name")
// 	}

// 	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: value})
// 	if err != nil {
// 		panic("could not marshal result to JSON")
// 	}

// 	return res, nil
// }

// // nolint: unparam
// func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
// 	whois := keeper.GetWhois(ctx, path[0])

// 	res, err := codec.MarshalJSONIndent(keeper.cdc, whois)
// 	if err != nil {
// 		panic("could not marshal result to JSON")
// 	}

// 	return res, nil
// }

// func queryNames(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
// 	var namesList types.QueryResNames

// 	iterator := keeper.GetNamesIterator(ctx)

// 	for ; iterator.Valid(); iterator.Next() {
// 		namesList = append(namesList, string(iterator.Key()))
// 	}

// 	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
// 	if err != nil {
// 		panic("could not marshal result to JSON")
// 	}

// 	return res, nil
// }

// nolint: unparam
func queryEventResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	value := keeper.ResolveEvent(ctx, path[0])

	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResEventResolve{Value: value})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

// nolint: unparam
func queryWhoseEvent(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	whoseevent := keeper.GetWhoseEvent(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, whoseevent)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryAllEvents(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var namesList types.QueryResAllEvents

	iterator := keeper.GetEventsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		namesList = append(namesList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
