package runemap

import (
	"maps"
)

type RuneMap map[string]int

func CreateRuneMap(baseWords []string) RuneMap {
	runeMap := RuneMap{}
	for _, word := range baseWords {
		for _, character := range word {
			runeMap[string(character)] += 1
		}
	}
	return runeMap
}

func (rm RuneMap) RemoveRunesFromWord(word []rune) RuneMap {
	newRuneMap := maps.Clone(rm)
	for _, character := range word {
		newRuneMap[string(character)] -= 1
		if newRuneMap[string(character)] == 0 {
			delete(newRuneMap, string(character))
		}
	}
	return newRuneMap
}

func (rm RuneMap) IsWordValid(word string) bool {
	runeMapFromWord := CreateRuneMap([]string{word})
	for key, count := range runeMapFromWord {
		if rm[key] < count {
			return false
		}
	}
	return true
}
