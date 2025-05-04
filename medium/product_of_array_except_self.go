package medium

func countNoOfZeroes(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		}
	}
	return count
}

func findProductExcludingZero(nums []int) int {
	product := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			product *= nums[i]
		}
	}
	return product
}

func productExceptSelf(nums []int) []int {
	no_of_zeroes := countNoOfZeroes(nums)

	if no_of_zeroes == 0 {
		all_product := findProductExcludingZero(nums)
		for i := 0; i < len(nums); i++ {
			nums[i] = all_product / nums[i]
		}
	} else if no_of_zeroes == 1 {
		all_product := findProductExcludingZero(nums)
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				nums[i] = all_product
			} else {
				nums[i] = 0
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			nums[i] = 0
		}
	}
	return nums
}
