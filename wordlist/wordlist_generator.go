package wordlist

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strings"
)

func ReadAndCreateWordList(filePath string, baseWords []string) [][]rune {
	unfilteredRuneList := createRuneList(filePath)
	fmt.Println("Unfiltered List length: ", len(unfilteredRuneList))
	baseWordRuneMap := map[rune]int{}
	for _, word := range baseWords {
		for _, character := range word {
			baseWordRuneMap[character] += 1
		}
	}
	filteredRuneList := filterWordList(unfilteredRuneList, baseWordRuneMap)

	fmt.Println("Filtered List length: ", len(filteredRuneList))
	return filteredRuneList
}

func filterWordList(candidateRuneList [][]rune, baseWordRuneMap map[rune]int) [][]rune {
	filteredRuneList := [][]rune{}
	for _, candidateWord := range candidateRuneList {
		if !wordContainsInvalidRunes(candidateWord, baseWordRuneMap) {
			filteredRuneList = append(filteredRuneList, candidateWord)
		}
	}
	return filteredRuneList
}

func createRuneList(filePath string) [][]rune {
	runeList := [][]rune{}
	fileReader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(fileReader)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.ToLower(line)
		runeList = append(runeList, []rune(line))
	}
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}
	return runeList
}

func wordContainsInvalidRunes(word []rune, baseRuneList map[rune]int) bool {
	cloneRuneMap := maps.Clone(baseRuneList)
	for _, character := range word {
		count, present := cloneRuneMap[character]
		if !present || count == 0 {
			return true
		} else {
			cloneRuneMap[character] -= 1
		}
	}
	return false
}
