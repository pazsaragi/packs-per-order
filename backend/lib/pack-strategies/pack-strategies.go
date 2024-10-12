package packstrategies

type PackResult struct {
	Packs map[int]int
}

// PackStrategy defines the interface that all pack strategies must implement
type PackStrategy interface {
	// FindIdealPack takes an order quantity and returns the ideal pack combination
	FindIdealPack(order int, packSizes []int) PackResult
}
