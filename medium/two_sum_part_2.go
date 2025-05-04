package medium

func twoSum(nums []int, target int) []int {
	l := 0
	h := len(nums) - 1
	var result [2]int
	for l < h {
		if nums[l]+nums[h] == target {
			result[0] = l + 1
			result[1] = h + 1
			break
		} else if nums[l]+nums[h] > target {
			h--
		} else {
			l++
		}
	}
	return result[:]
}
