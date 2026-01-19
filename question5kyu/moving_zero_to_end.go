package question5kyu

func MoveZeros(arr []int) []int {
	newArr := []int{}
	zeroCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			newArr = append(newArr, arr[i])
			continue
		}
		zeroCount++
	}

	for zeroCount > 0 {
		newArr = append(newArr, 0)
		zeroCount--
	}

	return newArr
}
