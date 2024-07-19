package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()
	array := readArray(scanner)
	x := array[1]
	array = readArray(scanner)
	mx, id, exist := trip(array, x)
	if exist {
		fmt.Printf("%d %d\n", mx, id)
	} else {
		fmt.Println(-1)
	}
}

func trip(arr []int, x int) (int, int, bool) {
	l, r := 0, 0
	maxL, id := 0, -1
	curr := 0

	for r < len(arr) {
		fmt.Println(len(arr), curr, l, r)
		switch {
		case curr > x:
			curr -= arr[l]
			l += 1
		case curr < x:
			curr += arr[r]
			r += 1
		case curr == x:
			if maxL < r-l+1 {
				id = l
				maxL = r - l + 1
			}
			curr -= arr[l]
			l++
		}
	}

	if id == -1 {
		return 0, 0, false
	}

	return maxL, id, true
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
