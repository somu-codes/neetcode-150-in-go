package medium

import (
	"strconv"
)

// Encode and Decode Strings
// Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.

// Please implement encode and decode

// Example 1:

// Input: ["neet","code","love","you"]

// Output:["neet","code","love","you"]
// Example 2:

// Input: ["we","say",":","yes"]

// Output: ["we","say",":","yes"]
// Constraints:

// 0 <= strs.length < 100
// 0 <= strs[i].length < 200
// strs[i] contains only UTF-8 characters.

func encode(strs []string) string {
	var result string
	for _, value := range strs {
		length := strconv.Itoa(len(value))
		result = result + length + "#" + value
	}
	return result
}

func getCountAndIndex(s string, index int) (count int, newIndex int) {
	count = 0
	for i := index; i < len(s) && s[i] != '#'; i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			panic("Not a number, error while decoding")
		}

		c := int(s[i] - '0')
		count = count*10 + c
		newIndex = i + 1
	}
	return count, newIndex
}

func decode(s string) []string {
	var result []string
	for i := 0; i < len(s); i++ {
		count, index := getCountAndIndex(s, i)
		if index+1+count > len(s) {
			panic("Out of boundss")
		}
		newString := s[index+1 : index+1+count]
		result = append(result, newString)
		i = index + count
	}
	return result
}
