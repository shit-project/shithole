package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Shit -
type Shit struct {
	ID        string         `json:"id"`
	ShitType  uint8          `json:"shit_type"`
	Amount    uint64         `json:"amount"`
	Informant sdk.AccAddress `json:"informant"`
	Comment   string         `jsong:"comment"`
}

func (s Shit) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
ShitType: %d
Amount: %d
Informants: %s
Comments: %s
`, s.ID, s.ShitType, s.Amount, s.Informant, s.Comment))
}

// QueryResShitIDs -
type QueryResShitIDs []string

func (n QueryResShitIDs) String() string {
	return strings.Join(n[:], "\n")
}

// Shits -
type Shits []*Shit
