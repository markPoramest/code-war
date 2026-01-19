package question5kyu

//Write an algorithm that takes an array and moves all the zeros to the end, preserving the order of the other elements.
//MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }

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
