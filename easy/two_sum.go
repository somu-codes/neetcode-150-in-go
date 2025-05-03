package easy

func twoSum(nums []int, target int) []int {
	var l, h int
	index := make(map[int]int)

	for i, value := range nums {
		if index[value] == 0 {
			index[value] = i + 1
		}
		remaining := target - value
		if index[remaining] != 0 && index[remaining] != i+1 {
			l = i
			h = index[remaining] - 1
			break
		}
	}
	return []int{l, h}
}
