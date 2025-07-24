package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	limitedqueue "github.com/jgonis/permutation-city/limitedQueue"
	"github.com/jgonis/permutation-city/runemap"
	"github.com/jgonis/permutation-city/wordlist"
)

type PermutationResult struct {
	wordList         []string
	averageFrequency uint64
}

func main() {
	flag.Parse()
	errLogger := log.New(os.Stderr, "", 0)

	args := flag.Args()
	if len(args) == 0 {
		errLogger.Println("Did not receive any arguments to generate permutations from")
	}

	baseWordRuneMap := runemap.CreateRuneMap(args)
	filteredWordList := wordlist.ReadAndCreateWordList("wordlist.txt", baseWordRuneMap)
	frequencyMap := wordlist.ReadAndCreateFrequencyMap("unigram_freq.csv")
	limitedResultList := limitedqueue.CreateLimitedQueue[PermutationResult](1000, func(a, b PermutationResult) int {
		if a.averageFrequency == b.averageFrequency {
			return 0
		} else if a.averageFrequency < b.averageFrequency {
			return 1
		} else {
			return -1
		}
	})
	// generatePermutations(baseWordRuneMap, filteredWordList, []string{}, frequencyMap, &limitedResultList)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for permutationResult := range calculateAverageFrequencyStream(ctx, uniqueWordListStream(ctx, generatePermutationsStream(ctx, baseWordRuneMap, filteredWordList)), frequencyMap) {
		limitedResultList.Insert(permutationResult)
	}

	for _, result := range limitedResultList.GetItems() {
		fmt.Printf("%s %v\n", strings.Join(result.wordList, " "), result.averageFrequency)
	}
}

func calculateFrequencyScore(sortedSlice []string, frequencyMap map[string]uint64) uint64 {
	var averageFrequencyScore uint64 = 0
	for _, resultWord := range sortedSlice {
		if score, ok := frequencyMap[resultWord]; ok {
			averageFrequencyScore += score
		}
	}
	averageFrequencyScore = averageFrequencyScore / uint64(len(sortedSlice))
	return averageFrequencyScore
}

func calculateAverageFrequencyStream(ctx context.Context, wordListStream <-chan []string, frequencyMap map[string]uint64) <-chan PermutationResult {
	permutationResultStream := make(chan PermutationResult)
	go func() {
		defer close(permutationResultStream)
		for wordList := range wordListStream {
			averageFrequencyScore := calculateFrequencyScore(wordList, frequencyMap)
			permResult := PermutationResult{
				wordList,
				averageFrequencyScore,
			}
			select {
			case <-ctx.Done():
				return
			case permutationResultStream <- permResult:
			}
		}
	}()
	return permutationResultStream
}

func uniqueWordListStream(ctx context.Context, wordlistStream <-chan []string) <-chan []string {
	uniqueWordlistStream := make(chan []string)
	go func() {
		defer close(uniqueWordlistStream)
		uniqueWordSet := map[string]bool{}
		for inputWordlist := range wordlistStream {
			sortedWordlist := slices.Clone(inputWordlist)
			slices.Sort(sortedWordlist)
			sortedWordlistAsString := strings.Join(sortedWordlist, " ")
			if _, found := uniqueWordSet[sortedWordlistAsString]; !found {
				uniqueWordSet[sortedWordlistAsString] = true
				select {
				case <-ctx.Done():
					return
				case uniqueWordlistStream <- sortedWordlist:
				}
			}
		}
	}()
	return uniqueWordlistStream
}

func generatePermutationsStream(ctx context.Context, validCharacters runemap.RuneMap, validWords []string) <-chan []string {
	wordListStream := make(chan []string)
	var generatePerms func(ctx context.Context, validCharacters runemap.RuneMap, validWords []string, currentResult []string, permChan chan<- []string)
	generatePerms = func(ctx context.Context, validCharacters runemap.RuneMap, validWords []string, currentResult []string, permChan chan<- []string) {
		for _, currentWord := range validWords {
			currentResult := slices.Clone(currentResult)
			currentResult = append(currentResult, currentWord)
			newRuneMap := validCharacters.RemoveRunesFromWord([]rune(currentWord))
			if len(newRuneMap) == 0 {
				select {
				case <-ctx.Done():
					return
				case permChan <- currentResult:
				}
			} else {
				newValidWords := wordlist.FilterWordList(validWords, newRuneMap)
				generatePerms(ctx, newRuneMap, newValidWords, currentResult, permChan)
			}
		}
	}
	go func() {
		defer close(wordListStream)
		generatePerms(ctx, validCharacters, validWords, []string{}, wordListStream)
	}()
	return wordListStream

}
