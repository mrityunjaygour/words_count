package service

import (
	"errors"
	"sort"
	"strings"
)

// Please create a small service that accepts as input a body of text, such as that from a book, and
// returns the top ten most-used words along with how many times they occur in the text.

type wordWithCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

type words []wordWithCount

func (l words) Len() int           { return len(l) }
func (l words) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l words) Less(i, j int) bool { return l[i].Count > l[j].Count }

//getMostUsedWords function take string as input and returns slice with data of top ten most used words
func getMostUsedWords(input string) (words, error) {

	if input == "" {
		return words{}, errors.New("input can not be blank")
	}

	mostUsed := make(map[string]int) // we created a map to assign word and its count
	wordsArray := strings.Split(input, " ")

	for _, word := range wordsArray {
		count := 1
		value, present := mostUsed[word] // first we are checking if that word is already present in the map
		if present {
			count = value + 1
			mostUsed[word] = count // if word is present we are increasing the occurence count
		} else {
			mostUsed[word] = count // if word is not present then we are assigning it into the map with count 1
		}
	}

	return getTopTenWords(mostUsed), nil
}

func getTopTenWords(wordsMap map[string]int) words {

	var wordsList words
	for word, count := range wordsMap {
		wordsList = append(wordsList, wordWithCount{word, count})
	}
	sort.Sort(wordsList)
	return wordsList[:10]

}
