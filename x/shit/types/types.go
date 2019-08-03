package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Shit -
type Shit struct {
	ID         string           `json:"id"`
	ShitType   uint8            `json:"shit_type"`
	Informants []sdk.AccAddress `json:"informants"`
	Comments   []string         `jsong:"comments"`
}

func (s Shit) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %d
ShitType: %d
Informants: %v
Comments: %v
`, s.ID, s.ShitType, s.Informants, s.Comments))
}

// QueryResShitIDs -
type QueryResShitIDs []string

func (n QueryResShitIDs) String() string {
	return strings.Join(n[:], "\n")
}

// Shits -
type Shits []*Shit
