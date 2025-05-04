package medium

import "sort"

type FrequencyPair struct {
	frequency int
	element   int
}

func topKFrequent(nums []int, k int) []int {
	freq_map := make(map[int]FrequencyPair)
	for _, value := range nums {
		var x FrequencyPair
		x.element = value
		x.frequency = freq_map[value].frequency + 1

		freq_map[value] = x
	}

	var freq_arr []FrequencyPair
	for _, value := range freq_map {
		freq_arr = append(freq_arr, value)
	}

	sort.Slice(freq_arr, func(i, j int) bool {
		return freq_arr[i].frequency > freq_arr[j].frequency
	})

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, freq_arr[i].element)
	}

	return result
}
