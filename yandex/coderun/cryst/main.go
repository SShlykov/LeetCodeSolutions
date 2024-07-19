package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CharCount struct {
	char  rune
	count int
}

func getStructure(s string) []CharCount {
	if len(s) == 0 {
		return nil
	}
	var structure []CharCount
	currentChar := rune(s[0])
	count := 0
	for _, char := range s {
		if char == currentChar {
			count++
		} else {
			structure = append(structure, CharCount{currentChar, count})
			currentChar = char
			count = 1
		}
	}
	structure = append(structure, CharCount{currentChar, count})
	return structure
}

func canBeEqual(s1, s2, s3 string) bool {
	struct1 := getStructure(s1)
	struct2 := getStructure(s2)
	struct3 := getStructure(s3)

	if len(struct1) != len(struct2) || len(struct1) != len(struct3) {
		return false
	}

	for i := range struct1 {
		if struct1[i].char != struct2[i].char || struct1[i].char != struct3[i].char {
			return false
		}
	}

	return true
}

func transformStringsToEqual(s1, s2, s3 string) string {
	if !canBeEqual(s1, s2, s3) {
		return "IMPOSSIBLE"
	}

	struct1 := getStructure(s1)
	struct2 := getStructure(s2)
	struct3 := getStructure(s3)

	var result strings.Builder

	for i := range struct1 {
		count1 := struct1[i].count
		count2 := struct2[i].count
		count3 := struct3[i].count
		minOperations := findMinOperations(count1, count2, count3)
		result.WriteString(strings.Repeat(string(struct1[i].char), minOperations))
	}

	return result.String()
}

func findMinOperations(a, b, c int) int {
	return min(a, min(b, c))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	s2, _ := reader.ReadString('\n')
	s3, _ := reader.ReadString('\n')

	s1 = strings.TrimSpace(s1)
	s2 = strings.TrimSpace(s2)
	s3 = strings.TrimSpace(s3)

	result := transformStringsToEqual(s1, s2, s3)
	fmt.Println(result)
}
