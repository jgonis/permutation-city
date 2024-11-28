package wordlist

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FilterWordlist(candidateRunes []rune, wordlist []string) []string {
	return []string{}
}

func ReadAndCreateWordList(filePath string, baseWords []string) [][]rune {
	unfilteredRuneList := createRuneList(filePath)
	fmt.Println("Unfiltered List length: ", len(unfilteredRuneList))
	baseWordRuneMap := map[rune]bool{}
	for _, word := range baseWords {
		for _, character := range word {
			baseWordRuneMap[character] = true
		}
	}
	filteredRuneList := [][]rune{}
	for _, candidateWord := range unfilteredRuneList {
		if !wordContainsInvalidRunes(candidateWord, baseWordRuneMap) {
			filteredRuneList = append(filteredRuneList, candidateWord)
		}
	}
	fmt.Println("Filtered List length: ", len(filteredRuneList))
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

func wordContainsInvalidRunes(word []rune, baseRuneList map[rune]bool) bool {
	for _, character := range word {
		_, present := baseRuneList[character]
		if !present {
			return true
		}
	}
	return false
}
