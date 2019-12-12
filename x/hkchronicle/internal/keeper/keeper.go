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

// Sets the entire Event metadata struct for an event
func (k Keeper) SetEvent(ctx sdk.Context, event string, Event types.Event) {
	if Event.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(event), k.cdc.MustMarshalBinaryBare(Event))
}

// Gets the entire Event metadata struct for an event
func (k Keeper) GetEvent(ctx sdk.Context, event string) types.Event {
	store := ctx.KVStore(k.storeKey)
	if !k.IsEventPresent(ctx, event) {
		return types.NewEvent()
	}
	bz := store.Get([]byte(event))
	var Event types.Event
	k.cdc.MustUnmarshalBinaryBare(bz, &Event)
	return Event
}

// Deletes the entire WhoseEevnt metadata struct for an event
func (k Keeper) DeleteEvent(ctx sdk.Context, event string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(event))
}

// ResolveEvent - returns the info that the event resolves to
func (k Keeper) ResolveEvent(ctx sdk.Context, event string) string {
	return k.GetEvent(ctx, event).Value
}

// SetThisEvent - sets the value and time that a event resolves to
func (k Keeper) SetThisEvent(ctx sdk.Context, event string, value string, time int64) {
	Event := k.GetEvent(ctx, event)
	Event.Value = value
	Event.Time = time
	k.SetEvent(ctx, event, Event)
}

// HasEventOwner - returns whether or not the event already has an owner
func (k Keeper) HasEventOwner(ctx sdk.Context, event string) bool {
	return !k.GetEvent(ctx, event).Owner.Empty()
}

// GetEventOwner - get the current owner of an event
func (k Keeper) GetEventOwner(ctx sdk.Context, event string) sdk.AccAddress {
	return k.GetEvent(ctx, event).Owner
}

func indexOf(element sdk.AccAddress, data []sdk.AccAddress) int {
	for k, v := range data {
		if element.Equals(v) {
			return k
		}
	}
	return -1 //not found.
}

func RemoveIndex(s []sdk.AccAddress, staker sdk.AccAddress) []sdk.AccAddress {
	index := indexOf(staker, s)
	return append(s[:index], s[index+1:]...)
}

// SetEventStaker - sets the current owner of an event
func (k Keeper) SetEventStaker(ctx sdk.Context, event string, staker sdk.AccAddress, value sdk.Coins, stakeType string) {
	Event := k.GetEvent(ctx, event)
	if stakeType == "stake" {
		Event.Stakers = append(Event.Stakers, staker)
		Event.Stake = Event.Stake.Add(value)

	} else if stakeType == "unstake" {
		Event.Stakers = RemoveIndex(Event.Stakers, staker)
		Event.Stake = Event.Stake.Sub(value)
	}
	k.SetEvent(ctx, event, Event)
}

// SetEventOwner - sets the current owner of an event
func (k Keeper) SetEventOwner(ctx sdk.Context, event string, owner sdk.AccAddress) {
	Event := k.GetEvent(ctx, event)
	Event.Owner = owner
	k.SetEvent(ctx, event, Event)
}

// GetEventPrice - gets the current price of an event
func (k Keeper) GetEventPrice(ctx sdk.Context, event string) sdk.Coins {
	return k.GetEvent(ctx, event).Price
}

// SetEventPrice - sets the current price of an event
func (k Keeper) SetEventPrice(ctx sdk.Context, event string, price sdk.Coins) {
	Event := k.GetEvent(ctx, event)
	Event.Price = price
	k.SetEvent(ctx, event, Event)
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

//-------fact-------//
// Gets the entire Fact metadata struct for a name
func (k Keeper) GetFact(ctx sdk.Context, title string) types.Fact {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(title))
	if bz == nil {
		//return an empty new fact if fact does not exist
		return types.NewFact()
	}
	var fact types.Fact
	k.cdc.MustUnmarshalBinaryBare(bz, &fact)
	return fact
}

// HasCreator - returns whether or not the title already has an creator
func (k Keeper) HasCreator(ctx sdk.Context, title string) bool {
	return !k.GetFact(ctx, title).Creator.Empty()
}

// Sets the entire Whois metadata struct for a name
func (k Keeper) SetFact(ctx sdk.Context, fact types.Fact) {
	if fact.Creator.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(fact.Title), k.cdc.MustMarshalBinaryBare(fact))
}

// GetCreator - get the creator of a fact
func (k Keeper) GetCreator(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetFact(ctx, name).Creator
}

// GetPrice - get the price of a fact
func (k Keeper) GetPrice(ctx sdk.Context, title string) sdk.Coins {
	return k.GetFact(ctx, title).Price
}
