package main

import "fmt"

// returns an NxN slice filled with "0"
func makeMatrix(n int) *[][]int {
	matrix := make([][]int, 0)
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}
	return &matrix
}

func check(matrix *[][]int, y int, x int, n int) bool {
	m := *matrix
	if y < 0 || y >= n {
		return false
	}
	if x < 0 || x >= n {
		return false
	}
	if m[y][x] == 1 {
		return false
	}
	return true
}

func loop(matrix *[][]int, y, x, n int) int {
	mtrx := *matrix
	solutions := 0
	if y == n-1 && x == n-1 {
		mtrx[y][x] = 0
		return 1
	}

	// up
	if check(&mtrx, y-1, x, n) {
		mtrx[y][x] = 1
		solutions = solutions + loop(&mtrx, y-1, x, n)
	}

	// right
	if check(&mtrx, y, x+1, n) {
		mtrx[y][x] = 1
		solutions = solutions + loop(&mtrx, y, x+1, n)
	}

	// down
	if check(&mtrx, y+1, x, n) {
		mtrx[y][x] = 1
		solutions = solutions + loop(&mtrx, y+1, x, n)
	}

	// left
	if check(&mtrx, y, x-1, n) {
		mtrx[y][x] = 1
		solutions = solutions + loop(&mtrx, y, x-1, n)
	}
	mtrx[y][x] = 0
	return solutions
}

func main() {
	var result = loop(makeMatrix(6), 0, 0, 6)
	fmt.Println(result)
}
