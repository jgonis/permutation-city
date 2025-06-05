package wordlist

import (
	"bufio"
	"maps"
	"os"
	"strconv"
	"strings"

	"github.com/jgonis/permutation-city/runemap"
)

func ReadAndCreateWordList(filePath string, runeMap runemap.RuneMap) []string {
	unfilteredRuneList := readWordsFromWordList(filePath)
	// fmt.Println("Unfiltered List length: ", len(unfilteredRuneList))
	filteredRuneList := FilterWordList(unfilteredRuneList, runeMap)

	// fmt.Println("Filtered List length: ", len(filteredRuneList))
	return filteredRuneList
}

func ReadAndCreateFrequencyMap(filepath string) map[string]uint64 {
	resultMap := map[string]uint64{}
	unprocessedWords := readWordsFromWordList(filepath)
	for _, word := range unprocessedWords {
		wordAndFrequencyString := strings.Split(word, ",")
		frequency, _ := strconv.ParseUint(wordAndFrequencyString[1], 10, 64)
		resultMap[wordAndFrequencyString[0]] = frequency
	}
	return resultMap
}

func FilterWordList(candidateWordList []string, baseWordRuneMap runemap.RuneMap) []string {
	filteredRuneList := []string{}
	for _, candidateWord := range candidateWordList {
		if !wordContainsInvalidRunes(candidateWord, baseWordRuneMap) {
			filteredRuneList = append(filteredRuneList, string(candidateWord))
		}
	}
	return filteredRuneList
}

func wordContainsInvalidRunes(word string, baseRuneList map[string]int) bool {
	cloneRuneMap := maps.Clone(baseRuneList)
	for _, character := range word {
		count, present := cloneRuneMap[string(character)]
		if !present || count == 0 {
			return true
		} else {
			cloneRuneMap[string(character)] -= 1
		}
	}
	return false
}

func readWordsFromWordList(wordlistFilePath string) []string {
	runeList := []string{}
	fileReader, err := os.Open(wordlistFilePath)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(fileReader)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.ToLower(line)
		runeList = append(runeList, line)
	}
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}
	return runeList
}
