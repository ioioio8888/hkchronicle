package hkchronicle

import (
	"fmt"

	"github.com/ioioio8888/hkchronicle/x/hkchronicle/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "hkchronicle" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetEvent:
			return handleMsgSetEvent(ctx, keeper, msg)
		case MsgBuyEvent:
			return handleMsgBuyEvent(ctx, keeper, msg)
		case MsgDeleteEvent:
			return handleMsgDeleteEvent(ctx, keeper, msg)
		case MsgStakeEvent:
			return handleMsgStakeEvent(ctx, keeper, msg)
		case MsgUnStakeEvent:
			return handleMsgUnStakeEvent(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized hkchronicle Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set event
func handleMsgSetEvent(ctx sdk.Context, keeper Keeper, msg MsgSetEvent) sdk.Result {
	if !msg.Owner.Equals(keeper.GetEventOwner(ctx, msg.Event)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetThisEvent(ctx, msg.Event, msg.Value, msg.Time) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                                      // return
}

// Handle a message to buy event
func handleMsgBuyEvent(ctx sdk.Context, keeper Keeper, msg MsgBuyEvent) sdk.Result {
	if keeper.GetEventPrice(ctx, msg.Event).IsAllGT(msg.Bid) { // Checks if the the bid price is greater than the price paid by the current owner
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}
	if keeper.HasEventOwner(ctx, msg.Event) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetEventOwner(ctx, msg.Event), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	keeper.SetEventOwner(ctx, msg.Event, msg.Buyer)
	keeper.SetEventPrice(ctx, msg.Event, msg.Bid)
	return sdk.Result{}
}

func isStaked(slice []sdk.AccAddress, val sdk.AccAddress) bool {
	for _, item := range slice {
		if item.Equals(val) {
			return true
		}
	}
	return false
}

// Handle a message to stake event
func handleMsgStakeEvent(ctx sdk.Context, keeper Keeper, msg MsgStakeEvent) sdk.Result {
	Event := keeper.GetEvent(ctx, msg.Event)
	if isStaked(Event.Stakers, msg.Staker) {
		return sdk.ErrInvalidAddress("Staker already staked on this event").Result()
	}

	_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Staker, msg.Bid) // If so, deduct the Bid amount from the sender
	if err != nil {
		return sdk.ErrInsufficientCoins("Staker does not have enough coins").Result()

	}
	keeper.SetEventStaker(ctx, msg.Event, msg.Staker, msg.Bid, "stake")
	return sdk.Result{}
}

// Handle a message to unstake an event
func handleMsgUnStakeEvent(ctx sdk.Context, keeper Keeper, msg MsgUnStakeEvent) sdk.Result {
	Event := keeper.GetEvent(ctx, msg.Event)
	if !isStaked(Event.Stakers, msg.Staker) {
		return sdk.ErrInvalidAddress("Staker did not staked on this event").Result()
	}

	_, err := keeper.CoinKeeper.AddCoins(ctx, msg.Staker, msg.Bid) // If so, deduct the Bid amount from the sender
	if err != nil {
		return err.Result()
	}
	keeper.SetEventStaker(ctx, msg.Event, msg.Staker, msg.Bid, "unstake")
	return sdk.Result{}
}

// Handle a message to delete evnet
func handleMsgDeleteEvent(ctx sdk.Context, keeper Keeper, msg MsgDeleteEvent) sdk.Result {
	if !keeper.IsEventPresent(ctx, msg.Event) {
		return types.ErrEventDoesNotExist(types.DefaultCodespace).Result()
	}
	if !msg.Owner.Equals(keeper.GetEventOwner(ctx, msg.Event)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}
	keeper.DeleteEvent(ctx, msg.Event)
	return sdk.Result{}
}

// Handle a message to create Fact
func handleMsgCreateFact(ctx sdk.Context, keeper Keeper, msg types.MsgCreateFact) sdk.Result {
	// Checks if the the bid price is greater than the price paid by the current owner
	if keeper.GetPrice(ctx, msg.Title).IsAllGT(msg.Bid) {
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}
	if keeper.HasCreator(ctx, msg.Title) {
		return sdk.ErrInternal("Same Title exists!").Result()
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Creator, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}

	fact := keeper.GetFact(ctx, msg.Title)
	fact.Title = msg.Title
	fact.Time = msg.Time
	fact.Place = msg.Place
	fact.Description = msg.Description
	fact.Creator = msg.Creator
	keeper.SetFact(ctx, fact)
	return sdk.Result{}
}

// Handle a message to edit Fact
func handleMsgEditFact(ctx sdk.Context, keeper Keeper, msg types.MsgEditFact) sdk.Result {

	if !msg.Editor.Equals(keeper.GetCreator(ctx, msg.Title)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Editor is not the Creator").Result() // If not, throw an error
	}

	fact := keeper.GetFact(ctx, msg.Title)
	fact.Title = msg.Title
	fact.Time = msg.Time
	fact.Place = msg.Place
	fact.Description = msg.Description
	keeper.SetFact(ctx, fact)
	return sdk.Result{}
}
