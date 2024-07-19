package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func closest(currposition int, stops []int) int {
	l, r := -1, len(stops)-1
	good := func(m int) bool {
		return stops[m] > currposition
	}

	for r-l > 1 {
		m := (l + r) / 2
		if good(m) {
			r = m
		} else {
			l = m
		}
	}

	if abs(stops[l]-currposition) <= abs(stops[r]-currposition) {
		return l
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Read n and k
	if !scanner.Scan() {
		log.Fatal("Failed to read n")
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	if !scanner.Scan() {
		log.Fatal("Failed to read k")
	}
	k, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// Read stops
	stops := make([]int, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			log.Fatal("Failed to read stop")
		}
		stops[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < k; i++ {
		if !scanner.Scan() {
			log.Fatal("Failed to read current position")
		}
		currposition, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(closest(currposition, stops) + 1)
	}
}
