package question5kyu

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

//In this kata, you will write a function that returns the positions and the values of the "peaks" (or local maxima) of a numeric array.
//
//For example, the array arr = [0, 1, 2, 5, 1, 0] has a peak at position 3 with a value of 5 (since arr[3] equals 5).
//
//The output will be returned as an object with two properties: pos and peaks. Both of these properties should be arrays. If there is no peak in the given array, then the output should be {pos: [], peaks: []}.
//
//Example: pickPeaks([3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3]) should return {pos: [3, 7], peaks: [6, 3]} (or equivalent in other languages)

func PickPeaks(array []int) PosPeaks {
	pos := []int{}
	peaks := []int{}
	nextPosition := 0
	for i := 1; i < len(array)-1; i++ {
		isLeftLess := true
		isRightLess := true
		for less := i - 1; less >= 0; less-- {
			if array[i] > array[less] {
				break
			}
			if array[i] < array[less] {
				isLeftLess = false
				break
			}
		}
		for more := i + 1; more < len(array); more++ {
			if array[i] == array[more] && more == len(array)-1 { // for ignore if edge is a peak
				isRightLess = false
				break
			}
			if array[i] < array[more] {
				isRightLess = false
				break
			}
			if array[i] > array[more] {
				nextPosition = more
				break
			}
		}
		if isLeftLess && isRightLess {
			pos = append(pos, i)
			peaks = append(peaks, array[i])
			i = nextPosition
		}
	}

	return PosPeaks{
		Pos:   pos,
		Peaks: peaks,
	}
}
