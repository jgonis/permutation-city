package main

import (
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/jgonis/permutation-city/runemap"
	"github.com/jgonis/permutation-city/wordlist"
)

type PermutationResult struct {
	wordList         []string
	averageFrequency uint64
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Did not receive any arguments to generate permutations from")
	}

	baseWordRuneMap := runemap.CreateRuneMap(args)
	filteredWordList := wordlist.ReadAndCreateWordList("wordlist.txt", baseWordRuneMap)
	frequencyMap := wordlist.ReadAndCreateFrequencyMap("unigram_freq.csv")
	fmt.Println("frequency map size: ", len(frequencyMap))
	fmt.Println("Calling recursive function now")
	finalResultList := jeff(baseWordRuneMap, filteredWordList, []string{})

	for i, _ := range finalResultList {
		var averageFrequencyScore uint64 = 0
		for _, word := range finalResultList[i].wordList {
			if score, ok := frequencyMap[word]; ok {
				averageFrequencyScore += score
			}
		}
		averageFrequencyScore = averageFrequencyScore / uint64(len(finalResultList[i].wordList))
		finalResultList[i].averageFrequency = averageFrequencyScore
	}
	slices.SortFunc(finalResultList, func(val1 PermutationResult, val2 PermutationResult) int {
		if val1.averageFrequency < val2.averageFrequency {
			return 1
		} else if val1.averageFrequency > val2.averageFrequency {
			return -1
		}
		return 0
	})
	for i := range min(1000, len(finalResultList)-1) {
		for _, word := range finalResultList[i].wordList {
			fmt.Printf("%s ", word)
		}
		fmt.Println(finalResultList[i].averageFrequency)
	}
	// read in words from frequency list
	// have the words be in a map with their frequency score as the value and the word itself as the key
	// Once you have that, iterate through the result list. For each item iterate through words in the list
	// and find them in the frequency map. Take their frequency score and add it to the total score for the line
	// then divide that score by the number of items in the line. Then store it in a struct and
	// use slices.BinarySearchFunc to find the index where the struct should be inserted in the frequency list
}

func jeff(validCharacters runemap.RuneMap, validWords []string, currentResultList []string) []PermutationResult {
	finalResultList := []PermutationResult{}
	for _, word := range validWords {
		newValidCharacters := validCharacters.RemoveRunesFromWord([]rune(word))
		if len(newValidCharacters) == 0 {
			resultWordList := slices.Clone(currentResultList)
			resultWordList = append(resultWordList, word)
			finalResultList = append(finalResultList, PermutationResult{
				resultWordList,
				0,
			})
		} else {
			newValidWords := wordlist.FilterWordList(validWords, newValidCharacters)
			// append current word to a copy of current word list
			newCurrentResultList := slices.Clone(currentResultList)
			newCurrentResultList = append(newCurrentResultList, word)
			finalResultList = append(finalResultList, jeff(newValidCharacters, newValidWords, newCurrentResultList)...)
		}
	}
	return finalResultList
}
