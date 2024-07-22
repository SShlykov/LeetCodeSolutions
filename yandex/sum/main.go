package main

import (
	"os"
	"strconv"
	"strings"
)

func getSum(a int, b int) int {
	return a + b
}

func main() {
	a, b := readInts()
	file, _ := os.Open("output.txt")
	defer file.Close()
	file.Write([]byte(strconv.Itoa(getSum(a, b))))

}

func readInts() (int, int) {
	bytes, _ := os.ReadFile("input.txt")
	str := string(bytes)
	ints := strings.Split(str, " ")
	a, _ := strconv.Atoi(ints[0])
	b, _ := strconv.Atoi(ints[1])
	return a, b
}
