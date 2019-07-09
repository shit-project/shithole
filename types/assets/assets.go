package assets

// MicroShitDenom -
const (
	MicroShitDenom = "ushit"
	MicroUnit      = int64(1e6)
)

// IsValidDenom -
func IsValidDenom(denom string) bool {
	return denom == MicroShitDenom
}
