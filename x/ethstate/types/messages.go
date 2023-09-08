package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/ingenuity-build/multierror"
)

const (
	TypeMsgSubmitSlotData     = "submit_slot_data"
	TypeMsgGetSlotDataFromEth = "get_slot_data_from_eth"
)

var (
	_ sdk.Msg            = &MsgSubmitSlotData{}
	_ sdk.Msg            = &MsgGetSlotDataFromEth{}
	_ legacytx.LegacyMsg = &MsgSubmitSlotData{}
	_ legacytx.LegacyMsg = &MsgGetSlotDataFromEth{}
)

// NewMsgSubmitSlotData - construct a msg to submit slot data
func NewMsgSubmitSlotData(slotData SlotData, fromAddress sdk.Address) *MsgSubmitSlotData {
	return &MsgSubmitSlotData{SlotData: &slotData, FromAddress: fromAddress.String()}
}

// Route Implements Msg.
func (msg MsgSubmitSlotData) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgSubmitSlotData) Type() string { return TypeMsgSubmitSlotData }

// ValidateBasic Implements Msg.
func (msg MsgSubmitSlotData) ValidateBasic() error {
	errs := make(map[string]error)

	// check from address
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		errs["FromAddress"] = err
	}

	// check slot data
	if msg.SlotData == nil {
		errs["SlotData"] = ErrSlotDataNil
	} else if err = msg.SlotData.Validate(); err != nil {
		errs["SlotData"] = err
	}

	return multierror.New(errs)
}

func (slotData *SlotData) Validate() error {
	if slotData.Height <= 0 || slotData.ContractAddress == "" || slotData.Slot == "" || slotData.Proof == nil {
		return ErrSlotDataInvalid
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSubmitSlotData) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg MsgSubmitSlotData) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// NewMsgGetSlotDataFromEth - construct a msg to get slot data from eth
func NewMsgGetSlotDataFromEth(contractAddress string, slot string, height uint64) *MsgGetSlotDataFromEth {
	return &MsgGetSlotDataFromEth{ContractAddress: contractAddress, Height: height, Slot: slot}
}

// Route Implements Msg.
func (msg MsgGetSlotDataFromEth) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgGetSlotDataFromEth) Type() string { return TypeMsgGetSlotDataFromEth }

// ValidateBasic Implements Msg.
func (msg MsgGetSlotDataFromEth) ValidateBasic() error {
	if msg.Height <= 0 || msg.ContractAddress == "" || msg.Slot == "" {
		return ErrSlotDataInvalid
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgGetSlotDataFromEth) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners Implements Msg.
func (msg MsgGetSlotDataFromEth) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
