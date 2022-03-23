package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const maxCountOfResult = 10

type frequencyOfWordsMap map[string]int

func Top10(text string) []string {
	words := parseWords(text)
	frequencyOfWords := calcFrequencyOfWords(words)
	sortedWords := getSortedWords(frequencyOfWords)

	return getSliceByMaxLen(sortedWords, maxCountOfResult)
}

func parseWords(text string) []string {
	if text == "" {
		return nil
	}

	return strings.Fields(text)
}

func calcFrequencyOfWords(words []string) frequencyOfWordsMap {
	result := make(frequencyOfWordsMap)
	for _, word := range words {
		if word == "" {
			continue
		}
		_, ok := result[word]
		if ok {
			result[word]++
		} else {
			result[word] = 1
		}
	}

	return result
}

func getSortedWords(frequencyOfWords frequencyOfWordsMap) []string {
	uniqWords := make([]string, 0, len(frequencyOfWords))
	for word := range frequencyOfWords {
		uniqWords = append(uniqWords, word)
	}

	sort.Slice(uniqWords, func(i, j int) bool {
		wordI := uniqWords[i]
		wordJ := uniqWords[j]
		wordFrequencyI := frequencyOfWords[wordI]
		wordFrequencyJ := frequencyOfWords[wordJ]
		if wordFrequencyI == wordFrequencyJ {
			return wordI < wordJ
		}

		return wordFrequencyI > wordFrequencyJ
	})

	return uniqWords
}

func getSliceByMaxLen(uniqWords []string, maxLen int) []string {
	if len(uniqWords) >= maxLen {
		return uniqWords[0:maxLen]
	}

	return uniqWords[0:]
}
