// Package leetcode327 https://leetcode.com/problems/count-of-range-sum/
package leetcode327

import "fmt"

func countRangeSum(nums []int, lower, upper int) int {
	fmt.Printf("nums=%v lower=%d upper=%d\n", nums, lower, upper)
	numsPrefix := make([]int, len(nums)+1)
	for i := range len(nums) {
		numsPrefix[i+1] = nums[i] + numsPrefix[i]
	}

	return merge(numsPrefix, lower, upper, 0)
}

func merge(nums []int, lower, upper int, depth int) int {
	prefix := ""
	for range depth {
		prefix = prefix + "\t"
	}
	fmt.Printf("%snums=%v\n", prefix, nums)

	count := 0
	if len(nums) <= 1 {
		return count
	}

	l, r := nums[0:len(nums)/2], nums[len(nums)/2:]

	count = merge(l, lower, upper, depth+1) + merge(r, lower, upper, depth+1)
	newNums := make([]int, len(nums))

	i, j := 0, 0
	for k := range len(nums) {
		fmt.Printf("%si=%d j=%d k=%d count=%d\n", prefix, i, j, k, count)

		if j >= len(r) || i >= len(l) {
			break
		}

		if lower <= r[j]-l[i] && r[j]-l[i] <= upper {
			fmt.Printf("%scount i=%d j=%d\n", prefix, i, j)
			count++
		}

		if l[i] < r[j] {
			newNums[k] = l[i]

			if i+1 >= len(l) {
				j++
			} else {
				i++
			}
		} else {
			newNums[k] = r[j]

			if j+1 >= len(r) {
				i++
			} else {
				j++
			}
		}
	}

	for i := range len(nums) {
		nums[i] = newNums[i]
	}

	fmt.Printf("%ssorted=%v\n", prefix, nums)
	fmt.Printf("%scount=%d\n", prefix, count)
	return count
}
