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
