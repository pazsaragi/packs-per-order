package main

import (
	"reflect"
	"testing"
)

// var PACK_SIZES = []int{250, 500, 1000, 2000, 5000}
var PACK_SIZES = []int{5000, 2000, 1000, 500, 250}
var TEST_CASES = []struct {
	order    int
	expected map[int]int
}{
	{1, map[int]int{250: 1}},
	{250, map[int]int{250: 1}},
	// The greedy approach fails at boundary cases, since it only
	// optimises for the least amount of packets.
	// {251, map[int]int{500: 1}},
	// {501, map[int]int{500: 1, 250: 1}},
	// {751, map[int]int{1000: 1}},
	{1001, map[int]int{1000: 1, 250: 1}},
	// {1251, map[int]int{1000: 1, 500: 1}},
	// {1501, map[int]int{1000: 1, 500: 1, 250: 1}},
	// {1751, map[int]int{1000: 2}},
	{2001, map[int]int{2000: 1, 250: 1}},
	{5001, map[int]int{5000: 1, 250: 1}},
	{12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
}

func TestExample(t *testing.T) {
	for _, tc := range TEST_CASES {
		result := FindIdealPack(tc.order, PACK_SIZES)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("For order %d, expected %v but got %v", tc.order, tc.expected, result)
		}
	}
}
