package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetEvent defines a SetEvent message
type MsgSetEvent struct {
	Event string         `json:"event"`
	Time  int64          `json:"time"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetEvent is a constructor function for MsgSetEvent
func NewMsgSetEvent(event string, time int64, value string, owner sdk.AccAddress) MsgSetEvent {
	return MsgSetEvent{
		Event: event,
		Time:  time,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetEvent) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetEvent) Type() string { return "set_event" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetEvent) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Event) == 0 || len(msg.Value) == 0 || msg.Time == 0 {
		return sdk.ErrUnknownRequest("Event, Time and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetEvent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgBuyEvent defines the BuyEvent message
type MsgBuyEvent struct {
	Event string         `json:"event"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

// NewMsgBuyEvent is the constructor function for MsgBuyEvent
func NewMsgBuyEvent(event string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyEvent {
	return MsgBuyEvent{
		Event: event,
		Bid:   bid,
		Buyer: buyer,
	}
}

// Route should return the name of the module
func (msg MsgBuyEvent) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyEvent) Type() string { return "buy_event" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyEvent) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Event) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyEvent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

// MsgDeleteEvent defines a DeleteEvent message
type MsgDeleteEvent struct {
	Event string         `json:"event"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgDeleteEvent is a constructor function for MsgDeleteEvent
func NewMsgDeleteEvent(event string, owner sdk.AccAddress) MsgDeleteEvent {
	return MsgDeleteEvent{
		Event: event,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgDeleteEvent) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteEvent) Type() string { return "delete_event" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteEvent) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Event) == 0 {
		return sdk.ErrUnknownRequest("Event cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteEvent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgBuyEvent defines the BuyEvent message
type MsgStakeEvent struct {
	Event  string         `json:"event"`
	Bid    sdk.Coins      `json:"bid"`
	Staker sdk.AccAddress `json:"staker"`
}

// NewMsgBuyEvent is the constructor function for MsgBuyEvent
func NewMsgStakeEvent(event string, bid sdk.Coins, staker sdk.AccAddress) MsgStakeEvent {
	return MsgStakeEvent{
		Event:  event,
		Bid:    bid,
		Staker: staker,
	}
}

// Route should return the name of the module
func (msg MsgStakeEvent) Route() string { return RouterKey }

// Type should return the action
func (msg MsgStakeEvent) Type() string { return "stake_event" }

// ValidateBasic runs stateless checks on the message
func (msg MsgStakeEvent) ValidateBasic() sdk.Error {
	if msg.Staker.Empty() {
		return sdk.ErrInvalidAddress(msg.Staker.String())
	}
	if len(msg.Event) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgStakeEvent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgStakeEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Staker}
}

// MsgBuyEvent defines the BuyEvent message
type MsgUnStakeEvent struct {
	Event  string         `json:"event"`
	Bid    sdk.Coins      `json:"bid"`
	Staker sdk.AccAddress `json:"staker"`
}

// NewMsgBuyEvent is the constructor function for MsgBuyEvent
func NewMsgUnStakeEvent(event string, bid sdk.Coins, staker sdk.AccAddress) MsgUnStakeEvent {
	return MsgUnStakeEvent{
		Event:  event,
		Bid:    bid,
		Staker: staker,
	}
}

// Route should return the name of the module
func (msg MsgUnStakeEvent) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUnStakeEvent) Type() string { return "unstake_event" }

// ValidateBasic runs stateless checks on the message
func (msg MsgUnStakeEvent) ValidateBasic() sdk.Error {
	if msg.Staker.Empty() {
		return sdk.ErrInvalidAddress(msg.Staker.String())
	}
	if len(msg.Event) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnStakeEvent) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUnStakeEvent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Staker}
}
