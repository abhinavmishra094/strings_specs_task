package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// complexity Medium, time O(N)
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

// complexity Easy, time O(N)
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

// complexity easy, time O(N)
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

// complexity Medium, time O(N)
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

// complexity Medium, time O(N)
func generate(isCorrect bool) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	numItems := r1.Intn(10) + 2
	result := ""
	for i := 0; i < numItems; i++ {
		result += fmt.Sprintf("%d-%s-", r1.Intn(1000), randStringBytes(1+r1.Intn(10)))
	}
	if isCorrect {
		return strings.Trim(result, "-")
	} else {
		return result + "----"
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func main() {
	println(generate(true))
	println(generate(false))
}
