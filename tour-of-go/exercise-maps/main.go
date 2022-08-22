package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	arr := strings.Fields(s)
	for _, word := range arr {
		_, ok := res[word]
		if ok {
			res[word] += 1
		} else {
			res[word] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
