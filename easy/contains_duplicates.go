package easy

func containsDuplicate(nums []int) bool {
	exists := make(map[int]bool)
	for _, value := range nums {
		if exists[value] {
			return true
		}
		exists[value] = true
	}
	return false
}
