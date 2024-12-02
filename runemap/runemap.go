package runemap

import "maps"

type RuneMap map[rune]int

func CreateRuneMap(baseWords []string) RuneMap {
	runeMap := RuneMap{}
	for _, word := range baseWords {
		for _, character := range word {
			runeMap[character] += 1
		}
	}
	return runeMap
}

func (rm *RuneMap) RemoveRunesFromWord(word []rune) RuneMap {
	newRuneMap := maps.Clone(*rm)
	for _, character := range word {
		newRuneMap[character] -= 1
		if newRuneMap[character] == 0 {
			delete(newRuneMap, character)
		}
	}
	return newRuneMap
}
