package main

import (
	"strconv"
	"strings"
)

func testValidity(input string) bool {
	words := strings.Split(input, "-")
	wCount := 0
	nCount := 0
	for _, word := range words {
		if word == "" {
			return false
		} else if _, err := strconv.Atoi(word); err == nil {
			nCount++
		} else {
			wCount++
		}
		if nCount-wCount == 1 || nCount == wCount {
			continue
		} else {
			return false
		}
	}
	return true
}

func averageNumber(inp string) float32 {
	words := strings.Split(inp, "-")
	nCount := 0
	sum := 0
	for _, word := range words {
		if val, err := strconv.Atoi(word); err == nil {
			nCount++
			sum += val
		}
	}
	return float32(sum) / float32(nCount)
}

func wholeStory(inp string) string {
	words := strings.Split(inp, "-")
	story := ""
	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		} else {
			story += word + " "
		}
	}
	return strings.TrimSpace(story)
}

func storyStats(inp string) (string, string, float32, []string) {
	words := strings.Split(inp, "-")
	var shortestWord string
	var longestWord string
	nWords := 0
	totalLenWords := 0
	var listWord []string
	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		} else {
			if shortestWord == "" || len(word) < len(shortestWord) {
				shortestWord = word
			}
			if longestWord == "" || len(word) > len(longestWord) {
				longestWord = word
			}
			nWords++
			totalLenWords += len(word)
		}
	}
	avgWordLength := float32(totalLenWords) / float32(nWords)
	avgWordLengthRoundedUp := int(avgWordLength) + 1
	avgWordLengthRoundedDown := int(avgWordLength)
	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		} else {
			if len(word) == avgWordLengthRoundedDown || len(word) == avgWordLengthRoundedUp {

				listWord = append(listWord, word)
			}
		}
	}
	return shortestWord, longestWord, avgWordLength, listWord
}
