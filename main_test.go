package main

import (
	"testing"
)

func TestTestValidity(t *testing.T) {
	table := []struct {
		input   string
		isValid bool
	}{
		{"23-ab-48-caba-56-haha", true},
		{"23-ab-48--56-haha", false},
		{"23-ab-48-ram-caba-56-haha", false},
	}

	for _, testCase := range table {
		isValid := testValidity(testCase.input)
		if isValid != testCase.isValid {
			t.Errorf("testValidity was incorrect for input %s, got: %v, want: %v.", testCase.input, isValid, testCase.isValid)
		}
	}

}

func TestAverageNumber(t *testing.T) {
	table := []struct {
		input string
		avg   float32
	}{
		{"23-ab-48-caba-55-haha", 42.0},
		{"12-ab-34--56-haha", 34.0},
		{"23-ab-48-ram-caba--haha", 35.5},
	}

	for _, testCase := range table {
		avg := averageNumber(testCase.input)
		if avg != testCase.avg {
			t.Errorf("AverageNumber was incorrect for input %s, got: %v, want: %v.", testCase.input, avg, testCase.avg)
		}
	}

}

func TestWholeStory(t *testing.T) {
	table := []struct {
		input string
		story string
	}{
		{"23-ab-48-caba-560-haha", "ab caba haha"},
		{"", ""},
		{"1-hello-2-world", "hello world"},
	}

	for _, testCase := range table {
		story := wholeStory(testCase.input)
		if story != testCase.story {
			t.Errorf("WholeStory was incorrect for input %s, got: %v, want: %v.", testCase.input, story, testCase.story)
		}
	}

}

func TestStoryStats(t *testing.T) {
	table := []struct {
		input         string
		shortestWord  string
		longestWord   string
		avgWordLength float32
		listWord      []string
	}{
		{"23-ab-48-caba-560-hahaaa", "ab", "hahaaa", 4.0, []string{"caba"}},
		{"23-a-48-bc-de-fg-560-rhjk", "a", "rhjk", 2.2, []string{"bc", "de", "fg"}},
		{"23-a-48-bcd-ded-fgd-560-rhjk", "a", "rhjk", 2.8, []string{"bcd", "ded", "fgd"}},
	}

	for _, testCase := range table {
		shortestWord, longestWord, avgWordLength, listWord := storyStats(testCase.input)
		if shortestWord != testCase.shortestWord || longestWord != testCase.longestWord || avgWordLength != testCase.avgWordLength || !assertListEqual(listWord, testCase.listWord) {
			t.Errorf(
				"StoryStats was incorrect for input %s", testCase.input,
			)
		}
	}

}

func assertListEqual(lst1 []string, lst2 []string) bool {
	if len(lst1) != len(lst2) {
		return false
	}
	for i := 0; i < len(lst1); i++ {
		if lst1[i] != lst2[i] {
			return false
		}
	}
	return true
}
