package katas

import (
	"sort"
)

func FindOdd(seq []int) int {
	sort.Slice(seq, func(i, j int) bool {
		return seq[i] < seq[j]
	})
	count := 0
	for i := 0; i < len(seq); i++ {
		for j := 0; j < len(seq); j++ {
			if seq[i] == seq[j] {
				count++
			}
		}
		if count%2 != 0 {
			return seq[i]
		}
	}
	return 0
}
