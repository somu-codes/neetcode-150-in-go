package medium

func computeAlphabeticValue(s string) [26]int {
	var freq [26]int
	for _, ch := range s {
		freq[ch-'a']++
	}
	return freq
}

func groupAnagrams(strs []string) [][]string {
	grouped := make(map[[26]int][]string)

	for _, str := range strs {
		key := computeAlphabeticValue(str)
		grouped[key] = append(grouped[key], str)
	}

	var result [][]string
	for _, group := range grouped {
		result = append(result, group)
	}
	return result
}
