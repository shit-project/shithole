package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Shit -
type Shit struct {
	ID            string         `json:"id"`
	Type    uint8          `json:"Type"`
	Informants         []sdk.AccAddress `json:"informants"`
	Comments	[]string `jsong:"comments"`
}

func (s Shit) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %d
Type: %d
Informants: %v
Comments: %v
`, s.ID, s.Type, s.Informants, s.Comments))
}

// QueryResShitIDs -
type QueryResShitIDs []string

func (n QueryResShitIDs) String() string {
	return strings.Join(n[:], "\n")
}

// Shits -
type Shits []*Shit
