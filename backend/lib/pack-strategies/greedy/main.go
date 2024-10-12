package main

import "sort"

func FindIdealPack(order int, packSizes []int) map[int]int {
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	result := make(map[int]int)
	remaining := order

	for _, size := range packSizes {
		if size <= remaining {
			count := remaining / size
			result[size] = count
			remaining -= count * size
		}

		if remaining == 0 {
			break
		}
	}

	// If there's still a remaining amount, add one more of the smallest pack
	if remaining > 0 {
		result[packSizes[len(packSizes)-1]]++
	}

	return result
}
