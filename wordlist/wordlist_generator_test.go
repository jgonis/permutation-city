package wordlist_test

import (
	"fmt"
	"testing"
)

func TestFilterWordList(t *testing.T) {
	testCases := map[string]struct {
		candidateRunes [][]rune
		wordList       []string
		expectedRunes  [][]rune
	}{
		"Test 1": {
			candidateRunes: [][]rune{[]rune("car"), []rune("et"), []rune("caret"), []rune("zoo")},
			wordList:       []string{"car", "et"},
			expectedRunes:  [][]rune{[]rune("car"), []rune("et"), []rune("caret")},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			fmt.Println(tc.candidateRunes)
		})
	}
}
