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

// Handle a message to stake event
func handleMsgStakeEvent(ctx sdk.Context, keeper Keeper, msg MsgStakeEvent) sdk.Result {
	_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Staker, msg.Bid) // If so, deduct the Bid amount from the sender
	if err != nil {
		return sdk.ErrInsufficientCoins("Staker does not have enough coins").Result()

	}
	keeper.SetEventStaker(ctx, msg.Event, msg.Staker, msg.Bid)
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
