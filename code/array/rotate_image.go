package array

import "fmt"

func printMatrix(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			fmt.Printf(" %d", mat[i][j])
		}
		fmt.Println(" ")
	}
}

func rotateMatrix90DegreesByAntiClockwise(matrix [][]int) {
	n := len(matrix)
	for x := 0; x < n/2; x++ {
		for y := x; y < n-x-1; y++ {
			temp := matrix[x][y]
			matrix[x][y] = matrix[y][n-x-1]
			matrix[y][n-x-1] = matrix[n-x-1][n-y-1]
			matrix[n-x-1][n-y-1] = matrix[n-y-1][x]
			matrix[n-y-1][x] = temp
		}
	}
	printMatrix(matrix)
}

func rotateMatrix90DegreesByClockwise(matrix [][]int) {
	n := len(matrix)
	for x := 0; x < n/2; x++ {
		for y := x; y < n-x-1; y++ {
			temp := matrix[x][y]
			matrix[x][y] = matrix[n-y-1][x]
			matrix[n-y-1][x] = matrix[n-x-1][n-y-1]
			matrix[n-x-1][n-y-1] = matrix[y][n-x-1]
			matrix[y][n-x-1] = temp
		}
	}
	printMatrix(matrix)
}

func Rotate(matrix [][]int) {
	rotateMatrix90DegreesByClockwise(matrix)
}
