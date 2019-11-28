package hkchronicle

import (
	"github.com/ioioio8888/hkchronicle/x/hkchronicle/internal/keeper"
	"github.com/ioioio8888/hkchronicle/x/hkchronicle/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper  = keeper.NewKeeper
	NewQuerier = keeper.NewQuerier
	// NewMsgBuyName    = types.NewMsgBuyName
	// NewMsgSetName    = types.NewMsgSetName
	// NewMsgDeleteName = types.NewMsgDeleteName
	// NewWhois         = types.NewWhois

	NewMsgBuyEvent    = types.NewMsgBuyEvent
	NewMsgSetEvent    = types.NewMsgSetEvent
	NewMsgDeleteEvent = types.NewMsgDeleteEvent
	NewWhoseEvent     = types.NewWhoseEvent

	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper         = keeper.Keeper
	MsgSetEvent    = types.MsgSetEvent
	MsgBuyEvent    = types.MsgBuyEvent
	MsgDeleteEvent = types.MsgDeleteEvent
	// MsgSetName           = types.MsgSetName
	// MsgBuyName           = types.MsgBuyName
	// MsgDeleteName        = types.MsgDeleteName
	QueryResEventResolve = types.QueryResEventResolve
	QueryResAllEvents    = types.QueryResAllEvents
	// QueryResResolve      = types.QueryResResolve
	// QueryResNames        = types.QueryResNames
	WhoseEvent = types.WhoseEvent
	// Whois                = types.Whois
)
