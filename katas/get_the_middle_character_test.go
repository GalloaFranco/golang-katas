package katas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMiddleBytesCharacter(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"GivenAnEvenValueReturnTwoMiddleCharacters": func(t *testing.T) {
			result := GetMiddleBytes("test")
			assert.Equal(t, "es", result)
		},
		"GivenAnOddValueReturnTheMiddleCharacter": func(t *testing.T) {
			result := GetMiddleBytes("testing")
			assert.Equal(t, "t", result)
		},
		"GivenOnlyOneCharacterValueReturnTheCharacter": func(t *testing.T) {
			result := GetMiddleBytes("A")
			assert.Equal(t, "A", result)
		},
	}
	for name, tc := range testCases {
		t.Run(name, tc)
	}
}

func TestGetMiddleStringsCharacter(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"GivenAnEvenValueReturnTwoMiddleCharacters": func(t *testing.T) {
			result := GetMiddleStrings("test")
			assert.Equal(t, "es", result)
		},
		"GivenAnOddValueReturnTheMiddleCharacter": func(t *testing.T) {
			result := GetMiddleStrings("testing")
			assert.Equal(t, "t", result)
		},
		"GivenOnlyOneCharacterValueReturnTheCharacter": func(t *testing.T) {
			result := GetMiddleStrings("A")
			assert.Equal(t, "A", result)
		},
	}
	for name, tc := range testCases {
		t.Run(name, tc)
	}
}
