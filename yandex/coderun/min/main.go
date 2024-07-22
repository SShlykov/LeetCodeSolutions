package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}

	largestSquareSide, topLeftX, topLeftY := findLargestSquareUsingGraph(matrix, n, m)
	fmt.Println(largestSquareSide)
	fmt.Println(topLeftX+1, topLeftY+1)
}

func findLargestSquareUsingGraph(matrix [][]int, n, m int) (int, int, int) {
	// Directions for right, down, down-right diagonal, and down-left diagonal
	directions := []Point{{0, 1}, {1, 0}, {1, 1}, {1, -1}}

	// Create visited matrix
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	maxSide := 0
	topLeftX := 0
	topLeftY := 0

	// DFS to explore the connected components
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		visited[x][y] = true
		sideLength := 1

		for _, dir := range directions {
			nx, ny := x+dir.x, y+dir.y
			if nx >= 0 && nx < n && ny >= 0 && ny < m && matrix[nx][ny] == 1 && !visited[nx][ny] {
				sideLength = max(sideLength, 1+dfs(nx, ny))
			}
		}

		return sideLength
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 && !visited[i][j] {
				sideLength := dfs(i, j)
				if sideLength > maxSide {
					maxSide = sideLength
					topLeftX = i
					topLeftY = j
				}
			}
		}
	}

	return maxSide, topLeftX, topLeftY
}
