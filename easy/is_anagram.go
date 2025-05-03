package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq_s [26]int

	for i := 0; i < len(s); i++ {
		freq_s[s[i]-'a']++
		freq_s[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if freq_s[i] != 0 {
			return false
		}
	}

	return true
}
