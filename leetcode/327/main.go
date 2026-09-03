// Package leetcode327 https://leetcode.com/problems/count-of-range-sum/
package leetcode327

func countRangeSum(nums []int, lower, upper int) int {
	numsPrefix := make([]int, len(nums)+1)
	for i := range len(nums) {
		numsPrefix[i+1] = nums[i] + numsPrefix[i]
	}

	return merge(numsPrefix, lower, upper)
}

func merge(nums []int, lower, upper int) int {
	count := 0
	if len(nums) <= 1 {
		return count
	}

	left, right := nums[0:len(nums)/2], nums[len(nums)/2:]

	count = merge(left, lower, upper) + merge(right, lower, upper)
	var result []int

	i, j := 0, 0
	for i < len(left) && j < len(right) {

		if lower <= right[j]-left[i] && right[j]-left[i] <= upper {
			count++
		}
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	for i := range len(nums) {
		nums[i] = result[i]
	}

	return count
}
