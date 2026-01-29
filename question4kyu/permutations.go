package question4kyu

import (
	"sort"
)

// ref: https://www.codewars.com/kata/5254ca2719453dcc0b00027d

func Permutations(s string) []string {
	resultMapper := map[string]bool{}

	resultMapper = permute(s, 0, resultMapper)

	res := []string{}
	for k := range resultMapper {
		res = append(res, k)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func permute(s string, currentIndex int, result map[string]bool) map[string]bool {
	if currentIndex == len(s) {
		result[s] = true
		return result
	}

	for i := currentIndex; i < len(s); i++ {
		newString := swap(s, i, currentIndex)
		result[newString] = true

		result = permute(newString, currentIndex+1, result)
	}

	return result
}

func swap(s string, currentIndex int, swapIndex int) string {
	bytesStr := []byte(s)
	bytesStr[currentIndex], bytesStr[swapIndex] = bytesStr[swapIndex], bytesStr[currentIndex]

	return string(bytesStr)
}
