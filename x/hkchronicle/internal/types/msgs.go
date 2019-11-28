package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// // MsgSetName defines a SetName message
// type MsgSetName struct {
// 	Name  string         `json:"name"`
// 	Value string         `json:"value"`
// 	Owner sdk.AccAddress `json:"owner"`
// }

// // NewMsgSetName is a constructor function for MsgSetName
// func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
// 	return MsgSetName{
// 		Name:  name,
// 		Value: value,
// 		Owner: owner,
// 	}
// }

// // Route should return the name of the module
// func (msg MsgSetName) Route() string { return RouterKey }

// // Type should return the action
// func (msg MsgSetName) Type() string { return "set_name" }

// // ValidateBasic runs stateless checks on the message
// func (msg MsgSetName) ValidateBasic() sdk.Error {
// 	if msg.Owner.Empty() {
// 		return sdk.ErrInvalidAddress(msg.Owner.String())
// 	}
// 	if len(msg.Name) == 0 || len(msg.Value) == 0 {
// 		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
// 	}
// 	return nil
// }

// // GetSignBytes encodes the message for signing
// func (msg MsgSetName) GetSignBytes() []byte {
// 	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
// }

// // GetSigners defines whose signature is required
// func (msg MsgSetName) GetSigners() []sdk.AccAddress {
// 	return []sdk.AccAddress{msg.Owner}
// }

// // MsgBuyName defines the BuyName message
// type MsgBuyName struct {
// 	Name  string         `json:"name"`
// 	Bid   sdk.Coins      `json:"bid"`
// 	Buyer sdk.AccAddress `json:"buyer"`
// }

// // NewMsgBuyName is the constructor function for MsgBuyName
// func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
// 	return MsgBuyName{
// 		Name:  name,
// 		Bid:   bid,
// 		Buyer: buyer,
// 	}
// }

// // Route should return the name of the module
// func (msg MsgBuyName) Route() string { return RouterKey }

// // Type should return the action
// func (msg MsgBuyName) Type() string { return "buy_name" }

// // ValidateBasic runs stateless checks on the message
// func (msg MsgBuyName) ValidateBasic() sdk.Error {
// 	if msg.Buyer.Empty() {
// 		return sdk.ErrInvalidAddress(msg.Buyer.String())
// 	}
// 	if len(msg.Name) == 0 {
// 		return sdk.ErrUnknownRequest("Name cannot be empty")
// 	}
// 	if !msg.Bid.IsAllPositive() {
// 		return sdk.ErrInsufficientCoins("Bids must be positive")
// 	}
// 	return nil
// }

// // GetSignBytes encodes the message for signing
// func (msg MsgBuyName) GetSignBytes() []byte {
// 	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
// }

// // GetSigners defines whose signature is required
// func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
// 	return []sdk.AccAddress{msg.Buyer}
// }

// // MsgDeleteName defines a DeleteName message
// type MsgDeleteName struct {
// 	Name  string         `json:"name"`
// 	Owner sdk.AccAddress `json:"owner"`
// }

// // NewMsgDeleteName is a constructor function for MsgDeleteName
// func NewMsgDeleteName(name string, owner sdk.AccAddress) MsgDeleteName {
// 	return MsgDeleteName{
// 		Name:  name,
// 		Owner: owner,
// 	}
// }

// // Route should return the name of the module
// func (msg MsgDeleteName) Route() string { return RouterKey }

// // Type should return the action
// func (msg MsgDeleteName) Type() string { return "delete_name" }

// // ValidateBasic runs stateless checks on the message
// func (msg MsgDeleteName) ValidateBasic() sdk.Error {
// 	if msg.Owner.Empty() {
// 		return sdk.ErrInvalidAddress(msg.Owner.String())
// 	}
// 	if len(msg.Name) == 0 {
// 		return sdk.ErrUnknownRequest("Name cannot be empty")
// 	}
// 	return nil
// }

// // GetSignBytes encodes the message for signing
// func (msg MsgDeleteName) GetSignBytes() []byte {
// 	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
// }

// // GetSigners defines whose signature is required
// func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
// 	return []sdk.AccAddress{msg.Owner}
// }

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
