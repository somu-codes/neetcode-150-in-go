package medium

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	answer := 1
	curr := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			curr++
		} else {
			curr = 1
		}
		if curr > answer {
			answer = curr
		}
	}

	return answer
}
