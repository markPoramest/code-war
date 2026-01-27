package question4kyu

import (
	"fmt"
	"strings"
)

// ref: https://www.codewars.com/kata/51ba717bb08c1cd60f00002f

func Solution(list []int) string {
	result := ""
	for i := 0; i < len(list); i++ {
		count := 1
		for j := 1; j < len(list); j++ {
			if i+j < len(list) && list[i]+j == list[i+j] {
				count++
			} else if count >= 3 {
				fmt.Println(count)
				result += fmt.Sprintf("%d-%d,", list[i], list[i+j-1])
				break
			} else {
				for k := i; k < i+count; k++ {
					result += fmt.Sprintf("%d,", list[k])
				}
				break
			}
		}
		i = i + count - 1
	}

	return strings.Trim(result, ",")
}
