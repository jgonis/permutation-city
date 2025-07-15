package wordlist

import (
	"slices"
	"testing"

	"github.com/jgonis/permutation-city/runemap"
)

func TestWordContainsInvalidRunes(t *testing.T) {
	testCases := map[string]struct {
		word        []rune
		baseRuneMap map[string]int
		expected    bool
	}{
		"Word contains only valid runes": {
			word:        []rune("car"),
			baseRuneMap: map[string]int{"c": 1, "a": 1, "r": 1},
			expected:    false,
		},
		"Word contains invalid runes": {
			word:        []rune("carz"),
			baseRuneMap: map[string]int{"c": 1, "a": 1, "r": 1},
			expected:    true,
		},
		"Word contains more runes than allowed": {
			word:        []rune("carr"),
			baseRuneMap: map[string]int{"c": 1, "a": 1, "r": 1},
			expected:    true,
		},
		"Word contains valid runes with repetition": {
			word:        []rune("carr"),
			baseRuneMap: map[string]int{"c": 1, "a": 1, "r": 2},
			expected:    false,
		},
		"Empty word": {
			word:        []rune(""),
			baseRuneMap: map[string]int{"c": 1, "a": 1, "r": 1},
			expected:    false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := wordContainsInvalidRunes(string(tc.word), tc.baseRuneMap)
			if result != tc.expected {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}

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
