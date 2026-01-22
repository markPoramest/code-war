package question5kyu

import (
	"strings"
)

//ref: https://www.codewars.com/kata/52bc74d4ac05d0945d00054e

func FirstNonRepeating(str string) string {
	alreadyDuplicate := []uint8{}

	lowerCase := strings.ToLower(str)
	for i := 0; i < len(str); i++ {
		hasFound := false
		for _, v := range alreadyDuplicate {
			if v == lowerCase[i] {
				hasFound = true
			}
		}
		if !hasFound {
			for j := i + 1; j < len(str); j++ {
				if lowerCase[i] == lowerCase[j] {
					hasFound = true
					alreadyDuplicate = append(alreadyDuplicate, lowerCase[i])
					break
				}
			}
		}

		if !hasFound {
			return string(str[i])
		}
	}

	return ""
}
