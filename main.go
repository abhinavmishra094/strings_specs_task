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
