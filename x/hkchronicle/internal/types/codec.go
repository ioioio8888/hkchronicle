package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetEvent{}, "hkchronicle/SetEvent", nil)
	cdc.RegisterConcrete(MsgBuyEvent{}, "hkchronicle/BuyEvent", nil)
	cdc.RegisterConcrete(MsgDeleteEvent{}, "hkchronicle/DeleteEvent", nil)
}
