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

	NewMsgBuyEvent    = types.NewMsgBuyEvent
	NewMsgSetEvent    = types.NewMsgSetEvent
	NewMsgDeleteEvent = types.NewMsgDeleteEvent
	NewEvent          = types.NewEvent

	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper               = keeper.Keeper
	MsgSetEvent          = types.MsgSetEvent
	MsgBuyEvent          = types.MsgBuyEvent
	MsgDeleteEvent       = types.MsgDeleteEvent
	MsgStakeEvent        = types.MsgStakeEvent
	MsgUnStakeEvent      = types.MsgUnStakeEvent
	QueryResEventResolve = types.QueryResEventResolve
	QueryResAllEvents    = types.QueryResAllEvents
	Event                = types.Event
)
