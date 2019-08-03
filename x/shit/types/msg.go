package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey -
const RouterKey = ModuleName

/*
Shit
*/

// MsgShit -
type MsgShit struct {
	ID        string
	ShitType  uint8
	Amount    uint16
	Informant sdk.AccAddress
	Comment   string
}

// NewMsgShit -
func NewMsgShit(id string, shitType uint8, amount uint16, informant sdk.AccAddress, comment string) MsgShit {
	return MsgShit{
		ID:        id,
		ShitType:  shitType,
		Amount:    amount,
		Informant: informant,
		Comment:   comment,
	}
}

// Route -
func (msg MsgShit) Route() string {
	return RouterKey
}

// Type -
func (msg MsgShit) Type() string {
	return "shit"
}

// ValidateBasic -
func (msg MsgShit) ValidateBasic() sdk.Error {
	if msg.Informant.Empty() {
		return sdk.ErrInvalidAddress(msg.Informant.String())
	}

	if len(msg.ID) == 0 {
		return sdk.ErrUnknownRequest("Must specify ID, ShitType.")
	}

	if msg.Amount < 100 {
		return sdk.ErrUnknownRequest("Amount must be greater than 100.")
	}

	if msg.ShitType < 1 {
		return sdk.ErrUnknownRequest("ShitType must be greater than 0.")
	}

	return nil
}

// GetSignBytes -
func (msg MsgShit) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners -
func (msg MsgShit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Informant}
}

/*
MsgSorry
*/

// MsgSorry -
type MsgSorry struct {
	ID        string
	ShitType  uint8
	Amount    uint16
	Informant sdk.AccAddress
	Opponent  sdk.AccAddress
	Comment   string
}

// NewMsgSorry -
func NewMsgSorry(id string, owner sdk.AccAddress, nonce string) MsgSorry {
	return MsgSorry{
		ID:        id,
		ShitType:  owner,
		Amount:    nonce,
		Informant: informant,
		Opponent:  opponent,
		Comment:   comment,
	}
}

// Route -
func (msg MsgDeployNonce) Route() string {
	return "rand"
}

// Type -
func (msg MsgDeployNonce) Type() string {
	return "deploy_round"
}

// ValidateBasic -
func (msg MsgDeployNonce) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	if len(msg.ID) == 0 {
		return sdk.ErrUnknownRequest("Must specify ID, NonceHash.")
	}

	return nil
}

// GetSignBytes -
func (msg MsgDeployNonce) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners -
func (msg MsgDeployNonce) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

/*
AddTargets
*/

// MsgAddTargets -
type MsgAddTargets struct {
	ID      string
	Owner   sdk.AccAddress
	Targets []string
}

// NewMsgAddTargets -
func NewMsgAddTargets(id string, owner sdk.AccAddress, targets []string) MsgAddTargets {
	return MsgAddTargets{
		ID:      id,
		Owner:   owner,
		Targets: targets,
	}
}

// Route -
func (msg MsgAddTargets) Route() string {
	return "rand"
}

// Type -
func (msg MsgAddTargets) Type() string {
	return "add_targets"
}

// ValidateBasic - 모집단 추가 ValidateBasic
func (msg MsgAddTargets) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.ID) == 0 {
		return sdk.ErrUnknownRequest("Must specify ID.")
	}

	if msg.Targets == nil {
		return sdk.ErrUnknownRequest("Must specify Targets.")
	}

	return nil
}

// GetSignBytes -
func (msg MsgAddTargets) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners -
func (msg MsgAddTargets) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

/*
MsgUpdateTargets
*/

// MsgUpdateTargets -
type MsgUpdateTargets struct {
	ID      string
	Owner   sdk.AccAddress
	Targets []string
}

// NewMsgUpdateTargets -
func NewMsgUpdateTargets(id string, owner sdk.AccAddress, targets []string) MsgUpdateTargets {
	return MsgUpdateTargets{
		ID:      id,
		Owner:   owner,
		Targets: targets,
	}
}

// Route -
func (msg MsgUpdateTargets) Route() string {
	return "rand"
}

// Type -
func (msg MsgUpdateTargets) Type() string {
	return "update_targets"
}

// ValidateBasic -
func (msg MsgUpdateTargets) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.ID) == 0 {
		return sdk.ErrUnknownRequest("Must specify ID.")
	}

	if msg.Targets == nil {
		return sdk.ErrUnknownRequest("Must specify Targets.")
	}

	return nil
}

// GetSignBytes -
func (msg MsgUpdateTargets) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners -
func (msg MsgUpdateTargets) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
