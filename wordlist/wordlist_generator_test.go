package wordlist

import (
	"slices"
	"testing"

	"github.com/jgonis/permutation-city/runemap"
)

func TestWordContainsInvalidRunes(t *testing.T) {
	testCases := map[string]struct {
		word        []rune
		baseRuneMap map[rune]int
		expected    bool
	}{
		"Word contains only valid runes": {
			word:        []rune("car"),
			baseRuneMap: map[rune]int{'c': 1, 'a': 1, 'r': 1},
			expected:    false,
		},
		"Word contains invalid runes": {
			word:        []rune("carz"),
			baseRuneMap: map[rune]int{'c': 1, 'a': 1, 'r': 1},
			expected:    true,
		},
		"Word contains more runes than allowed": {
			word:        []rune("carr"),
			baseRuneMap: map[rune]int{'c': 1, 'a': 1, 'r': 1},
			expected:    true,
		},
		"Word contains valid runes with repetition": {
			word:        []rune("carr"),
			baseRuneMap: map[rune]int{'c': 1, 'a': 1, 'r': 2},
			expected:    false,
		},
		"Empty word": {
			word:        []rune(""),
			baseRuneMap: map[rune]int{'c': 1, 'a': 1, 'r': 1},
			expected:    false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := wordContainsInvalidRunes(tc.word, tc.baseRuneMap)
			if result != tc.expected {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}

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
		"Word with too many runes than are in candidate words": {
			candidateRunes: [][]rune{[]rune("car"), []rune("et"), []rune("caret"), []rune("carret"), []rune("zoo")},
			wordList:       []string{"car", "et"},
			expectedRunes:  [][]rune{[]rune("car"), []rune("et"), []rune("caret")},
		},
		"Empty candidate runes": {
			candidateRunes: [][]rune{},
			wordList:       []string{"car", "et"},
			expectedRunes:  [][]rune{},
		},
		"Empty word list": {
			candidateRunes: [][]rune{[]rune("car"), []rune("et"), []rune("caret"), []rune("zoo")},
			wordList:       []string{},
			expectedRunes:  [][]rune{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			baseWordRuneMap := runemap.CreateRuneMap(tc.wordList)
			result := filterWordList(tc.candidateRunes, baseWordRuneMap)
			if !slices.EqualFunc(result, tc.expectedRunes, func(a, b []rune) bool {
				return slices.Equal(a, b)
			}) {
				t.Errorf("got %v, want %v", result, tc.expectedRunes)
			}
		})
	}
}
