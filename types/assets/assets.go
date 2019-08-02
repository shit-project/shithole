package assets

const (
	// MicroShitDenom -
	MicroShitDenom = "ushit"

	// MicroUnit -
	MicroUnit = int64(1e6)
)

// IsValidDenom -
func IsValidDenom(denom string) bool {
	return denom == MicroShitDenom
}
