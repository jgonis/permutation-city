package main

import (
	"fmt"

	"github.com/jgonis/permutation-city/permutations"
	"github.com/jgonis/permutation-city/wordlist"
)

func main() {
	fmt.Println("Hello, World!")
	permutations.GeneratePermutations(3)
	wordlist.ReadAndCreateWordList("wordlist.txt", []string{"permutation", "city"})

	// generate permutations
	// for each permutations generate a string based on the permutation characters
	// check the string to see if all characters in it can form a word string
	//How to see if it can form a word string?
	// Look at first character and see if it works as a word
	// if yes then look at the next character and perform same check
	// if no, append the next character to the previous character and see if it works as a word
	// if we reach the end of the string, we couldn't form a word string
	// if yes output word string

	//iterate through the list of candidate words
	//for each word figure out the remaining characters left from the input words and the number of characters left in the string
	//create a new list of candidate words that don't use unavailable characters and which aren't too long
	//call a method with the new list of candidate words and the remaining characters, and a slice with the current word at the front

	// in the new method
	// first check to see if the length of the length of the current word slice is equal to the length of the string,
	// if yes we have succeeded and should return the word list
	// if not,
	//if the candidate word list is empty
	//we cannot continue and should return false, no matches found
	//otherwise
	// if not, iterate through the list of candidate words
	// for each word figure out the remaining characters left from the input words and the number of characters left in the string
	// recursively call yourself with the new list of candidate words and the remaining characters, and a slice with the current word at the front
}

func permfinder(targetLength int, candidateWords []string, remainingCharacters map[rune]int, currentWords []string) {
	if currentWordsLength(currentWords) == targetLength {
		fmt.Println("Found a match: ", currentWords)
		return
	}
	if len(candidateWords) == 0 {
		fmt.Println("Can't proceed, no candidate words left")
		return
	}
	for _, word := range candidateWords {
		// for each word figure out the remaining characters left from the input words and the number of characters left in the string
		// create a new list of candidate words that don't use unavailable characters and which aren't too long
		// call a method with the new list of candidate words and the remaining characters, and a slice with the current word at the front
	}
}

func currentWordsLength(currentWords []string) int {
	length := 0
	for _, word := range currentWords {
		length += len(word)
	}
	return length
}

// find words in string
// start at first character and take a substring and see if you can find a word matching that substring
// if you can, then move to the next character and restart the process
// if you can't then move to the next character and now you have a one character large substring, repeat search
// if you try to move to the next character and you have exceeded the length of the string, you couldn't find a match
// if the start character is 1 greater than the length of string, then you created a sentence from the words and should return the word list
// func findWords(start, end int, indices []int, characterList []rune, dictionary map[string]bool, resultWordList []string) {
// 	if start == len(characterList) {
// 		return resultWordList
// 	}
// 	builder := strings.Builder{}
// 	offset := 0
// 	for {
// 		// start at first character and take a substring and see if you can find a word matching that substring
// 		builder.WriteRune(characterList[start+i])
// 		result, present := dictionary[builder.String()]
// 		if present {

// 		}
// 	}
// 	// if you can, then move to the next character and restart the process
// 	// if you can't then move to the next character and now you have a one character large substring, repeat search
// 	// if you try to move to the next character and you have exceeded the length of the string, you couldn't find a match
// 	// if the start character is 1 greater than the length of string, then you created a sentence from the words and should return the word list
// }
