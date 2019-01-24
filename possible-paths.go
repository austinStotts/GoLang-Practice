package main

import (
	"fmt"
)

// returns an NxN slice filled with "0"
func makeMatrix(n int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}
	return matrix
}

func check(matrix [][]int, y int, x int, n int) bool {
	if y < 0 || y >= n {
		return false
	}
	if x < 0 || x >= n {
		return false
	}
	if matrix[y][x] == 1 {
		return false
	}
	return true
}

func loop(matrix [][]int, y, x, n int) int {
	solutions := 0
	if y == n-1 && x == n-1 {
		matrix[y][x] = 0
		return 1
	}

	// up
	if check(matrix, y-1, x, n) {
		matrix[y][x] = 1
		solutions = solutions + loop(matrix, y-1, x, n)
	}

	// right
	if check(matrix, y, x+1, n) {
		matrix[y][x] = 1
		solutions = solutions + loop(matrix, y, x+1, n)
	}

	// down
	if check(matrix, y+1, x, n) {
		matrix[y][x] = 1
		solutions = solutions + loop(matrix, y+1, x, n)
	}

	// left
	if check(matrix, y, x-1, n) {
		matrix[y][x] = 1
		solutions = solutions + loop(matrix, y, x-1, n)
	}
	matrix[y][x] = 0
	return solutions
}

func main() {
	fmt.Println(loop(makeMatrix(6), 0, 0, 6))
}
