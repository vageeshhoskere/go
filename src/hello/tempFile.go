package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	wordMap := map[string]int{}
	for _, str := range strs {
		cnt := wordMap[str]
		wordMap[str] = cnt + 1
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
