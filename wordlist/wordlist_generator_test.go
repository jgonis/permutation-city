package wordlist

import (
	"slices"
	"testing"

	"github.com/jgonis/permutation-city/runemap"
)

func TestFilterWordList(t *testing.T) {
	testCases := map[string]struct {
		candidateRunes []string
		wordList       []string
		expectedRunes  []string
	}{
		"Runes not in candidate words": {
			candidateRunes: []string{"car", "et", "caret", "zoo"},
			wordList:       []string{"car", "et"},
			expectedRunes:  []string{"car", "et", "caret"},
		},
		"Word with too many runes than are in candidate words": {
			candidateRunes: []string{"car", "et", "caret", "carret", "zoo"},
			wordList:       []string{"car", "et"},
			expectedRunes:  []string{"car", "et", "caret"},
		},
		"Empty candidate runes": {
			candidateRunes: []string{},
			wordList:       []string{"car", "et"},
			expectedRunes:  []string{},
		},
		"Empty word list": {
			candidateRunes: []string{"car", "et", "caret", "zoo"},
			wordList:       []string{},
			expectedRunes:  []string{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			baseWordRuneMap := runemap.CreateRuneMap(tc.wordList)
			result := FilterWordList(tc.candidateRunes, baseWordRuneMap)
			if !slices.EqualFunc(result, tc.expectedRunes, func(a, b string) bool {
				return a == b
			}) {
				t.Errorf("got %v, want %v", result, tc.expectedRunes)
			}
		})
	}
}
