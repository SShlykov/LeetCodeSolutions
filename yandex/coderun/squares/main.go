package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MOD = 1_000_000_007

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	a := make([]int, n)
	basketStr := strings.Split(scanner.Text(), " ")
	for i := range basketStr {
		a[i], _ = strconv.Atoi(basketStr[i])
	}

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < q; i++ {
		scanner.Scan()
		query := strings.Split(scanner.Text(), " ")
		t, _ := strconv.Atoi(query[0])
		l, _ := strconv.Atoi(query[1])
		r, _ := strconv.Atoi(query[2])
		l, r = l-1, r-1

		switch t {
		case 0:
			for j := l; j <= r; j++ {
				a[j]++
			}
		case 1:
			count := 1
			for j := l; j <= r; j++ {
				count = (count * a[j]) % MOD
			}
			fmt.Fprintln(writer, count)
		}
	}
}
