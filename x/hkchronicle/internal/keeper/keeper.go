package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/ioioio8888/hkchronicle/x/hkchronicle/internal/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the hkchronicle Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Sets the entire WhoseEvent metadata struct for an event
func (k Keeper) SetWhoseEvent(ctx sdk.Context, event string, whoseEvent types.WhoseEvent) {
	if whoseEvent.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(event), k.cdc.MustMarshalBinaryBare(whoseEvent))
}

// Gets the entire WhoseEvent metadata struct for an event
func (k Keeper) GetWhoseEvent(ctx sdk.Context, event string) types.WhoseEvent {
	store := ctx.KVStore(k.storeKey)
	if !k.IsEventPresent(ctx, event) {
		return types.NewWhoseEvent()
	}
	bz := store.Get([]byte(event))
	var whoseEvent types.WhoseEvent
	k.cdc.MustUnmarshalBinaryBare(bz, &whoseEvent)
	return whoseEvent
}

// Deletes the entire WhoseEevnt metadata struct for an event
func (k Keeper) DeleteWhoseEvent(ctx sdk.Context, event string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(event))
}

// ResolveEvent - returns the info that the event resolves to
func (k Keeper) ResolveEvent(ctx sdk.Context, event string) string {
	return k.GetWhoseEvent(ctx, event).Value
}

// SetEvent - sets the value and time that a event resolves to
func (k Keeper) SetEvent(ctx sdk.Context, event string, value string, time int64) {
	whoseEvent := k.GetWhoseEvent(ctx, event)
	whoseEvent.Value = value
	whoseEvent.Time = time
	k.SetWhoseEvent(ctx, event, whoseEvent)
}

// HasEventOwner - returns whether or not the event already has an owner
func (k Keeper) HasEventOwner(ctx sdk.Context, event string) bool {
	return !k.GetWhoseEvent(ctx, event).Owner.Empty()
}

// GetEventOwner - get the current owner of an event
func (k Keeper) GetEventOwner(ctx sdk.Context, event string) sdk.AccAddress {
	return k.GetWhoseEvent(ctx, event).Owner
}

// SetEventOwner - sets the current owner of an event
func (k Keeper) SetEventOwner(ctx sdk.Context, event string, owner sdk.AccAddress) {
	whoseEvent := k.GetWhoseEvent(ctx, event)
	whoseEvent.Owner = owner
	k.SetWhoseEvent(ctx, event, whoseEvent)
}

// GetEventPrice - gets the current price of an event
func (k Keeper) GetEventPrice(ctx sdk.Context, event string) sdk.Coins {
	return k.GetWhoseEvent(ctx, event).Price
}

// SetEventPrice - sets the current price of an event
func (k Keeper) SetEventPrice(ctx sdk.Context, event string, price sdk.Coins) {
	whoseEvent := k.GetWhoseEvent(ctx, event)
	whoseEvent.Price = price
	k.SetWhoseEvent(ctx, event, whoseEvent)
}

// Check if the event is present in the store or not
func (k Keeper) IsEventPresent(ctx sdk.Context, event string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(event))
}

// Get an iterator over all events in which the keys are the events and the values are the whois
func (k Keeper) GetEventsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}
