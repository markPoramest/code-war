package question4kyu

import (
	"math"
	"sort"
	"strconv"
)

// ref: https://www.codewars.com/kata/55983863da40caa2c900004e

func NextBigger(n int) int {
	number := strconv.Itoa(n)
	digits := []byte(number)
	pivot := -1
	for i := len(digits) - 2; i >= 0; i-- {
		if digits[i] < digits[i+1] {
			pivot = i
			break
		}
	}

	if pivot == -1 {
		return pivot
	}

	minValue := math.MaxInt64
	index := -1
	for k := pivot + 1; k < len(digits); k++ {
		if int(digits[k]) <= minValue && int(digits[k]) > int(digits[pivot]) {
			minValue = int(digits[k])
			index = k
		}
	}
	digits[pivot], digits[index] = digits[index], digits[pivot]

	sort.Slice(digits[pivot+1:], func(a, b int) bool {
		return digits[pivot+1+a] < digits[pivot+1+b]
	})

	resultStr := string(digits)
	resultInt, _ := strconv.Atoi(resultStr)
	return resultInt
}
