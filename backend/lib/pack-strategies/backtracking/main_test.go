package backtracking

import (
	strategy "packs-per-order/lib/pack-strategies"
	"reflect"
	"testing"
)

var PACK_SIZES = []int{5000, 2000, 1000, 500, 250}
var TEST_CASES = []struct {
	order    int
	expected strategy.PackResult
}{
	{1, strategy.PackResult{Packs: map[int]int{250: 1}}},
	{250, strategy.PackResult{Packs: map[int]int{250: 1}}},
	{251, strategy.PackResult{Packs: map[int]int{500: 1}}},
	{501, strategy.PackResult{Packs: map[int]int{500: 1, 250: 1}}},
	{751, strategy.PackResult{Packs: map[int]int{1000: 1}}},
	{1001, strategy.PackResult{Packs: map[int]int{1000: 1, 250: 1}}},
	{2001, strategy.PackResult{Packs: map[int]int{2000: 1, 250: 1}}},
	{5001, strategy.PackResult{Packs: map[int]int{5000: 1, 250: 1}}},
	{12001, strategy.PackResult{Packs: map[int]int{5000: 2, 2000: 1, 250: 1}}},
}

func TestExample(t *testing.T) {
	generator := NewPackageStrategy(PACK_SIZES)
	for _, tc := range TEST_CASES {
		result := generator.FindIdealPack(tc.order)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("For order %d, expected %v but got %v", tc.order, tc.expected, result)
		}
	}
}
