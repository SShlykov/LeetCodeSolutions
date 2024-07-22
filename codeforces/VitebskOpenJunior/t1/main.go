package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	factsmemo = map[int]int{
		0: 1,
		1: 1,
	}
)

func main() {
	array := readArray(makeScanner())
	n, k := array[0], array[1]

	res := factorial(n) / (factorial(n-k) * factorial(k))

	fmt.Println(res)
}

func factorial(i int) int {
	if val, exist := factsmemo[i]; exist {
		return val
	}

	return i * factorial(i-1)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}
func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
