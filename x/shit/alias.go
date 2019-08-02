package rand

import (
	"github.com/shit-project/shithole/x/rand/types"
)

const (
	// ModuleName -
	ModuleName = types.ModuleName

	// RouterKey -
	RouterKey = types.RouterKey

	// StoreKey -
	StoreKey = types.StoreKey
)

var (
	// NewMsgNewRound -
	NewMsgNewRound = types.NewMsgNewRound

	// NewMsgDeployNonce -
	NewMsgDeployNonce = types.NewMsgDeployNonce

	// NewMsgAddTargets -
	NewMsgAddTargets = types.NewMsgAddTargets

	// NewMsgUpdateTargets -
	NewMsgUpdateTargets = types.NewMsgUpdateTargets

	// ModuleCdc -
	ModuleCdc = types.ModuleCdc

	// RegisterCodec -
	RegisterCodec = types.RegisterCodec
)

type (
	// MsgNewRound -
	MsgNewRound = types.MsgNewRound

	// MsgDeployNonce -
	MsgDeployNonce = types.MsgDeployNonce

	// MsgAddTargets -
	MsgAddTargets = types.MsgAddTargets

	// MsgUpdateTargets -
	MsgUpdateTargets = types.MsgUpdateTargets

	// Seed -
	//Seed = types.Seed

	// Round -
	Round = types.Round

	// Rounds -
	Rounds = types.Rounds

	// QueryResRoundIDs -
	QueryResRoundIDs = types.QueryResRoundIDs
)
