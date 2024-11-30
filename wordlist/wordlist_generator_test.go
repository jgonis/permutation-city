package wordlist_test

import (
	"testing"

	"github.com/jgonis/permutation-city/wordlist"
)

func TestFilterWordList(t *testing.T) {
	testCases := map[string]struct {
		candidateRunes [][]rune
		wordList       []string
		expectedRunes  [][]rune
	}{
		"Runes not in candidate words": {
			candidateRunes: [][]rune{[]rune("car"), []rune("et"), []rune("caret"), []rune("zoo")},
			wordList:       []string{"car", "et"},
			expectedRunes:  [][]rune{[]rune("car"), []rune("et"), []rune("caret")},
		},
		// "Word with too many runes than are in candidate words": {
		// 	candidateRunes: [][]rune{[]rune("car"), []rune("et"), []rune("caret"), []rune("carret"), []rune("zoo")},
		// 	wordList:       []string{"car", "et"},
		// 	expectedRunes:  [][]rune{[]rune("car"), []rune("et"), []rune("caret")},
		// },
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			wordlist.FilterWordlist()
		})
	}
}
