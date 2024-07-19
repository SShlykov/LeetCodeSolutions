package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mostFrequentPair(text string) string {
	pairsCount := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		for i := 0; i < len(word)-1; i++ {
			pair := word[i : i+2]
			pairsCount[pair]++
		}
	}

	var mostFrequent string
	maxCount := 0

	for pair, count := range pairsCount {
		if count > maxCount || (count == maxCount && pair > mostFrequent) {
			mostFrequent = pair
			maxCount = count
		}
	}

	return mostFrequent
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Println(mostFrequentPair(text))
}
