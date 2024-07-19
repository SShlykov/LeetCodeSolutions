package main

import (
	"fmt"
	"log"
	"math"
)

func canFirstTeamHaveMoreStudents(A, B, N int) bool {
	return float64(A) > math.Ceil(float64(B)/float64(N))
}

func main() {
	var a, b, n int
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	if canFirstTeamHaveMoreStudents(a, b, n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
