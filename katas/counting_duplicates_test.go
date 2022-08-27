package katas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingDuplicate(t *testing.T) {

	testCases := map[string]func(t *testing.T){
		"GivenANoRepeatValueReturnZero": func(t *testing.T) {
			result := Duplicate_count("abcde")
			assert.Equal(t, 0, result)
		},
		"GivenARepeatedValueReturnTwo": func(t *testing.T) {
			result := Duplicate_count("aabbcde")
			assert.Equal(t, 2, result)
		},
		"GivenARepeatedValueInUpperCaseReturnTwo": func(t *testing.T) {
			result := Duplicate_count("aabBcde")
			assert.Equal(t, 2, result)
		},
		"GivenARepeatedValueSeveralTimesReturnOne": func(t *testing.T) {
			result := Duplicate_count("indivisibility")
			assert.Equal(t, 1, result)
		},
		"GivenARepeatedValueWithNumbersReturnTwo": func(t *testing.T) {
			result := Duplicate_count("aA11")
			assert.Equal(t, 2, result)
		},
		"GivenARepeatedValueAllInUpperCaseReturnTwo": func(t *testing.T) {
			result := Duplicate_count("ABBA")
			assert.Equal(t, 2, result)
		},
	}
	for name, tc := range testCases {
		t.Run(name, tc)
	}
}
