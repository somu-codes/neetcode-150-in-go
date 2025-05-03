package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq_s [26]int
	var freq_t [26]int

	for i := 0; i < len(s); i++ {
		freq_s[s[i]-'a']++
		freq_t[t[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if freq_s[i] != freq_t[i] {
			return false
		}
	}

	return true
}
