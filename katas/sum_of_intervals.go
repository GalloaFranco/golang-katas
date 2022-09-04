package katas

import (
	"fmt"
	"sort"
)

func SumOfIntervals() int {
	arr := [][2]int{{1, 4}, {7, 10}, {3, 5}}

	sort.Slice(arr[:], func(i, j int) bool {
		for x := range arr[i] {
			if arr[i][x] == arr[j][x] {
				continue
			}
			return arr[i][x] < arr[j][x]
		}
		return false
	})

	uniqueArr := [][2]int{arr[0]}
	var accArr []int
	var result int
Loop:
	for _, v := range arr {
		for _, duplicate := range uniqueArr {
			if (v[0] != duplicate[0] && v[1] != duplicate[1]) && (v[1] > duplicate[0] && v[0] > duplicate[0]) {
				if v[0] < duplicate[1] && v[1] > duplicate[1] {
					remaining := [][2]int{{0, v[1] - duplicate[1]}}
					uniqueArr = append(uniqueArr, remaining[0])
				} else {
					uniqueArr = append(uniqueArr, v)
				}
				continue Loop
			}
		}
	}

	for _, value := range uniqueArr {
		val := value[1] - value[0]
		accArr = append(accArr, val)
	}

	for _, partialResult := range accArr {
		result += partialResult
	}
	fmt.Printf("Result: %v", result)
	return result
}
