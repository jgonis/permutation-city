package runemap

import (
	"maps"
	"testing"
)

func TestCreateRuneMap(t *testing.T) {
	testCases := map[string]struct {
		baseWords []string
		want      RuneMap
	}{
		"Single word": {
			baseWords: []string{"caret"},
			want:      RuneMap{"c": 1, "a": 1, "r": 1, "e": 1, "t": 1},
		},
		"Multiple words": {
			baseWords: []string{"car", "et"},
			want:      RuneMap{"c": 1, "a": 1, "r": 1, "e": 1, "t": 1},
		},
		"Multiple words, with repeated runes": {
			baseWords: []string{"car", "et", "caret"},
			want:      RuneMap{"c": 2, "a": 2, "r": 2, "e": 2, "t": 2},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := CreateRuneMap(testCase.baseWords)
			if !maps.Equal(got, testCase.want) {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestRemoveRunesFromWord(t *testing.T) {
	testCases := map[string]struct {
		startingMap RuneMap
		word        []rune
		expected    RuneMap
	}{
		"Remove single rune": {
			startingMap: RuneMap{"c": 1, "a": 1, "r": 1, "e": 1, "t": 1},
			word:        []rune{'c'},
			expected:    RuneMap{"a": 1, "r": 1, "e": 1, "t": 1},
		},
		"Remove multiple runes": {
			startingMap: RuneMap{"c": 1, "a": 1, "r": 1, "e": 1, "t": 1},
			word:        []rune{'c', 'a', 'r'},
			expected:    RuneMap{"e": 1, "t": 1},
		},
		"Remove runes with repetition": {
			startingMap: RuneMap{"c": 2, "a": 2, "r": 2, "e": 2, "t": 2},
			word:        []rune{'c', 'a', 'r', 'e', 't'},
			expected:    RuneMap{"c": 1, "a": 1, "r": 1, "e": 1, "t": 1},
		},
		"Remove all runes": {
			startingMap: RuneMap{"c": 1, "a": 1, "r": 1, "e": 1, "t": 1},
			word:        []rune{'c', 'a', 'r', 'e', 't'},
			expected:    RuneMap{},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			updatedRuneMap := testCase.startingMap.RemoveRunesFromWord(testCase.word)
			if !maps.Equal(updatedRuneMap, testCase.expected) {
				t.Errorf("got %v, want %v", testCase.startingMap, testCase.expected)
			}
		})
	}
}
