package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc -
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec -
func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(MsgShit{}, "shit/Shit", nil)
	cdc.RegisterConcrete(MsgSorry{}, "rand/Sorry", nil)
}
