package leetcode327

import (
	"math"
	"testing"
)

func TestCountRangeSum(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		lower int
		upper int
		want  int
	}{
		{
			name:  "example one",
			nums:  []int{-2, 5, -1},
			lower: -2,
			upper: 2,
			want:  3,
		},
		{
			name:  "single zero",
			nums:  []int{0},
			lower: 0,
			upper: 0,
			want:  1,
		},
		{
			name:  "single value in range",
			nums:  []int{5},
			lower: 5,
			upper: 5,
			want:  1,
		},
		{
			name:  "single value outside range",
			nums:  []int{5},
			lower: 1,
			upper: 4,
			want:  0,
		},
		{
			name:  "all zeros",
			nums:  []int{0, 0, 0},
			lower: 0,
			upper: 0,
			want:  6,
		},
		{
			name:  "positive numbers",
			nums:  []int{1, 2, 3},
			lower: 3,
			upper: 5,
			want:  4,
		},
		{
			name:  "negative numbers",
			nums:  []int{-1, -2, -3},
			lower: -5,
			upper: -3,
			want:  3,
		},
		{
			name:  "exact zero sums",
			nums:  []int{1, -1, 1},
			lower: 0,
			upper: 0,
			want:  2,
		},
		{
			name:  "whole array only",
			nums:  []int{1, 2, 3},
			lower: 6,
			upper: 6,
			want:  1,
		},
		{
			name:  "no matching ranges",
			nums:  []int{1, 2, 3},
			lower: 10,
			upper: 20,
			want:  0,
		},
		{
			name:  "32 bit prefix sum overflow",
			nums:  []int{math.MaxInt32, math.MaxInt32},
			lower: -1,
			upper: 1,
			want:  0,
		},
		{
			name:  "large values cancel",
			nums:  []int{math.MaxInt32, -math.MaxInt32},
			lower: 0,
			upper: 0,
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countRangeSum(tt.nums, tt.lower, tt.upper)

			if got != tt.want {
				t.Errorf(
					"countRangeSum(%v, %d, %d) = %d; want %d",
					tt.nums, tt.lower, tt.upper, got, tt.want,
				)
			}
		})
	}
}
