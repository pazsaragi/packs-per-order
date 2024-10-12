package backtracking

import (
	strategy "packs-per-order/lib/pack-strategies"
	"sort"
)

// BacktrackingMemoStrategy extends the PackStrategy
type BacktrackingMemoStrategy struct {
	strategy.PackStrategy
	memo      map[int]strategy.PackResult
	packSizes []int
}

func NewPackageStrategy(packSizes []int) *BacktrackingMemoStrategy {
	// Sort pack sizes from smallest to largest
	sortedSizes := make([]int, len(packSizes))
	copy(sortedSizes, packSizes)
	sort.Slice(sortedSizes, func(i, j int) bool {
		return sortedSizes[i] < sortedSizes[j]
	})

	return &BacktrackingMemoStrategy{
		memo:      make(map[int]strategy.PackResult),
		packSizes: sortedSizes,
	}
}

/*
The approach:

[Optimsation]: use the greedy approach for orders larger than the largest pack size.

Find all combinations for a given order using backtracking.

Iterate through each combination and optimise for:

1) Least amount of items
2) Least amount of packs

[Optimisation]: memoize results.
*/
func (p *BacktrackingMemoStrategy) FindIdealPack(order int) strategy.PackResult {
	// If the order exists in the memo,
	// return the existing result
	if result, exists := p.memo[order]; exists {
		return result
	}

	// For orders greater than the maximum, we can cut down the search space by using a greedy approach
	// for the largest pack size. i.e. for 12,001 we can take
	// 2 x 5,000 packs and just finding the optimum solution for 2,001.
	largestSize := p.packSizes[len(p.packSizes)-1]
	largestSizeCount := order / largestSize
	remainingOrder := order % largestSize

	baseCombination := make([]int, largestSizeCount)
	for i := range baseCombination {
		baseCombination[i] = largestSize
	}

	allCombinations := [][]int{}
	p.findCombinationsHelper(remainingOrder, []int{}, &allCombinations)

	var bestCombination []int
	minTotalItems := int(^uint(0) >> 1) // Max int value
	minPackCount := len(p.packSizes)    // Initialize with worst case

	for _, combination := range allCombinations {
		finalCombo := append(baseCombination, combination...)
		totalItems := sum(finalCombo)
		packCount := len(finalCombo)

		// if the current combination items is less than or equal to the current minimum
		// and the number of packs is less, then update the best combination
		if totalItems < minTotalItems || (totalItems == minTotalItems && packCount < minPackCount) {
			bestCombination = finalCombo
			minTotalItems = totalItems
			minPackCount = packCount
		}
	}

	memoResult := strategy.PackResult{
		Packs: convertCombinationToResult(bestCombination),
	}

	// Memoize the result
	p.memo[order] = memoResult

	return memoResult
}

func convertCombinationToResult(combination []int) map[int]int {
	result := make(map[int]int)
	for _, packSize := range combination {
		result[packSize]++
	}
	return result
}

func sum(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

/*
This helper uses backtracking to find all combinations of pack sizes for a given order.

TODO: can we prune the results to avoid calculating combinations for combinations than
have more items and more packs than the current best candidate.
*/
func (p *BacktrackingMemoStrategy) findCombinationsHelper(remainingOrder int, currentCombination []int, allCombinations *[][]int) {
	if remainingOrder <= 0 {
		// If the order is exactly fulfilled, record the current combination
		combinationCopy := make([]int, len(currentCombination))
		copy(combinationCopy, currentCombination)
		*allCombinations = append(*allCombinations, combinationCopy)
		return
	}

	// Explore all pack sizes
	for _, packSize := range p.packSizes {
		// Add the current pack size to the combination and recurse
		currentCombination = append(currentCombination, packSize)
		p.findCombinationsHelper(remainingOrder-packSize, currentCombination, allCombinations)
		// Backtrack
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}
