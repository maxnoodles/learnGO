package main

import "fmt"

func noneRepeating(s string) int {
	lastIndexMap := make(map[byte]int)
	start, maxLen := 0, 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastIndexMap[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastIndexMap[ch] = i
	}
	return maxLen
}

func main() {
	fmt.Println(
		noneRepeating("abcabc"),
		noneRepeating(""),
		noneRepeating("abcdefg"),
		noneRepeating("aabbccabc"),
	)
}
