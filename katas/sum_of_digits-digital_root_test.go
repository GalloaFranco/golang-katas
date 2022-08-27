package katas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"Given_16_Return_7": func(t *testing.T) {
			result := DigitalRoot(16)
			assert.Equal(t, 7, result)
		},
		"Given_942_Return_6": func(t *testing.T) {
			result := DigitalRoot(942)
			assert.Equal(t, 6, result)
		},
		"Given_132189_Return_6": func(t *testing.T) {
			result := DigitalRoot(132189)
			assert.Equal(t, 6, result)
		},
		"Given_493193_Return_11": func(t *testing.T) {
			result := DigitalRoot(493193)
			assert.Equal(t, 2, result)
		},
	}
	for name, tc := range testCases {
		t.Run(name, tc)
	}
}
