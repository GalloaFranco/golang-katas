package katas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindOdd(t *testing.T) {
	case1 := []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}
	case2 := []int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}
	case3 := []int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5}
	case4 := []int{10}
	case5 := []int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}
	case6 := []int{}
	testCases := map[string]func(t *testing.T){
		"Case 1": func(t *testing.T) {
			result := FindOdd(case1)
			assert.Equal(t, 5, result)
		},
		"Case 2": func(t *testing.T) {
			result := FindOdd(case2)
			assert.Equal(t, -1, result)
		},
		"Case 3": func(t *testing.T) {
			result := FindOdd(case3)
			assert.Equal(t, 5, result)
		},
		"Case 4": func(t *testing.T) {
			result := FindOdd(case4)
			assert.Equal(t, 10, result)
		},
		"Case 5": func(t *testing.T) {
			result := FindOdd(case5)
			assert.Equal(t, 10, result)
		},
		"Case 6": func(t *testing.T) {
			result := FindOdd(case6)
			assert.Equal(t, 0, result)
		},
	}
	for name, tc := range testCases {
		t.Run(name, tc)
	}
}
