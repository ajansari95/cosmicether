package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/gogoproto/proto"
)

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgSubmitSlotData{}, "cosmicether/MsgSubmitSlotData", nil)
	cdc.RegisterConcrete(&MsgGetSlotDataFromEth{}, "cosmicether/MsgGetSlotDataFromEth", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	impls := []proto.Message{&MsgSubmitSlotData{}, &MsgGetSlotDataFromEth{}}
	for _, impl := range impls {
		typeURL := "/" + proto.MessageName(impl)
		fmt.Print(typeURL)
	}

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSubmitSlotData{},
		&MsgGetSlotDataFromEth{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
